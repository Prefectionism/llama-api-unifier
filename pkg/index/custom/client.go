package custom

import (
	"context"
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_ index.Provider = (*Client)(nil)
)

type Client struct {
	url string

	client IndexClient
}

type Option func(*Client