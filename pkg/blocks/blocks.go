package blocks

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type Block interface {
	AddAttributes(attr map[string]cty.Value)
	GetBlock() *hclwrite.Block
}
