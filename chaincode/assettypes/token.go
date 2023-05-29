package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a book
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
			Writers:  []string{`org3MSP`, `org2MSP`, `org1MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			// Composite Key
			Tag:         "prop",
			Label:       "Proprietario",
			Description: "Proprietario",
			DataType:    "string",
			Writers:     []string{`org2MSP`, `org1MSP`, "org3MSP", "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Required:    true,
		},
		{
			Tag:         "quant",
			Label:       "Quant",
			Description: "ID",
			DataType:    "string",
			Writers:     []string{`org2MSP`, `org1MSP`, `org3MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Required:    true,
		},
	},
}
