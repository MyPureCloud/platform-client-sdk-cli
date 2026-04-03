package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventidentifiersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventidentifiersDud struct { 
    

}

// Externaleventidentifiers
type Externaleventidentifiers struct { 
    // ExternalIds - The list of identifiers of the contact
    ExternalIds []Externaleventexternalid `json:"externalIds"`

}

// String returns a JSON representation of the model
func (o *Externaleventidentifiers) String() string {
     o.ExternalIds = []Externaleventexternalid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventidentifiers) MarshalJSON() ([]byte, error) {
    type Alias Externaleventidentifiers

    if ExternaleventidentifiersMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventidentifiersMarshalled = true

    return json.Marshal(&struct {
        
        ExternalIds []Externaleventexternalid `json:"externalIds"`
        *Alias
    }{

        
        ExternalIds: []Externaleventexternalid{{}},
        

        Alias: (*Alias)(u),
    })
}

