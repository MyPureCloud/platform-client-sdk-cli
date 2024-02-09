package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestcontextpatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestcontextpatternDud struct { 
    

}

// Requestcontextpattern
type Requestcontextpattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Requestentitytypecriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Requestcontextpattern) String() string {
     o.Criteria = []Requestentitytypecriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestcontextpattern) MarshalJSON() ([]byte, error) {
    type Alias Requestcontextpattern

    if RequestcontextpatternMarshalled {
        return []byte("{}"), nil
    }
    RequestcontextpatternMarshalled = true

    return json.Marshal(&struct {
        
        Criteria []Requestentitytypecriteria `json:"criteria"`
        *Alias
    }{

        
        Criteria: []Requestentitytypecriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

