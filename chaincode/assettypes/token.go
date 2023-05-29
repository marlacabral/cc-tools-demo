package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Token = assets.AssetType{
	Tag:         "token",
	Label:       "Token",
	Description: "Token",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "ID",
			DataType: "string",
		},
		{
			// Composite Key
			Tag:         "prop",
			Label:       "Proprietario",
			Description: "Proprietario",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "quant",
			Label:       "Quant",
			Description: "ID",
			DataType:    "string",
			Required:    true,
		},
	},
}
