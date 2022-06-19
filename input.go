package solc

type Input struct {
	Language string              `json:"language,omitempty"`
	Sources  map[string]SourceIn `json:"sources,omitempty"`
	Settings Settings            `json:"settings,omitempty"`
}

type SourceIn struct {
	Content   string   `json:"content,omitempty"`
	Keccak256 string   `json:"keccak256,omitempty"`
	Urls      []string `json:"urls,omitempty"`
}

type Settings struct {
	Optimizer       Optimizer                      `json:"optimizer,omitempty"`
	OutputSelection map[string]map[string][]string `json:"outputSelection,omitempty"`
}

type Optimizer struct {
	Enabled bool `json:"enabled,omitempty"`
	Runs    int  `json:"runs,omitempty"`
}
