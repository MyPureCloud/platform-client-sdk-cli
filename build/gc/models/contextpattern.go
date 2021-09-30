package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContextpatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContextpatternDud struct { 
    

}

// Contextpattern
type Contextpattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Entitytypecriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Contextpattern) String() string {
    
    
     o.Criteria = []Entitytypecriteria{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contextpattern) MarshalJSON() ([]byte, error) {
    type Alias Contextpattern

    if ContextpatternMarshalled {
        return []byte("{}"), nil
    }
    ContextpatternMarshalled = true

    return json.Marshal(&struct { 
        Criteria []Entitytypecriteria `json:"criteria"`
        
        *Alias
    }{
        

        
        Criteria: []Entitytypecriteria{{}},
        

        
        Alias: (*Alias)(u),
    })
}

