package google

import "errors"

type Option struct {
	Client string

	Encoding string

	SourceLang string
	TargetLang string

	Text string

	HL  string
	OTF string
	DT  []string

	DJ string
}

func (o *Option) Validate() error {
	if o.Text == "" {
		return errors.New("The text must not be empty")
	} else if len(o.Text) > 500 {
		return errors.New("The text length must be less than the 500")
	} else if o.TargetLang == "" {
		return errors.New("The targe language is required")
	}
	return nil
}

func (o *Option) GetClient() string {
	return "at"
}

func (g *Option) GetDT() []string {
	return []string{"t", "ld", "qca", "rm", "bd"}
}

func (g *Option) GetDJ() string {
	return "1"
}

func (o *Option) GetText() string {
	return o.Text
}

func (o *Option) GetSourceLang() string {
	if o.SourceLang != "" {
		return o.SourceLang
	}
	return "auto"
}

func (o *Option) GetTargetLang() string {
	if o.TargetLang != "" {
		return o.TargetLang
	}
	return "en"
}

func (o *Option) GetHL() string {
	if o.HL != "" {
		return o.HL
	}
	return o.GetTargetLang()
}

func (o *Option) GetEncoding() string {
	if o.Encoding != "" {
		return o.Encoding
	}
	return "UTF-8"
}

func (o Option) GetOTF() string {
	return "2"
}
