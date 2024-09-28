package log

import (
	"fmt"

	proto "github.com/Egor123qwe/logs-storage/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	proto.LogsClient
	conn *grpc.ClientConn
}

func New(host string, port int) (*Client, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := &Client{
		LogsClient: proto.NewLogsClient(conn),
		conn:       conn,
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
