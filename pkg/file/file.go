package file

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/ishankhare07/go-terraform/pkg/blocks"
)

type File interface {
	AddBlock(blocks.Block)
	Output() []byte
}

type HCLFile struct {
	file *hclwrite.File
}

func (h *HCLFile) AddBlock(b blocks.Block) {
	h.file.Body().AppendBlock(b.GetBlock())
}

func (h *HCLFile) Output() []byte {
	return h.file.Bytes()
}

func NewFile() File {
	return &HCLFile{
		file: hclwrite.NewEmptyFile(),
	}
}
