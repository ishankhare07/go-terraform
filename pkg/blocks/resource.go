package blocks

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type _resource struct {
	Type         string `hcl:"type,label"`
	ResourceName string `hcl:",label"`
}

type Resource struct {
	HclBlock *hclwrite.Block
}

func (r *Resource) AddAttributes(attrs map[string]cty.Value) {
	for k, v := range attrs {
		r.HclBlock.Body().SetAttributeValue(k, v)
	}
}

func (r *Resource) GetBlock() *hclwrite.Block {
	return r.HclBlock
}

func NewResource(t, resourceName string) Block {
	_r := _resource{
		Type:         t,
		ResourceName: resourceName,
	}

	return &Resource{
		HclBlock: gohcl.EncodeAsBlock(_r, "resource"),
	}
}
