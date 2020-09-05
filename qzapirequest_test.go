package qzgo

import (
	"fmt"
	"testing"
	"time"
)

func TestQzgo(t *testing.T) {
	client, err := NewClientLogin("ncepu", "20189999999", "666666")
	if err != nil {
		panic(err)
	}

	current, err := client.GetCurrentTime(time.Now().Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	fmt.Println(current)

	current, err = client.GetCurrentTime("")
	if err != nil {
		panic(err)
	}
	fmt.Println(current)

	kb, err := client.GetKbcxAzc("20189999999", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(kb)
}
