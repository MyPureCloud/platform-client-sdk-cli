package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesorganizationDud struct { 
    

}

// Capabilitiesorganization
type Capabilitiesorganization struct { 
    // Products
    Products []string `json:"products"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesorganization) String() string {
     o.Products = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesorganization) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesorganization

    if CapabilitiesorganizationMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Products []string `json:"products"`
        *Alias
    }{

        
        Products: []string{""},
        

        Alias: (*Alias)(u),
    })
}

