package services

import (
	"math/rand"
	"testing"
	"time"
)

func TestSh(t *testing.T) {
	check := make(map[string]string)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 100000; i++ {
		data := RandStringBytes(r1)
		for _, v := range data {
			if _, ok := check[string(v)]; !ok {
				check[string(v)] = "ok"
			}
		}
		if len(data) != 10 {
			t.Errorf("Output length is not what expected")
		}
	}
	if len(letterBytes) != len(check) {
		t.Errorf("The algorythm is not what expected")
	}
}
