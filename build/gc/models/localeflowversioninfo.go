package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocaleflowversioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocaleflowversioninfoDud struct { 
    

}

// Localeflowversioninfo
type Localeflowversioninfo struct { 
    // NluVersion - NLU Version Info for this flow version.
    NluVersion Addressableentityref `json:"nluVersion"`

}

// String returns a JSON representation of the model
func (o *Localeflowversioninfo) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localeflowversioninfo) MarshalJSON() ([]byte, error) {
    type Alias Localeflowversioninfo

    if LocaleflowversioninfoMarshalled {
        return []byte("{}"), nil
    }
    LocaleflowversioninfoMarshalled = true

    return json.Marshal(&struct {
        
        NluVersion Addressableentityref `json:"nluVersion"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

