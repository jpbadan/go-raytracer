package main

import (
	"reflect"
	"testing"
)

func TestRay(t *testing.T) {
	checkResult := func(t testing.TB, got, want Vec3) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	r := Ray{Point3{0, 0, 0}, Vec3{1, 1, 1}}

	t.Run("Test if Ray returns it's origin", func(t *testing.T) {
		got := r.origin()
		want := Vec3{0, 0, 0}
		checkResult(t, got, want)
	})

	t.Run("Test if Ray returns it's direction", func(t *testing.T) {
		got := r.direction()
		want := Vec3{1, 1, 1}
		checkResult(t, got, want)
	})

	t.Run("Test if Ray returns it's equation result", func(t *testing.T) {
		got := r.at(10)
		want := Vec3{10, 10, 10}
		checkResult(t, got, want)
	})

}
