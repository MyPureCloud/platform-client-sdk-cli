package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitactivitycodelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitactivitycodelistingDud struct { 
    

}

// Businessunitactivitycodelisting
type Businessunitactivitycodelisting struct { 
    // Entities
    Entities []Businessunitactivitycode `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Businessunitactivitycodelisting) String() string {
     o.Entities = []Businessunitactivitycode{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitactivitycodelisting) MarshalJSON() ([]byte, error) {
    type Alias Businessunitactivitycodelisting

    if BusinessunitactivitycodelistingMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitactivitycodelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Businessunitactivitycode `json:"entities"`
        *Alias
    }{

        
        Entities: []Businessunitactivitycode{{}},
        

        Alias: (*Alias)(u),
    })
}

