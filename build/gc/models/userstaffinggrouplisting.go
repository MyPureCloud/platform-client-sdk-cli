package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstaffinggrouplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstaffinggrouplistingDud struct { 
    

}

// Userstaffinggrouplisting
type Userstaffinggrouplisting struct { 
    // Entities
    Entities []Userstaffinggroupresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Userstaffinggrouplisting) String() string {
     o.Entities = []Userstaffinggroupresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstaffinggrouplisting) MarshalJSON() ([]byte, error) {
    type Alias Userstaffinggrouplisting

    if UserstaffinggrouplistingMarshalled {
        return []byte("{}"), nil
    }
    UserstaffinggrouplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userstaffinggroupresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Userstaffinggroupresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

