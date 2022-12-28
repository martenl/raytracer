package geometry

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "Empty matrix",
			args: args{height: 0, width: 0},
			want: Matrix{height: 0, width: 0, data: make([][]float64, 0)},
		},
		{
			name: "Non-empty matrix",
			args: args{height: 2, width: 3},
			want: Matrix{height: 2, width: 3, data: func() [][]float64 {
				myData := make([][]float64, 2)
				for i := 0; i < 2; i++ {
					myData[i] = make([]float64, 3)
				}
				return myData
			}()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.height, tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateWithValues(t *testing.T) {
	type args struct {
		height int
		width  int
		values []float64
	}
	tests := []struct {
		name  string
		args  args
		wantM Matrix
	}{
		{
			name:  "Empty Matrix",
			args:  args{width: 0, height: 0},
			wantM: Create(0, 0),
		},
		{
			name: "Non-empty square Matrix",
			args: args{width: 2, height: 2, values: []float64{1, 2, 3, 4}},
			wantM: func() Matrix {
				m := Create(2, 2)
				values := []float64{1, 2, 3, 4}
				idx := 0
				for i := 0; i < 2; i++ {
					for j := 0; j < 2; j++ {
						m.Set(j, i, values[idx])
						idx++
					}
				}
				return m
			}(),
		},
		{
			name: "Non-empty non-square Matrix",
			args: args{width: 2, height: 3, values: []float64{1, 2, 3, 4, 5, 6}},
			wantM: func() Matrix {
				m := Create(3, 2)
				values := []float64{1, 2, 3, 4, 5, 6}
				idx := 0
				for i := 0; i < 3; i++ {
					for j := 0; j < 2; j++ {
						m.Set(j, i, values[idx])
						idx++
					}
				}
				return m
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := CreateWithValues(tt.args.height, tt.args.width, tt.args.values...); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("CreateWithValues() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestIdentity(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantM Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := Identity(tt.args.n); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("Identity() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		m2 Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		{
			name: "Add empty Matrix",
			fields: fields{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
			args: args{
				Matrix{
					height: 0,
					width:  0,
					data:   [][]float64{},
				}},
			want: Matrix{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
		},
		{
			name: "Add non-empty Matrix",
			fields: fields{
				height: 2,
				width:  2,
				data:   [][]float64{{1,2},{3,4}},
			},
			args: args{
				Matrix{
					height: 2,
					width:  2,
					data:   [][]float64{{4,3},{2,1}},
				}},
			want: Matrix{
				height: 2,
				width:  2,
				data:   [][]float64{{5,5},{5,5}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Add(tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_At(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		x int
		y int
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
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.At(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cofactor(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		row    int
		column int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCofactor float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotCofactor := m.Cofactor(tt.args.row, tt.args.column); gotCofactor != tt.wantCofactor {
				t.Errorf("Cofactor() = %v, want %v", gotCofactor, tt.wantCofactor)
			}
		})
	}
}

func TestMatrix_Determinant(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name            string
		fields          fields
		wantDeterminant float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotDeterminant := m.Determinant(); gotDeterminant != tt.wantDeterminant {
				t.Errorf("Determinant() = %v, want %v", gotDeterminant, tt.wantDeterminant)
			}
		})
	}
}

func TestMatrix_Dims(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name       string
		fields     fields
		wantHeight int
		wantWidth  int
	}{
		{
			name:       "Empty matrix dimensions",
			fields:     fields{height: 0, width: 0},
			wantHeight: 0,
			wantWidth:  0,
		},
		{
			name:       "Non-empty matrix dimensions",
			fields:     fields{height: 2, width: 7},
			wantHeight: 2,
			wantWidth:  7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			gotHeight, gotWidth := m.Dims()
			if gotHeight != tt.wantHeight {
				t.Errorf("Dims() gotHeight = %v, want %v", gotHeight, tt.wantHeight)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("Dims() gotWidth = %v, want %v", gotWidth, tt.wantWidth)
			}
		})
	}
}

func TestMatrix_Equals(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		m2 Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Empty matrix equals empty matrix",
			fields: fields{height: 0, width: 0, data: [][]float64{}},
			args: args{
				m2: Matrix{
					height: 0,
					width:  0,
					data:   [][]float64{},
				},
			},
			want: true,
		},
		{
			name:   "Empty matrix doesn't equal non-empty matrix",
			fields: fields{height: 0, width: 0},
			args: args{
				m2: Matrix{
					height: 1,
					width:  1,
					data:   [][]float64{{5}},
				},
			},
			want: false,
		},
		{
			name:   "Non-empty matrix equals the same matrix",
			fields: fields{height: 2, width: 2, data: [][]float64{{1, 2}, {3, 4}}},
			args: args{
				m2: Matrix{
					height: 2,
					width:  2,
					data:   [][]float64{{1, 2}, {3, 4}},
				},
			},
			want: true,
		},
		{
			name:   "Non-empty matrix doesn't equal matrix with different data",
			fields: fields{height: 2, width: 2, data: [][]float64{{1, 2}, {3, 4}}},
			args: args{
				m2: Matrix{
					height: 2,
					width:  2,
					data:   [][]float64{{2, 1}, {4, 3}},
				},
			},
			want: false,
		},
		{
			name:   "Non-empty matrix doesn't equal matrix with different dimensions",
			fields: fields{height: 2, width: 2, data: [][]float64{{1, 2}, {3, 4}}},
			args: args{
				m2: Matrix{
					height: 1,
					width:  2,
					data:   [][]float64{{1, 2}},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Equals(tt.args.m2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Height(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name       string
		fields     fields
		wantHeight int
	}{
		{
			name:       "Empty matrix height",
			fields:     fields{height: 0, width: 0},
			wantHeight: 0,
		},
		{
			name:       "Non-empty matrix height",
			fields:     fields{height: 2, width: 7},
			wantHeight: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotHeight := m.Height(); gotHeight != tt.wantHeight {
				t.Errorf("Height() = %v, want %v", gotHeight, tt.wantHeight)
			}
		})
	}
}

func TestMatrix_Inverse(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Inverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Minor(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		row    int
		column int
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
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Minor(tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("Minor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Multiply(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		m2 Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM3 Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotM3 := m.Multiply(tt.args.m2); !reflect.DeepEqual(gotM3, tt.wantM3) {
				t.Errorf("Multiply() = %v, want %v", gotM3, tt.wantM3)
			}
		})
	}
}

func TestMatrix_Print(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			m.Print()
		})
	}
}

func TestMatrix_RotateXaxis(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		rad float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.RotateXaxis(tt.args.rad); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateXaxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_RotateYaxis(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		rad float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.RotateYaxis(tt.args.rad); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateYaxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_RotateZaxis(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		rad float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.RotateZaxis(tt.args.rad); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateZaxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_ScalarMultiply(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		scalar float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{

		{
			name: "Multipy empty Matrix with scalar",
			fields: fields{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
			args: args{
				scalar: 2},
			want: Matrix{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
		},
		{
			name: "Multipy non-empty Matrix with scalar",
			fields: fields{
				height: 2,
				width:  2,
				data:   [][]float64{{1, 2}, {3, 4}},
			},
			args: args{scalar: 2.0},
			want: Matrix{
				height: 2,
				width:  2,
				data:   [][]float64{{2, 4}, {6, 8}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.ScalarMultiply(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalarMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Scale(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		{
			name:   "Identity Matrix scaling",
			fields: fields{height: 4, width: 4, data: [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
			args:   args{x: 1, y: 1, z: 1},
			want:   Matrix{height: 4, width: 4, data: [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
		},
		{
			name:   "Empty Vector scaling",
			fields: fields{height: 4, width: 1, data: [][]float64{{0}, {0}, {0}, {0}}},
			args:   args{x: 2, y: 3, z: 4},
			want:   Matrix{height: 4, width: 1, data: [][]float64{{0}, {0}, {0}, {0}}},
		},
		{
			name:   "Vector scaling",
			fields: fields{height: 4, width: 1, data: [][]float64{{1}, {1}, {1}, {0}}},
			args:   args{x: 2, y: 3, z: 4},
			want:   Matrix{height: 4, width: 1, data: [][]float64{{2}, {3}, {4}, {0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Scale(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		x     int
		y     int
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Set(tt.args.x, tt.args.y, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Shear(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		xy float64
		xz float64
		yx float64
		yz float64
		zx float64
		zy float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		{
			name: "Shearing",
			fields: fields{
				height: 4,
				width:  1,
				data:   [][]float64{{2}, {3}, {4}, {1}}},
			args: args{
				xy: 1,
				xz: 0,
				yx: 0,
				yz: 0,
				zx: 0,
				zy: 0,
			},
			want: Matrix{
				height: 4,
				width:  1,
				data:   [][]float64{{5}, {3}, {4}, {1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Shear(tt.args.xy, tt.args.xz, tt.args.yx, tt.args.yz, tt.args.zx, tt.args.zy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Submatrix(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		row    int
		column int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantSubmatrix Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotSubmatrix := m.Submatrix(tt.args.row, tt.args.column); !reflect.DeepEqual(gotSubmatrix, tt.wantSubmatrix) {
				t.Errorf("Submatrix() = %v, want %v", gotSubmatrix, tt.wantSubmatrix)
			}
		})
	}
}

func TestMatrix_Subtract(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		m2 Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{

		{
			name: "Subtract empty Matrix",
			fields: fields{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
			args: args{
				Matrix{
					height: 0,
					width:  0,
					data:   [][]float64{},
				}},
			want: Matrix{
				height: 0,
				width:  0,
				data:   [][]float64{},
			},
		},
		{
			name: "Subtract non-empty Matrix",
			fields: fields{
				height: 2,
				width:  2,
				data:   [][]float64{{5, 5}, {5, 5}},
			},
			args: args{
				Matrix{
					height: 2,
					width:  2,
					data:   [][]float64{{4, 3}, {2, 1}},
				}},
			want: Matrix{
				height: 2,
				width:  2,
				data:   [][]float64{{1, 2}, {3, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Subtract(tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Translate(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		{
			name: "Translate a point",
			fields: fields{
				height: 4,
				width:  1,
				data:   [][]float64{{-3}, {4}, {5}, {1}}, //-3, 4, 5
			},
			args: args{
				x: 5,
				y: -3,
				z: 2,
			},
			want: Matrix{
				height: 4,
				width:  1,
				data:   [][]float64{{2}, {1}, {7}, {1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.Translate(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		wantM2 Matrix
	}{
		{
			name:   "Transpose of empty matrix",
			fields: fields{height: 0, width: 0, data: [][]float64{}},
			wantM2: Matrix{height: 0, width: 0, data: [][]float64{}},
		},
		{
			name:   "Transpose of non-empty matrix",
			fields: fields{height: 4, width: 4, data: [][]float64{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}}},
			wantM2: Matrix{height: 4, width: 4, data: [][]float64{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}},
		},
		{
			name:   "Transpose of identity matrix",
			fields: fields{height: 4, width: 4, data: [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
			wantM2: Identity(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotM2 := m.Transpose(); !reflect.DeepEqual(gotM2, tt.wantM2) {
				t.Errorf("Transpose() = %v, want %v", gotM2, tt.wantM2)
			}
		})
	}
}

func TestMatrix_Width(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	tests := []struct {
		name      string
		fields    fields
		wantWidth int
	}{
		{
			name:      "Empty matrix width",
			fields:    fields{height: 0, width: 0},
			wantWidth: 0,
		},
		{
			name:      "Non-empty matrix width",
			fields:    fields{height: 2, width: 7},
			wantWidth: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotWidth := m.Width(); gotWidth != tt.wantWidth {
				t.Errorf("Width() = %v, want %v", gotWidth, tt.wantWidth)
			}
		})
	}
}

func TestMatrix_rotate(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		rad  float64
		axis rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if got := m.rotate(tt.args.rad, tt.args.axis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_submatrix(t *testing.T) {
	type fields struct {
		height int
		width  int
		data   [][]float64
	}
	type args struct {
		row    int
		column int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantSubmatrix Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				height: tt.fields.height,
				width:  tt.fields.width,
				data:   tt.fields.data,
			}
			if gotSubmatrix := m.submatrix(tt.args.row, tt.args.column); !reflect.DeepEqual(gotSubmatrix, tt.wantSubmatrix) {
				t.Errorf("submatrix() = %v, want %v", gotSubmatrix, tt.wantSubmatrix)
			}
		})
	}
}
