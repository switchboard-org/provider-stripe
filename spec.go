package main

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type ProviderSchema interface {
	Value() cty.Value
	Parse(hcl.Body) hcldec.ObjectSpec
}

type Block interface {
}

type Attribute interface {
}

type BlockSpec struct {
	Labels       []string
	NestedBlocks []Block
	Attributes   []Attribute
}

type SingleBlockSpec struct {
	NestedBlocks  []Block
	AttributeSpec []Attribute
}

type AttributeSpec struct {
	Labels   []string
	Value    cty.Value
	Type     cty.Type
	Optional bool
}
