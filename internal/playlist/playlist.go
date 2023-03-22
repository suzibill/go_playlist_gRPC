package main

import (
	"container/list"
	"context"
	"fmt"
	"time"
)

const (
	Play int = iota + 1
	Pause
	Next
	Prev
	AddSong
)

const (
	Tik time.Duration = 100 * time.Millisecond
)

type commands struct {
	command int
	song    song
}

type playlist struct {
	songs       *list.List
	currentSong *list.Element
}

type song struct {
	name     string
	duration time.Duration
}

func (p *playlist) addSong(s song) {
	p.songs.PushBack(s)
	if p.currentSong == nil {
		p.currentSong = p.songs.Back()
	}
}

func (p *playlist) Next() *song {
	if p.currentSong == nil || p.currentSong.Next() == nil {
		return nil
	}
	p.currentSong = p.currentSong.Next()
	tmp := p.currentSong.Value.(song)
	return &tmp
}

func (p *playlist) Prev() *song {
	if p.currentSong == nil || p.currentSong.Prev() == nil {
		return nil
	}
	p.currentSong = p.currentSong.Prev()
	tmp := p.currentSong.Value.(song)
	return &tmp
}

func (p *playlist) CurrentSong() *song {
	if p.currentSong == nil {
		return nil
	}
	tmp := p.currentSong.Value.(song)
	return &tmp
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	commandCh := make(chan commands)
	go doWork(ctx, commandCh)
	commandCh <- commands{command: Play}
	time.Sleep(3 * time.Second)
	commandCh <- commands{command: Next}
	commandCh <- commands{command: AddSong, song: song{name: "song4", duration: 1 * time.Second}}
	commandCh <- commands{command: Prev}
	commandCh <- commands{command: Prev}
	commandCh <- commands{command: Play}
	commandCh <- commands{command: Play}
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// graceful shutdown
	// нейминг переменных
	// http сервер
	// прологировать что песня началась
	// 100 милисекунд константа
	// вынести повторяющие блоки кода в отдельные функции, либо через анонимные функции, либо через общую структура

}

func doWork(ctx context.Context, commandCh <-chan commands) {
	p := playlist{songs: list.New(),
		currentSong: nil,
	}
	p.addSong(song{name: "aboba", duration: 1 * time.Second})
	p.addSong(song{name: "amogus", duration: 1 * time.Second})
	p.addSong(song{name: "shleppa", duration: 1 * time.Second})
	t := time.NewTicker(time.Second * 3)
	var duration int64
	//t.Stop()
	defer t.Stop()
	for {
		select {
		case cmd := <-commandCh:
			switch cmd.command {
			case Play:
				t.Reset(Tik)
				if duration == 0 {
					s := p.CurrentSong()
					if s != nil {
						duration = int64(s.duration / Tik)
					}
				}
			case Pause:
				t.Stop()
			case Next:
				s := p.Next()
				if s != nil {
					duration = int64(s.duration / Tik)
					t.Reset(Tik)
				}
			case Prev:
				s := p.Prev()
				if s != nil {
					duration = int64(s.duration / Tik)
					t.Reset(Tik)
				}
			case AddSong:
				p.addSong(cmd.song)
			}
		case <-ctx.Done():
			fmt.Println("The end by context")
			return
		case <-t.C:
			duration--
			fmt.Printf("tik ")
			if duration == 0 {
				s := p.Next()
				if s == nil {
					fmt.Println("The end by playlist")
					return
				}
				duration = int64(s.duration / Tik)
				t.Reset(Tik)
				fmt.Printf("next song %s\n", s.name)
			}
		}
	}
}
