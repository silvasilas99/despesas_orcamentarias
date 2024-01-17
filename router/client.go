package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Path/to/model"
)

// Fetch is exported ...
func FetchCrypto(fiat string, crypto string) (string, error) {
	URL := "http://dados.recife.pe.gov.br/api/3/action/datastore_search?resource_id=d4d8a7f0-d4be-4397-b950-f0c991438111&limit=5 "
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var cResp model.Cryptoresponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
