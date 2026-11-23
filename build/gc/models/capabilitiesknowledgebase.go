package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesknowledgebaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesknowledgebaseDud struct { 
    

}

// Capabilitiesknowledgebase
type Capabilitiesknowledgebase struct { 
    // Language - Language to use for knowledgebase handling
    Language string `json:"language"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesknowledgebase) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesknowledgebase) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesknowledgebase

    if CapabilitiesknowledgebaseMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesknowledgebaseMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

