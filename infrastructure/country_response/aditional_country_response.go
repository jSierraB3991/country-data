package countryresponse

type Name struct {
	Common     string                `json:"common"`
	Official   string                `json:"official"`
	NativeName map[string]NativeName `json:"nativeName"`
}

type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type IDD struct {
	Root     string   `json:"root"`
	Suffixes []string `json:"suffixes"`
}

type Translation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Demonym struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}

type CoatOfArms struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
}

type CapitalInfo struct {
	Latlng []float64 `json:"latlng"`
}
