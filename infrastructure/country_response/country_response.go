package countryresponse

type Country struct {
	Name         Name                   `json:"name"`
	TLD          []string               `json:"tld"`
	CCA2         string                 `json:"cca2"`
	CCN3         string                 `json:"ccn3"`
	CCA3         string                 `json:"cca3"`
	CIOC         string                 `json:"cioc"`
	Independent  bool                   `json:"independent"`
	Status       string                 `json:"status"`
	UnMember     bool                   `json:"unMember"`
	Currencies   map[string]Currency    `json:"currencies"`
	IDD          IDD                    `json:"idd"`
	Capital      []string               `json:"capital"`
	AltSpellings []string               `json:"altSpellings"`
	Region       string                 `json:"region"`
	Subregion    string                 `json:"subregion"`
	Languages    map[string]string      `json:"languages"`
	Translations map[string]Translation `json:"translations"`
	Latlng       []float64              `json:"latlng"`
	Landlocked   bool                   `json:"landlocked"`
	Borders      []string               `json:"borders"`
	Area         float64                `json:"area"`
	Demonyms     map[string]Demonym     `json:"demonyms"`
	Flag         string                 `json:"flag"`
	Maps         Maps                   `json:"maps"`
	Population   int                    `json:"population"`
	Gini         map[string]float64     `json:"gini"`
	FIFA         string                 `json:"fifa"`
	Car          Car                    `json:"car"`
	Timezones    []string               `json:"timezones"`
	Continents   []string               `json:"continents"`
	Flags        Flags                  `json:"flags"`
	CoatOfArms   CoatOfArms             `json:"coatOfArms"`
	StartOfWeek  string                 `json:"startOfWeek"`
	CapitalInfo  CapitalInfo            `json:"capitalInfo"`
}
