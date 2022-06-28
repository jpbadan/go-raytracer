package entities

import (
	"go-raytracer/lib"
	"reflect"
	"testing"
)

func TestRay(t *testing.T) {
	checkResult := func(t testing.TB, got, want lib.Vec) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	r := Ray{lib.Point{X: 0, Y: 0, Z: 0}, lib.Vec{X: 1, Y: 1, Z: 1}}

	t.Run("Test if Ray returns it's origin", func(t *testing.T) {
		got := r.Origin()
		want := lib.Vec{X: 0, Y: 0, Z: 0}
		checkResult(t, got, want)
	})

	t.Run("Test if Ray returns it's direction", func(t *testing.T) {
		got := r.Direction()
		want := lib.Vec{X: 1, Y: 1, Z: 1}
		checkResult(t, got, want)
	})

	t.Run("Test if Ray returns it's equation result", func(t *testing.T) {
		got := r.At(10)
		want := lib.Vec{X: 10, Y: 10, Z: 10}
		checkResult(t, got, want)
	})

}
