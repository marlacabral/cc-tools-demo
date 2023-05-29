package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Token = assets.AssetType{
	Tag:         "token",
	Label:       "Token",
	Description: "Token",
	Props: []assets.AssetProp{
		{
			Tag:      "id",
			Label:    "ID",
			DataType: "string",
			Required: true,
			IsKey:    true,
		},
		{
			Tag:         "prop",
			Label:       "Proprietario",
			Description: "Proprietario",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "quant",
			Label:       "Quantidade",
			Description: "Quantidade",
			DataType:    "float64",
			Required:    true,
		},
	},
}
