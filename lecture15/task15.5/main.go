package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type weather struct {
	currentTemp string
	mu          sync.RWMutex
}

func (mi *weather) ReadTemp() string {
	mi.mu.RLock()
	defer mi.mu.RUnlock()
	return mi.currentTemp
}

func (mi *weather) ChangeTemp(v string) {
	mi.mu.Lock()
	defer mi.mu.Unlock()
	mi.currentTemp = v
}

func main() {
	w := weather{}
	w.mu = sync.RWMutex{}
	for j := 0; j < 100; j++ {
		go func() {
			w.ChangeTemp(strconv.Itoa(j))
		}()
		go func() {
			fmt.Println(w.ReadTemp())
		}()
	}
	time.Sleep(10 * time.Second)
}
