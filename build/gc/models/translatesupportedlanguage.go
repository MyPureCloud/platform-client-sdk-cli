package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranslatesupportedlanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranslatesupportedlanguageDud struct { 
    


    

}

// Translatesupportedlanguage
type Translatesupportedlanguage struct { 
    // LanguageName - Supported translation language name, natively spelled (ex. Español)
    LanguageName string `json:"languageName"`


    // LanguageCode - Supported translation language code. See - https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html#what-is-languages-supported
    LanguageCode string `json:"languageCode"`

}

// String returns a JSON representation of the model
func (o *Translatesupportedlanguage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Translatesupportedlanguage) MarshalJSON() ([]byte, error) {
    type Alias Translatesupportedlanguage

    if TranslatesupportedlanguageMarshalled {
        return []byte("{}"), nil
    }
    TranslatesupportedlanguageMarshalled = true

    return json.Marshal(&struct {
        
        LanguageName string `json:"languageName"`
        
        LanguageCode string `json:"languageCode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

