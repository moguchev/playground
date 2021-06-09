package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"encoding/json"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

type Config map[string]interface{}

type RealConfig struct {
	Version   string            `json:"Version"`
	ProjectID string            `json:"ProjectID"`
	Values    map[string]string `json:"Values"`
}

func ParseFile(path string) (*yaml.Node, error) {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("absolute path: %w", err)
	}

	file, err := os.Open(absolute)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	root := &yaml.Node{}
	if err = yaml.NewDecoder(file).Decode(root); err != nil {
		return nil, fmt.Errorf("decode file: %w", err)
	}

	return root, nil
}

func WriteToFile(path string, node *yaml.Node) error {
	filename, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("absolute path: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	if err = yaml.NewEncoder(file).Encode(node); err != nil {
		return fmt.Errorf("encode file: %w", err)
	}

	return nil
}

func GetRealTimeConfig(ctx context.Context, u *url.URL) (*RealConfig, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", response.StatusCode)
	}

	config := &RealConfig{}
	if err = json.NewDecoder(response.Body).Decode(config); err != nil {
		return nil, fmt.Errorf("decode body: %w", err)
	}

	return config, nil
}

// iterateNode will recursive look for the node following the identifier Node,
// attention: go-yaml has a node for the key and the value itself
func getValueNode(parent *yaml.Node, identifier string) *yaml.Node {
	if parent == nil {
		return nil
	}

	for i, child := range parent.Content {
		if child.Value == identifier {
			return parent.Content[i+1] // возвращаем ноду со значением
		}

		if len(child.Content) > 0 {
			if node := getValueNode(child, identifier); node != nil {
				return node
			}
		}
	}
	return nil
}

func actualizeRealtimeConfig(node *yaml.Node, values map[string]string) {
	for i := range node.Content {
		if node.Content[i].Kind != yaml.ScalarNode {
			continue
		}

		value, ok := values[node.Content[i].Value]
		if !ok {
			continue
		}

		actualizeOption(node.Content[i+1], node.Content[i].Value, value)
	}
}

func actualizeOption(optNode *yaml.Node, name, value string) {
	for i := range optNode.Content {
		if optNode.Content[i].Value == "value" {
			index := i + 1 // next node is value

			if optNode.Content[index].Value != value {
				oldValue := optNode.Content[index].Value

				log.Debugf("option \"%s\" was updated: %s -> %s", name, optNode.Content[index].Value, value)

				optNode.Content[index].Value = value
				optNode.Content[index].LineComment +=
					fmt.Sprintf("previous value: \"%s\"; updated at: \"%s\"", oldValue, time.Now().Format(time.RFC1123))
				break
			}
		}
	}
}

const (
	cfgpath     = "./values.yaml"
	newcfgpath  = "./values_new.yaml"
	urlpath     = "http://metatarifficator.geo.stg.s.o3.ru:84/config"
	realtimeKey = "realtimeConfig"
)

func main() {
	log.SetLevel(log.DebugLevel)

	root, err := ParseFile(cfgpath)
	if err != nil {
		log.Fatal(err)
	}

	realtimeConfigNode := getValueNode(root, realtimeKey)

	u, err := url.Parse(urlpath)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	realCfg, err := GetRealTimeConfig(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	actualizeRealtimeConfig(realtimeConfigNode, realCfg.Values)

	if err = WriteToFile(newcfgpath, root); err != nil {
		log.Fatal(err)
	}
}
