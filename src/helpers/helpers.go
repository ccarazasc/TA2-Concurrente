package helpers

import (
	"encoding/csv"
	"go-api/src/models"
	"io"
	"log"
	"net/http"
)

func ReadCSVFromUrl(url string) (models.Data, error) {
	var data models.Data
	csvFile, _ := http.Get(url)
	reader := csv.NewReader(csvFile.Body)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.Comma = ','
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if line[8]== "Aprobado" || line[8]== "Desaprobado" {
			data.Data = append(data.Data, models.RowData{
				Id:             line[0],
				Titular:        line[1],
				Ruc:            line[2],
				TituloProyecto: line[3],
				UnidadProyecto: line[4],
				Tipo:           line[5],
				Actividad:      line[6],
				FechaInicio:    line[7],
				Estado:         line[8],
				Descripcion:    line[9],
				Longitud:       line[10],
				Latitud:        line[11],
				Resolucion:     line[12],
				Label:          line[13],
			})
		}
	}
	return data, nil
}