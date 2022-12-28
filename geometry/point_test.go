package geometry

import (
	"reflect"
	"testing"
)

func TestCreatePoint(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name  string
		args  args
		wantP Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := CreatePoint(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("CreatePoint() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	type fields struct {
		data Matrix
	}
	type args struct {
		v Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				data: tt.fields.data,
			}
			if got := p.Add(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Multiply(t *testing.T) {
	type fields struct {
		data Matrix
	}
	type args struct {
		m Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				data: tt.fields.data,
			}
			if got := p.Multiply(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_X(t *testing.T) {
	type fields struct {
		data Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				data: tt.fields.data,
			}
			if got := p.X(); got != tt.want {
				t.Errorf("X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Y(t *testing.T) {
	type fields struct {
		data Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				data: tt.fields.data,
			}
			if got := p.Y(); got != tt.want {
				t.Errorf("Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Z(t *testing.T) {
	type fields struct {
		data Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				data: tt.fields.data,
			}
			if got := p.Z(); got != tt.want {
				t.Errorf("Z() = %v, want %v", got, tt.want)
			}
		})
	}
}
