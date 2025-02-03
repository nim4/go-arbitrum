package arbitrum

import (
	"context"

	"github.com/nim4/go-arbitrum/arbitrum_types"
	"github.com/nim4/go-arbitrum/core"
	"github.com/nim4/go-arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
