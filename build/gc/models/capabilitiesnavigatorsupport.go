package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesnavigatorsupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesnavigatorsupportDud struct { 
    

}

// Capabilitiesnavigatorsupport
type Capabilitiesnavigatorsupport struct { 
    // SupportedLanguages
    SupportedLanguages map[string]interface{} `json:"supportedLanguages"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesnavigatorsupport) String() string {
     o.SupportedLanguages = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesnavigatorsupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesnavigatorsupport

    if CapabilitiesnavigatorsupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesnavigatorsupportMarshalled = true

    return json.Marshal(&struct {
        
        SupportedLanguages map[string]interface{} `json:"supportedLanguages"`
        *Alias
    }{

        
        SupportedLanguages: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

