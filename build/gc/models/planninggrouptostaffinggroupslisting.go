package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggrouptostaffinggroupslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggrouptostaffinggroupslistingDud struct { 
    

}

// Planninggrouptostaffinggroupslisting
type Planninggrouptostaffinggroupslisting struct { 
    // Entities
    Entities []Planninggrouptostaffinggroupsresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Planninggrouptostaffinggroupslisting) String() string {
     o.Entities = []Planninggrouptostaffinggroupsresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggrouptostaffinggroupslisting) MarshalJSON() ([]byte, error) {
    type Alias Planninggrouptostaffinggroupslisting

    if PlanninggrouptostaffinggroupslistingMarshalled {
        return []byte("{}"), nil
    }
    PlanninggrouptostaffinggroupslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Planninggrouptostaffinggroupsresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Planninggrouptostaffinggroupsresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

