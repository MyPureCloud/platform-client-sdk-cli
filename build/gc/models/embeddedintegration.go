package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmbeddedintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmbeddedintegrationDud struct { 
    


    

}

// Embeddedintegration
type Embeddedintegration struct { 
    // EnableWhitelist
    EnableWhitelist bool `json:"enableWhitelist"`


    // DomainWhitelist
    DomainWhitelist []string `json:"domainWhitelist"`

}

// String returns a JSON representation of the model
func (o *Embeddedintegration) String() string {
    
     o.DomainWhitelist = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Embeddedintegration) MarshalJSON() ([]byte, error) {
    type Alias Embeddedintegration

    if EmbeddedintegrationMarshalled {
        return []byte("{}"), nil
    }
    EmbeddedintegrationMarshalled = true

    return json.Marshal(&struct {
        
        EnableWhitelist bool `json:"enableWhitelist"`
        
        DomainWhitelist []string `json:"domainWhitelist"`
        *Alias
    }{

        


        
        DomainWhitelist: []string{""},
        

        Alias: (*Alias)(u),
    })
}

