package countryresponse

type MetaDataResponse struct {
	Total     int    `json:"total"`
	Count     int    `json:"count"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	More      bool   `json:"more"`
	RequestId string `json:"request_id"`
	Duration  int    `json:"duration"`
}

type DataApiResponse struct {
	Objects []CountryDataResponse `json:"objects"`
	Meta    MetaDataResponse      `json:"meta"`
}

type ApiResponse struct {
	Data DataApiResponse `json:"data"`
}
