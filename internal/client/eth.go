package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	Client *ethclient.Client
}

func NewEthClient(rpcURL string) (*EthClient, error) {
	c, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &EthClient{Client: c}, nil
}

func (e *EthClient) GetLatestBlockByNumber(ctx context.Context) (uint64, error) {
	header, err := e.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (e *EthClient) Close() {
	e.Client.Close()
}

func ToBlockNumber(n uint64) *big.Int {
	return new(big.Int).SetUint64(n)
}
