package main

import "fmt"

// implementation
type PlayQuality interface {
	load(string)
}

type SDQuality struct{}
type HDQuality struct{}
type UltraHD struct{}

func (s *SDQuality) load(title string) {
	fmt.Printf("Streaming %s in SD Quality\n", title)
}

func (s *HDQuality) load(title string) {
	fmt.Printf("Streaming %s in HD Quality\n", title)
}

func (s *UltraHD) load(title string) {
	fmt.Printf("Streaming %s in UltraHD Quality\n", title)
}

//abstraction

type VideoPlayer interface {
	play(string)
}

type WebPlayer struct {
	vq PlayQuality
}

func NewWebPlayer(quality PlayQuality) *WebPlayer {
	return &WebPlayer{
		vq: quality,
	}
}

func (w *WebPlayer) play(title string) {
	fmt.Println("Playing video on Web player")
	w.vq.load(title)
}

type SmartTV struct {
	vq PlayQuality
}

func NewSmartTV(quality PlayQuality) *SmartTV {
	return &SmartTV{
		vq: quality,
	}
}

func (w *SmartTV) play(title string) {
	fmt.Println("Playing video on Smart TV")
	w.vq.load(title)
}

func main() {
	web := NewWebPlayer(&HDQuality{})
	web.play("Friends")

	tv := NewSmartTV(&SDQuality{})
	tv.play("Big Bang Theory")
}