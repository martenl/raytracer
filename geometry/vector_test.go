package geometry

import (
	"reflect"
	"testing"
)

func TestCreateVector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			args: args{x: 1, y: 2, z: 3},
			want: Vector{
				data: CreateWithValues(4, 1, 1, 2, 3, 0),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateVector(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Cross(t *testing.T) {
	type fields struct {
		data Matrix
	}
	type args struct {
		v2 Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Cross(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Dot(t *testing.T) {
	type fields struct {
		data Matrix
	}
	type args struct {
		v2 Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Length(t *testing.T) {
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
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Normalize(t *testing.T) {
	type fields struct {
		data Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_X(t *testing.T) {
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
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.X(); got != tt.want {
				t.Errorf("X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Y(t *testing.T) {
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
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Y(); got != tt.want {
				t.Errorf("Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Z(t *testing.T) {
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
			v := Vector{
				data: tt.fields.data,
			}
			if got := v.Z(); got != tt.want {
				t.Errorf("Z() = %v, want %v", got, tt.want)
			}
		})
	}
}
