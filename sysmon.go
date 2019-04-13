package main

type SysMon struct {
}

func NewSysMon() *SysMon {
	sm := &SysMon{}

	return sm
}

/* A go routine which will monitor the system and poke listeners */
func (sm *SysMon) Run() {

}
