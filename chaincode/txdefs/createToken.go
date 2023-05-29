package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// POST Method
var CreateToken = tx.Transaction{
	Tag:         "createToken",
	Label:       "Create Token",
	Description: "Create a Token",
	Method:      "POST",
	Callers:     []string{"$org3MSP", "$orgMSP"},

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "ID",
			Description: "ID",
			DataType:    "string",
			Required:    true,
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
			DataType:    "float64", // Change the data type to float64 for quant
			Required:    true,      // Make quant a required argument
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)
		prop, _ := req["prop"].(string)
		quant, _ := req["quant"].(float64)

		if quant <= 0.0 {
			return nil, errors.WrapError(nil, "Quantity must be greater than zero")
		}

		tokenMap := make(map[string]interface{})
		tokenMap["@assetType"] = "token"
		tokenMap["id"] = id
		tokenMap["prop"] = prop
		tokenMap["quant"] = quant

		tokenAsset, err := assets.NewAsset(tokenMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")

		}
		_, err = tokenAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		tokenJSON, nerr := json.Marshal(tokenAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return tokenJSON, nil
	},
}
