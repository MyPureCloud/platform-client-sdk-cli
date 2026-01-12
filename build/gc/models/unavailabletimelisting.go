package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnavailabletimelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnavailabletimelistingDud struct { 
    

}

// Unavailabletimelisting
type Unavailabletimelisting struct { 
    // Entities
    Entities []Unavailabletime `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Unavailabletimelisting) String() string {
     o.Entities = []Unavailabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unavailabletimelisting) MarshalJSON() ([]byte, error) {
    type Alias Unavailabletimelisting

    if UnavailabletimelistingMarshalled {
        return []byte("{}"), nil
    }
    UnavailabletimelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Unavailabletime `json:"entities"`
        *Alias
    }{

        
        Entities: []Unavailabletime{{}},
        

        Alias: (*Alias)(u),
    })
}

