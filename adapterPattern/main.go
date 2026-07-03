package main

import (
	"fmt"
)

type Client struct{}

func (c *Client) InsertLightningPortIntoComputer(comp Computer){
	fmt.Println("Inserting Lightning port into computer")
	comp.InsertLightningPort()
}

type Computer interface{
	InsertLightningPort()
}

type Mac struct{}

func (m *Mac) InsertLightningPort(){
	fmt.Println("Lightning port inserted into mac")
}

type Windows struct{}

func (w *Windows) InsertUSBPort(){
	fmt.Println("USB port inserted into windows")
}

type WindowsAdapter struct{
	windows *Windows
}

func (wa *WindowsAdapter) InsertLightningPort(){
	wa.windows.InsertUSBPort()
}

func main(){
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningPortIntoComputer(mac)
	windows := &Windows{}
	wa := &WindowsAdapter{windows: windows,}
	client.InsertLightningPortIntoComputer(wa)
}
