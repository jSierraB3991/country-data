package countryclient

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"

	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func GetCountries(countryUrl string) (*countryresponse.ApiResponse, error) {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: false},
		ForceAttemptHTTP2: false, // <--- esta línea desactiva HTTP/2
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", countryUrl, nil)
	if err != nil {
		return nil, err
	}

	// Establece un User-Agent como si fueras un navegador
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyGoClient/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse countryresponse.ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return nil, err
	}
	return &apiResponse, nil
}
