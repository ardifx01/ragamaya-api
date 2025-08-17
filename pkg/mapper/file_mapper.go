package mapper

import (
	"ragamaya-api/api/storages/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapFilesInputToModel(input dto.FilesInput) models.Files {
	var data models.Files
	mapstructure.Decode(input, &data)
	return data
}

func MapFilesModelToOutput(model models.Files) dto.FilesOutput {
	var output dto.FilesOutput
	mapstructure.Decode(model, &output)
	output.CreatedAt = model.CreatedAt
	return output
}
