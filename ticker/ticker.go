package main

import (
	"fmt"
	"time"
)

func GetDurationFromNow(hour, min, sec int) time.Duration {
	h, m, s := time.Now().Clock()

	t1 := time.Duration(h)*time.Hour +
		time.Duration(m)*time.Minute +
		time.Duration(s)*time.Second

	t2 := time.Duration(hour)*time.Hour +
		time.Duration(min)*time.Minute +
		time.Duration(sec)*time.Second

	diff := t2 - t1
	if diff >= 0 {
		return diff
	}
	diff += 24 * time.Hour
	return diff
}

func main() {
	stop := make(chan struct{})
	go func() {
		timer := time.NewTimer(GetDurationFromNow(8, 0, 0))
		fmt.Printf("Waiting until %s",
			time.Now().Add(GetDurationFromNow(8, 0, 0)).Format("15:05:05"))
		select {
		case <-timer.C:
			go func() {
				ticker := time.NewTicker(1 * time.Second)
				for {
					select {
					case <-stop:
						ticker.Stop()
						return
					case t := <-ticker.C:
						fmt.Println("Tick at", t)
					}
				}
			}()
		case <-stop:
			timer.Stop()
			return
		}
	}()
	fmt.Println("Server started")
	time.Sleep(60 * time.Second)
	fmt.Println("Server stopped")
	stop <- struct{}{}
}
