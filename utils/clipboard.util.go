package utils

import (
	"clip/database"
	"clip/types"
)

func SanitizeReturnData(dbData []database.Clipboard) []types.ReturnClipboard {
	returnData := make([]types.ReturnClipboard, len(dbData))
	for i, v := range dbData {
		returnData[i] = types.ReturnClipboard{
			Id:        int(v.ID),
			Datatype:  string(v.Datatype),
			Data:      string(v.Data),
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		}
	}
	return returnData
}
