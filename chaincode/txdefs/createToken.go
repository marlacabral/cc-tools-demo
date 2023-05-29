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
	Callers:     []string{"$org3MSP", "$org2MSP", "$org1MSP", "$orgMSP"},

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
			Label:       "Quant",
			Description: "ID",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)
		prop, _ := req["prop"].(string)
		quant, _ := req["quant"].(string)

		tokensMap := make(map[string]interface{})
		tokensMap["@assetType"] = "token"
		tokensMap["id"] = id
		tokensMap["prop"] = prop
		tokensMap["quant"] = quant

		tokensAsset, err := assets.NewAsset(tokensMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}
		_, err = tokensAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		tokenJSON, nerr := json.Marshal(tokensAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return tokenJSON, nil
	}}
