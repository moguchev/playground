package main

import (
	"context"
	"fmt"

	tr "gitlab.services.mts.ru/pepperpotts/api/oebs/transferservice"
	"google.golang.org/grpc"
)

var oebsCli tr.TransferOEBSClient

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		conn.Close()
	}()

	oebsCli := tr.NewTransferOEBSClient(conn)
	// req := tr.GetPositionForCardRequest{
	// 	Id:            2592089,
	// 	AssignmentsId: []int64{129845},
	// }
	req := tr.GetPositionForCardRequest{
		Id:            2973016,
		AssignmentsId: []int64{11167},
	}
	salaries, err := oebsCli.GetPositionForCard(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(salaries)
}
