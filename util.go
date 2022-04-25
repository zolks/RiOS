package main

import (
	"sync"
	"time"
)

type item struct {
	value      Flow
	insertTime int64
}

type FlowMapTTL struct {
	m map[string]*item
	l sync.Mutex
}

func NewFlowMapTTL(initialLen int, ttl int) (m *FlowMapTTL) {
	m = &FlowMapTTL{m: make(map[string]*item, initialLen)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.insertTime > int64(ttl) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *FlowMapTTL) Len() int {
	return len(m.m)
}

func (m *FlowMapTTL) Put(k string, v Flow) {
	m.l.Lock()
	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.insertTime = time.Now().Unix()
	m.l.Unlock()
}

func (m *FlowMapTTL) Get(k string) (v Flow) {
	m.l.Lock()
	if it, ok := m.m[k]; ok {
		v = it.value
	}
	m.l.Unlock()
	return
}
