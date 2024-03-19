package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggrouplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggrouplistingDud struct { 
    

}

// Staffinggrouplisting
type Staffinggrouplisting struct { 
    // Entities
    Entities []Staffinggroupresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Staffinggrouplisting) String() string {
     o.Entities = []Staffinggroupresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggrouplisting) MarshalJSON() ([]byte, error) {
    type Alias Staffinggrouplisting

    if StaffinggrouplistingMarshalled {
        return []byte("{}"), nil
    }
    StaffinggrouplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Staffinggroupresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Staffinggroupresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

