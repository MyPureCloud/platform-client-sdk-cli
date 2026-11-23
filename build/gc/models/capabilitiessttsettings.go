package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiessttsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiessttsettingsDud struct { 
    


    

}

// Capabilitiessttsettings
type Capabilitiessttsettings struct { 
    // DefaultEngine
    DefaultEngine string `json:"defaultEngine"`


    // LanguageOverrides
    LanguageOverrides map[string]Capabilitiessttlanguageoverride `json:"languageOverrides"`

}

// String returns a JSON representation of the model
func (o *Capabilitiessttsettings) String() string {
    
     o.LanguageOverrides = map[string]Capabilitiessttlanguageoverride{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiessttsettings) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiessttsettings

    if CapabilitiessttsettingsMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiessttsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DefaultEngine string `json:"defaultEngine"`
        
        LanguageOverrides map[string]Capabilitiessttlanguageoverride `json:"languageOverrides"`
        *Alias
    }{

        


        
        LanguageOverrides: map[string]Capabilitiessttlanguageoverride{"": {}},
        

        Alias: (*Alias)(u),
    })
}

