# google-translator

FREE Google Translate client for Go.

### Installation

Download and install it:
```
go get github.com/iEricKoh/google-translator
```

### Usage
```
import google "google-translator"

google.Translate("hello", "auto", "ja")
```

### API

#### Translate
```
func Translate(text, sl, tl string) (string, error)
```
Params:
* `text` - The text you wanna translate
* `sl` - Source language
* `tl` - Target language


Returns:
* `result` - The text after translated
* `error` - An error during translation
