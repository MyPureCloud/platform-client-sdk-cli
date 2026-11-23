package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesnlusupportedlanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesnlusupportedlanguageDud struct { 
    Language string `json:"language"`


    SupportedEntityTypes []string `json:"supportedEntityTypes"`


    

}

// Capabilitiesnlusupportedlanguage
type Capabilitiesnlusupportedlanguage struct { 
    


    


    // IntentClassification - Establishes if administration actions are available/allowed for the supported language
    IntentClassification bool `json:"intentClassification"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesnlusupportedlanguage) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesnlusupportedlanguage) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesnlusupportedlanguage

    if CapabilitiesnlusupportedlanguageMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesnlusupportedlanguageMarshalled = true

    return json.Marshal(&struct {
        
        IntentClassification bool `json:"intentClassification"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

