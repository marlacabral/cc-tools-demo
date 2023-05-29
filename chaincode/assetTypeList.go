package main

import (
	"github.com/goledgerdev/cc-tools-demo/chaincode/assettypes"
	"github.com/goledgerdev/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Proprietario,
	assettypes.Token,
	assettypes.Library,
	assettypes.Secret,
}
