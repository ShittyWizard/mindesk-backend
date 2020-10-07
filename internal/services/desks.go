package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"min-desk-backend/internal/models"
)

type DeskInfo struct {
	Id      primitive.ObjectID `json:"_id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Columns []ColumnInfo       `json:"columns,omitempty"`
}

type ColumnInfo struct {
	Id    primitive.ObjectID `json:"_id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Cards []*models.Card     `json:"cards,omitempty"`
}

func GetDeskInfoByDeskId(deskId primitive.ObjectID) DeskInfo {
	desk, _ := models.GetDeskById(deskId)
	var deskInfo DeskInfo
	deskInfo.Id = desk.Id
	deskInfo.Name = desk.Name
	columns := models.GetColumnsByDeskId(desk.Id)
	var columnInfos []ColumnInfo
	for _, column := range columns {
		cards := models.GetAllCardsByColumnId(column.Id)
		columnInfo := ColumnInfo{
			Id:    column.Id,
			Name:  column.Name,
			Cards: cards,
		}
		columnInfos = append(columnInfos, columnInfo)
	}
	deskInfo.Columns = columnInfos

	return deskInfo
}
