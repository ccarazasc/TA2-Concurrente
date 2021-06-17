package models

type RowData struct {
	CodigoHogar  string `json:"homeCode"`
	Departamento string `json:"department"`
	Provincia    string    `json:"province"`
	Distrito     string    `json:"district"`
	Area         string    `json:"area"`
	NuPersona    string    `json:"nuPerson"`
	Genero       int    `json:"genre"`
	Edad         int    `json:"age"`
}
