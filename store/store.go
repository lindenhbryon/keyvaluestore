package store

import "fmt"

var (
	accessChan chan string
	closeChan  chan struct{}
)

func Start() {

	go func() {
		//m := make(map[string]struct{})

		for {
			select {
			case <-accessChan:
				fmt.Println("accessing the store")
			case <-closeChan:
				fmt.Println("exiting working")
				return
			}
		}
	}()
}
