package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontextpatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontextpatternDud struct { 
    

}

// Patchcontextpattern
type Patchcontextpattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Patchentitytypecriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Patchcontextpattern) String() string {
     o.Criteria = []Patchentitytypecriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontextpattern) MarshalJSON() ([]byte, error) {
    type Alias Patchcontextpattern

    if PatchcontextpatternMarshalled {
        return []byte("{}"), nil
    }
    PatchcontextpatternMarshalled = true

    return json.Marshal(&struct {
        
        Criteria []Patchentitytypecriteria `json:"criteria"`
        *Alias
    }{

        
        Criteria: []Patchentitytypecriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

