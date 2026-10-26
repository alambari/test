package main

import (
	"reflect"
	"testing"
)

func TestGO1(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{-1, 1}, 0},
		{[]int{-1, -7, -5}, -2},
		{[]int{1, 2, 1, 6}, 3},
		{[]int{2, 3}, 1},
		{[]int{0, 1, 2}, 3},
		{nil, 0},
	}
	for _, c := range cases {
		if got := GO1(c.in); got != c.want {
			t.Errorf("GO1(%v)=%d want %d", c.in, got, c.want)
		}
	}
}

func TestGO2(t *testing.T) {
	if got, err := GO2([]int{0, 1, 0}, 1000, 1010); got != 0 || err == nil || err.Error() != "not found" {
		t.Errorf("got %d, %v", got, err)
	}
	if got, err := GO2([]int{0, 1, 0}, 1259, 2601); got != 62952 || err != nil {
		t.Errorf("got %d, %v", got, err)
	}
}

func TestGO3(t *testing.T) {
	if got := GO3([]int{1, 2}, []int{1, 3}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("got %v", got)
	}
	if got := GO3([]int{1, 2, 2}, []int{1, 2, 4}); !reflect.DeepEqual(got, []int{2, 1}) {
		t.Errorf("got %v", got)
	}
}

func TestGO4(t *testing.T) {
	want := map[int]int{1: 1, 2: 2, 3: 3, 4: 1}
	if got := GO4(1223334); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
	if got := GO4(0); !reflect.DeepEqual(got, map[int]int{0: 1}) {
		t.Errorf("got %v", got)
	}
}

func TestGO5(t *testing.T) {
	if got := GO5(1, 2, 3, 4); got != 14 {
		t.Errorf("got %d", got)
	}
	if got := GO5(1, 2, 3, 4, 5); got != 39 {
		t.Errorf("got %d", got)
	}
	if got := GO5(); got != 0 {
		t.Errorf("got %d", got)
	}
}
