package logic

import "testing"

// Out of bounds and equal values are being caught in the input validation phase
func TestCanBishopAttack(t *testing.T) {
	type args struct {
		bishopCoordinates []int
		targetCoordinates []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Can attack C1-G5",
			args: args{
				bishopCoordinates: []int{2, 0},
				targetCoordinates: []int{6, 4},
			},
			want: true,
		},
		{
			name: "Can't attack C2-C4",
			args: args{
				bishopCoordinates: []int{2, 1},
				targetCoordinates: []int{2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanBishopAttack(tt.args.bishopCoordinates, tt.args.targetCoordinates); got != tt.want {
				t.Errorf("CanBishopAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}
