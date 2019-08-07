package generator

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		info string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				info: "123321",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.info)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			t.Logf("got: %v", got)
			fmt.Printf("got: %v\n", got)
		})
	}
}
