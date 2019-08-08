package qrcodeclinet

import (
	"context"
	"errors"
	"time"

	pb "github.com/Basic-Components/qrcodegenerator/qrcodegeneratordeclare"

	grpc "google.golang.org/grpc"
)

// Client jwt的客户端类型
type Client struct {
	Address string
}

// New 创建客户端对象
func New(address string) *Client {
	return &Client{Address: address}
}

// Encode 获取二维码图片的base64字符串
func (client *Client) Encode(info string) (string, error) {
	conn, err := grpc.Dial(client.Address, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := pb.NewQrServiceClient(conn)
	// 设置请求上下文的过期时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := c.Encode(ctx, &pb.Request{Data: info})
	if err != nil {
		return "", err
	}
	if rs.Status.Status == pb.StatusData_ERROR {
		return "", errors.New(rs.Status.Msg)
	}
	return rs.Data, nil
}
