package ExpenseSearchController

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	//"github.com/silvasilas99/despesas_orcamentarias/models"
)

func get(w http.ResponseWriter, r *http.Request) {
	URL := "http://dados.recife.pe.gov.br/api/3/action/datastore_search?resource_id=d4d8a7f0-d4be-4397-b950-f0c991438111&limit=5"
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Não foi possivel obter os resultados da requisição.")
	}

	responseBody, error := io.ReadAll(resp.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := json.NewEncoder(w).Encode(responseBody)
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Response body: ", formattedData)

	defer resp.Body.Close()
}
