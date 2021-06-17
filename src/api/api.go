package api

import (
	"fmt"
	"net/http"
)

func GetData(res http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(res,"Hello SECOND TEST API")
}
