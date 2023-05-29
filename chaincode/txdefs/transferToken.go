package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// POST Method
var TransferToken = tx.Transaction{
	Tag:         "transferToken",
	Label:       "Transfer Token",
	Description: "transfer",
	Method:      "PUT",

	Args: []tx.Argument{
		{
			Tag:         "token",
			Label:       "token",
			Description: "token",
			DataType:    "->token",
			Required:    true,
		},
		{
			Tag:         "id",
			Label:       "id",
			Description: "id",
			DataType:    "string",
		},
		{
			Tag:         "quantidade",
			Label:       "quantidade",
			Description: "quantidade",
			DataType:    "number",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		tokenKey, ok := req["token"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter token must be an asset")
		}
		destino, ok := req["destino"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter destino must be an string")
		}
		id, ok := req["id"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter id must be an string")
		}
		quantidade, ok := req["quantidade"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter quanti must be a number")
		}
		tokenAsset, err := tokenKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		tokenMap := (map[string]interface{})(*tokenAsset)

		tokenMap["quantidade"] = tokenMap["quantidade"].(float64) - quantidade

		tokenMap, err = tokenAsset.Update(stub, tokenMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		destMap := make(map[string]interface{})
		destMap["@assetType"] = "token"
		destMap["id"] = id
		destMap["proprietario"] = destino
		destMap["quantidade"] = quantidade

		destAsset, err := assets.NewAsset(destMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new dest on channel
		_, err = destAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving destino on blockchain")
		}

		// Marshal asset back to JSON format
		tokenJSON, nerr := json.Marshal(tokenMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return tokenJSON, nil
	},
}
