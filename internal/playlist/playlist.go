package main

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"time"
)

const (
	Play int32 = iota + 1
	Pause
	Next
	Prev
	AddSong
)

const (
	TikDuration = 100 * time.Millisecond
)

type Commands struct {
	command int32
	song    Song
}

type Playlist struct {
	songs       *list.List
	currentSong *list.Element
}

type Song struct {
	id       int64
	name     string
	duration time.Duration
}

func (p *Playlist) AddSong(s Song) {
	p.songs.PushBack(s)
	if p.currentSong == nil {
		p.currentSong = p.songs.Back()
	}
}

func (p *Playlist) Next() *Song {
	if p.currentSong == nil || p.currentSong.Next() == nil {
		return nil
	}
	p.currentSong = p.currentSong.Next()
	tmp := p.currentSong.Value.(Song)
	return &tmp
}

func (p *Playlist) Prev() *Song {
	if p.currentSong == nil || p.currentSong.Prev() == nil {
		return nil
	}
	p.currentSong = p.currentSong.Prev()
	tmp := p.currentSong.Value.(Song)
	return &tmp
}

func (p *Playlist) CurrentSong() *Song {
	if p.currentSong == nil {
		return nil
	}
	tmp := p.currentSong.Value.(Song)
	return &tmp
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	commandCh := make(chan Commands)
	go doWork(ctx, commandCh)
	commandCh <- Commands{command: AddSong, song: Song{name: "Song1", duration: 7 * time.Second}}
	commandCh <- Commands{command: AddSong, song: Song{name: "Song2", duration: 7 * time.Second}}
	commandCh <- Commands{command: AddSong, song: Song{name: "Song3", duration: 7 * time.Second}}
	commandCh <- Commands{command: Play}
	time.Sleep(3 * time.Second)
	commandCh <- Commands{command: Next}
	time.Sleep(1 * time.Second)
	commandCh <- Commands{command: AddSong, song: Song{name: "Song4", duration: 1 * time.Second}}
	//commandCh <- Commands{command: Prev}

	time.Sleep(20 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// graceful shutdown
	// нейминг переменных ++
	// http сервер
	// прологировать что песня началась++
	// 100 милисекунд константа++
	// вынести повторяющие блоки кода в отдельные функции, либо через анонимные функции, либо через общую структура++

}

func doWork(ctx context.Context, commandCh <-chan Commands) {
	duration := int64(0)
	p := Playlist{songs: list.New()}
	s := p.CurrentSong()
	t := time.NewTicker(TikDuration)
	defer t.Stop()

	ChangeSong := func(s *Song) {
		if s != nil {
			duration = int64(s.duration / TikDuration)
			t.Reset(TikDuration)
		}
	}

	for {
		select {
		case cmd := <-commandCh:
			switch cmd.command {
			case Play:
				t.Reset(TikDuration)
				if duration == 0 {
					s = p.CurrentSong()
					ChangeSong(s)
					log.Println("Song started", s.name)
				} else {
					log.Println("Song resumed", s.name)
				}
			case Pause:
				t.Stop()
				log.Println("Pause")
			case Next:
				s = p.Next()
				ChangeSong(s)
				log.Println("Next")
			case Prev:
				s = p.Prev()
				ChangeSong(s)
				log.Println("Prev")
			case AddSong:
				p.AddSong(cmd.song)
				log.Printf("Song added %v", cmd.song.name)
			}
		case <-ctx.Done():
			log.Println("The end by context")
			return
		case <-t.C:
			duration--
			fmt.Printf("Tik ")
			if duration == 0 {
				log.Printf("end of song %v", s.name)
				s = p.Next()
				ChangeSong(s)
				if s == nil {
					log.Println("The end of Playlist")
					t.Stop()
				}
			}
		}
	}
}
