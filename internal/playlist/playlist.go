package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Playlist struct {
	songs       *list.List
	currentSong *list.List
	playing     bool
	pause       bool
	pauseTime   int
	sync.RWMutex
	stopCh  chan struct{}
	pauseCh chan struct{}
	playCh  chan struct{}
}

type Song struct {
	name     string
	duration int
}

func (p *Playlist) AddSong(song Song) {
	go func() {
		p.Lock()
		p.songs.PushBack(song)
		p.Unlock()
	}()

}

func (p *Playlist) Play() {
	go process(p)
}

func process(p *Playlist) {
	fmt.Println("Playing")
	p.playing = true
	p.pause = false
	p.pauseTime = 0
	for {
		select {
		case <-p.stopCh:
			fmt.Println("Stop")
			p.playing = false
			p.pause = false
			return
		case <-p.pauseCh:
			fmt.Println("Pause")
			p.playing = false
			p.pause = true
			p.pauseTime = 3 // TODO: get time of playing
			return
		default:
			fmt.Println("Playing")
		}
	}
}

func main() {
	p := Playlist{
		songs:       list.New(),
		currentSong: nil,
		playing:     false,
		pause:       false,
		pauseTime:   0,
		stopCh:      make(chan struct{}),
		pauseCh:     make(chan struct{}),
		playCh:      make(chan struct{}),
	}

	p.songs.PushBack(Song{name: "song1", duration: 5})
	p.songs.PushBack(Song{name: "song2", duration: 5})
	p.songs.PushBack(Song{name: "song3", duration: 5})

	for e := p.songs.Front(); e != nil; e = e.Next() {
		go func(song Song) {
			fmt.Println(song.name)
			p.Play()
			time.Sleep(time.Duration(song.duration) * time.Second)
		}(e.Value.(Song))
	}

	go p.AddSong(Song{name: "song4", duration: 5})

	time.Sleep(10 * time.Second)
	fmt.Println(p)
	fmt.Println("Hello, playground")

}
