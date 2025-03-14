package mymath

import (
	"reflect"
	"testing"
)

func TestMean(t *testing.T) {
	m := Math{}
	got := m.Mean([]int{1, 2, 3, 4, 5})
	want := 3.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	gotMed := m.Median([]int{2, 3, 4})
	wantM := 3
	if gotMed != float64(wantM) {
		t.Errorf("got %2.f, want %d", gotMed, wantM)
	}

	gotMode := m.Mode([]int{1, 2, 2, 3, 3, 3, 4, 5})
	wantedMode := []int{3}
	if !reflect.DeepEqual(gotMode, wantedMode) {
		t.Errorf("got %d equal %d", gotMode, wantedMode)
	}
	empMode := m.Mode([]int{})
	wantempMode := []int{}
	if !reflect.DeepEqual(empMode, wantempMode) {
		t.Errorf("want %d got %d", empMode, wantempMode)
	}

	gotRange := m.Range([]int{3, 5, 6, 2, 1})
	wantRange := 5
	if gotRange != wantRange {
		t.Errorf("got %d want %d", gotRange, wantRange)
	}

	gotMax := m.Max([]int{8, 4, 3, 1})
	wantedMax := 8
	if gotMax != wantedMax {
		t.Errorf("got %d want %d", gotMax, wantedMax)
	}

	gotMin, err := m.Min([]int{4, 5, 2, 9})
	wantMin := 2
	if err != nil || gotMin != wantMin {
		t.Errorf("got %d, want %d", gotMin, wantMin)
	}
}
