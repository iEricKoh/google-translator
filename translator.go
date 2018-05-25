package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultAPIURL = "https://translate.google.cn/translate_a/single"
)

func Translate(text, sl, tl string) (string, error) {
	option := &Option{
		Text:       text,
		SourceLang: sl,
		TargetLang: tl,
	}

	url, err := GenerateURL(option)

	if err != nil {
		return "", err
	}

	return doTranslate(url)
}

func doTranslate(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "AndroidTranslate/5.3.0.RC02.130475354-53000263 5.1 phone TRANSLATE_OPM5_TEST_1")

	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	var transl Translation

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&transl)

	if err != nil {
		return "", err
	}

	result := ""
	for _, v := range transl.Sentences {
		trans := v.Trans
		if trans != "" {
			result += " " + trans
		}
	}

	return result, nil
}

func GenerateURL(opt *Option) (string, error) {
	if err := opt.Validate(); err != nil {
		return "", err
	}

	params := &url.Values{}
	params.Set("client", opt.GetClient())
	params.Set("dj", opt.GetDJ())

	for _, dt := range opt.GetDT() {
		params.Add("dt", dt)
	}

	params.Set("hl", opt.GetHL())
	params.Set("ie", opt.GetEncoding())

	params.Set("oe", opt.GetEncoding())
	params.Set("otf", opt.GetOTF())
	params.Set("q", opt.GetText())

	params.Set("sl", opt.GetSourceLang())
	params.Set("tl", opt.GetTargetLang())

	return fmt.Sprintf("%s?%s", defaultAPIURL, params.Encode()), nil
}
