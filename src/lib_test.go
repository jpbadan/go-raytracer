package main

import (
	"image/color"
	"math"
	"reflect"
	"testing"
)

func TestVec3(t *testing.T) {
	checkVec3Result := func(t testing.TB, got, want Vec3) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	checkResult := func(t testing.TB, got, want float32) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	v1 := Vec3{1, 2, 3}
	v2 := Vec3{4, 5, 6}
	var st float32 = 10.0

	t.Run("test sum of two Vec3", func(t *testing.T) {
		got := v1.sum(v2)
		want := Vec3{5, 7, 9}
		checkVec3Result(t, got, want)
	})

	t.Run("test multiplication Vec3 by t", func(t *testing.T) {
		got := v1.scale(st)
		want := Vec3{1 * st, 2 * st, 3 * st}
		checkVec3Result(t, got, want)
	})

	t.Run("test division of Vec3 by scalar", func(t *testing.T) {
		got := v1.div(st)
		want := Vec3{1 / st, 2 / st, 3 / st}
		checkVec3Result(t, got, want)
	})

	t.Run("test negation of Vec3", func(t *testing.T) {
		got := v1.neg()
		want := Vec3{-1, -2, -3}
		checkVec3Result(t, got, want)
	})

	t.Run("test length of Vec3", func(t *testing.T) {
		got := v1.length()
		var want float32 = float32(math.Pow(14, 0.5))
		checkResult(t, got, want)
	})

	t.Run("test length squared of Vec3", func(t *testing.T) {
		got := v1.lengthSquared()
		var want float32 = 14.0
		checkResult(t, got, want)
	})

	t.Run("test multiplication of two Vec3", func(t *testing.T) {
		got := v1.mult(v2)
		want := Vec3{4, 10, 18}
		checkVec3Result(t, got, want)
	})

	t.Run("test dot multiplication of two Vec3", func(t *testing.T) {
		got := v1.dot(v2)
		want := float32(32.0)
		checkResult(t, got, want)
	})

	t.Run("test cross multiplication of two Vec3", func(t *testing.T) {
		got := v1.cross(v2)
		want := Vec3{-3, 6, -3}
		checkVec3Result(t, got, want)
	})

	t.Run("test unit vector of Vec3", func(t *testing.T) {
		got := v1.unit()
		sqrt14 := float32(math.Pow(14, 0.5))
		want := Vec3{1 / sqrt14, 2 / sqrt14, 3 / sqrt14}
		checkVec3Result(t, got, want)
	})

	t.Run("test subtraction vector of two Vec3", func(t *testing.T) {
		got := v1.sub(v2)
		want := Vec3{-3, -3, -3}
		checkVec3Result(t, got, want)
	})

}

func TestColor(t *testing.T) {
	color1 := Color{0.0, 0.5, 1.0}
	t.Run("test RGBA conversion", func(t *testing.T) {
		got := color1.toRGBA()
		want := color.RGBA{0, 127, 255, 255}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
