package google

type Sentence struct {
	Trans       string `json:"trans,omitempty"`
	Orig        string `json:"orig,omitempty"`
	Backend     int    `json:"backend,omitempty"`
	SrcTranslit string `json:"src_translit,omitempty"`
}

type Translation struct {
	Sentences []*Sentence `json:"sentences"`
}
