package main

import (
	"fmt"
	"sync"
	"time"
)

type someData struct {
	stopCh  chan struct{}
	pauseCh chan struct{}
	playCh  chan struct{}
	counter int
	dataP   int
	data    [][]int
	sync.RWMutex
}

func (d *someData) Stop() {
	fmt.Println("Press stop")
	d.stopCh <- struct{}{}
}

func (d *someData) Pause() {
	fmt.Println("Press pause")
	d.pauseCh <- struct{}{}
}

func (d *someData) Play() {
	fmt.Println("Press play")
	go doWork(d)
}

func (d *someData) Next() {
	fmt.Println("Press next")
	d.dataP++
	if d.dataP == len(d.data) {
		d.dataP = 0
	}
}

func (d *someData) Prev() {
	fmt.Println("Press prev")
	d.dataP--
	if d.dataP == -1 {
		d.dataP = len(d.data) - 1
	}
}

func (d *someData) AddDate(data []int) {
	fmt.Println("Press add")
	go func() {
		d.Lock()
		d.data = append(d.data, data)
		d.Unlock()
	}()
}

func main() {
	d := someData{stopCh: make(chan struct{}), pauseCh: make(chan struct{}), playCh: make(chan struct{}), data: [][]int{{1, 2}, {3, 4}}}
	d.Play()
	time.Sleep(4 * time.Second)
	d.Pause()
	time.Sleep(2 * time.Second)
	d.Play()
	d.AddDate([]int{228, 1337})
	time.Sleep(10 * time.Second)
	d.Stop()
	time.Sleep(2 * time.Second)
}

func doWork(d *someData) {
	for {
		select {
		case <-d.stopCh:
			fmt.Println("stop")
			d.counter = 0
			d.dataP = 0
			return
		case <-d.pauseCh:
			fmt.Println("pause")
			return
		default:
			println(d.data[d.dataP][d.counter%len(d.data[d.dataP])])
			d.counter++
			if d.counter%len(d.data[d.dataP]) == 0 {
				fmt.Println("auto next")
				d.dataP++
				if d.dataP == len(d.data) {
					d.dataP = 0
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
