package countryclient

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func GetCountries(countryUrl string) ([]countryresponse.Country, error) {

	resp, err := http.Get(countryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var countries []countryresponse.Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return nil, err
	}
	return countries, nil
}
