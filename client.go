package switchboardv2

import (
	"github.com/gagliardetto/solana-go/rpc"
	"go.uber.org/zap"
)

// Client interacts with Pyth via Solana's JSON-RPC API.
//
// Do not instantiate Client directly, use NewClient instead.
type Client struct {
	Env          Env
	RPC          *rpc.Client
	WebSocketURL string
	Log          *zap.Logger

}

func NewClient(env Env, rpcURL string, wsURL string) *Client {
	return &Client{
		Env:          env,
		RPC:          rpc.New(rpcURL),
		WebSocketURL: wsURL,
		Log:          zap.NewNop(),

	}
}
