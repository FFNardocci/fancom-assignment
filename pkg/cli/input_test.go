package cli

import (
	"reflect"
	"testing"
)

func TestValidateAndTranlsateInputString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Success [0,0]",
			args: args{
				input: "a1",
			},
			want:    []int{0, 0},
			wantErr: false,
		},
		{
			name: "Success [7,7]",
			args: args{
				input: "h8",
			},
			want:    []int{7, 7},
			wantErr: false,
		},
		{
			name: "Fail: input string too long",
			args: args{
				input: "a12",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail: invalid input string",
			args: args{
				input: "bb",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail: input string out of 8x8 range 1",
			args: args{
				input: "h9",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail: input string out of 8x8 range 2",
			args: args{
				input: "i1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateAndTranlsateInputString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAndTranlsateInputString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAndTranlsateInputString() = %v, want %v", got, tt.want)
			}
		})
	}
}
