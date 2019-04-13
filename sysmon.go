package main

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type SysMon struct {
}

func NewSysMon() *SysMon {
	sm := &SysMon{}

	return sm
}

/* A go routine which will monitor the system and poke listeners */
func (sm *SysMon) Run() {

	for {
		time.Sleep(1 * time.Second)

		s := time.Now().String()
		gWindow.SetTime(s)

		p, err := cpu.Percent(0, true)
		if err == nil {
			//log.Infof("CPU: %v", p)
			gWindow.SetCPU(p)
		}

	}
}

func (sm *SysMon) NumCPUs() int {
	p, err := cpu.Percent(1*time.Millisecond, true)

	if err != nil {
		log.Errorf("Didn't get any CPUs: %v", err)
		return 0
	}

	return len(p)
}
