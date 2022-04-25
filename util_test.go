package main

import (
	"testing"
	"time"
)

func TestFlowMapTTL(t *testing.T) {

	flowMap := NewFlowMapTTL(10, 2)

	flowMap.Put("0800997777", Flow{Id: 11})
	flowMap.Put("0800997778", Flow{Id: 12})
	flowMap.Put("0800997779", Flow{Id: 13})
	flowMap.Put("0800997710", Flow{Id: 14})
	flowMap.Put("0800997711", Flow{Id: 15})
	time.Sleep(3 * time.Second)
	flowMap.Put("0800997712", Flow{Id: 16})
	flowMap.Put("0800997713", Flow{Id: 17})
	flowMap.Put("0800997714", Flow{Id: 18})
	flowMap.Put("0800997715", Flow{Id: 19})
	flowMap.Put("0800997716", Flow{Id: 20})

	flow1 := flowMap.Get("0800997777")
	if flow1.Id != 0 {
		t.Error("id:0800997777 FOUND!")
	}

	flow2 := flowMap.Get("0800997712")
	if flow2.Id != 16 {
		t.Error("id:0800997712 NOT FOUND!")
	}

}
