package main

import (
	"log"
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

func NewFlowMapTTL(initialLen int, ttl int, cleanInterval time.Duration) (m *FlowMapTTL) {
	m = &FlowMapTTL{m: make(map[string]*item, initialLen)}
	go func() {
		for now := range time.Tick(cleanInterval) {
			m.l.Lock()
			var newest int64
			oldest := now.Unix()
			for k, v := range m.m {
				if now.Unix()-v.insertTime > int64(ttl) {
					delete(m.m, k)
				} else if v.insertTime < oldest {
					oldest = v.insertTime
				}
				if v.insertTime > newest {
					newest = v.insertTime
				}
			}
			if len(m.m) == 0 {
				log.Printf("FlowMap statistics: active(0) newest(0) oldest(0) ttl(%v).", ttl)
			} else {
				log.Printf("FlowMap statistics: active(%v) newest(%v) oldest(%v) ttl(%v).", len(m.m), time.Unix(newest, 0), time.Unix(oldest, 0), ttl)
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
