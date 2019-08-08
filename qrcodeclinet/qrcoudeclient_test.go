package qrcodeclinet

import (
	"fmt"
	"testing"
)

func TestClient_Encode(t *testing.T) {
	type args struct {
		info string
	}
	tests := []struct {
		name   string
		client *Client
		args   args
	}{
		{
			name:   "test1",
			client: New("localhost:5000"),
			args:   args{info: "12345"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.Encode(tt.args.info)
			if err != nil {
				t.Errorf("Client.Encode() error = %v", err)
				return
			}
			fmt.Printf("got :\n %v \n", got)
		})
	}
}
