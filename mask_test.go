package qr

import "testing"

func Test_getScoreRule1(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 97},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule1(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoreRule2(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule2(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoreRule3(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule3(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoreRule4(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule4(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoreRule5(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 0},
		{"Test 2", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, I, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, O, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, O, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule5(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoreRule6(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoreRule6(tt.args.canvas); got != tt.want {
				t.Errorf("getScoreRule6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScore(t *testing.T) {
	type args struct {
		canvas *[][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{&[][]bool{
			{I, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O},
			{O, I, O, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, O},
			{O, O, O, I, I, O, O, O, I, O, I, O, O, O, I, O, I, I, I, I, O},
			{I, O, I, I, I, I, I, I, O, O, I, O, O, I, I, O, I, O, O, O, O},
			{I, I, I, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, O, O, O},
			{I, O, I, I, O, I, I, I, O, O, I, O, I, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, I, O, I, O, O, O},
			{I, I, I, I, I, I, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, I, I, O, I, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, O, O, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, I, I, O, I, O, I, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, I, I, O},
		}}, 405},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScore(tt.args.canvas); got != tt.want {
				t.Errorf("getScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
