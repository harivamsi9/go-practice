package main

import "fmt"

type obeserver interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (dl *DataListener) onUpdate(data string) {

	fmt.Println("Listener:", dl.Name, " Updated with data: ", data)
	// dl.Name = data

}
