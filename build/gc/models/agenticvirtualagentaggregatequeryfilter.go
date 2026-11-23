package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentaggregatequeryfilterDud struct { 
    


    


    

}

// Agenticvirtualagentaggregatequeryfilter
type Agenticvirtualagentaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agenticvirtualagentaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agenticvirtualagentaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentaggregatequeryfilter) String() string {
    
     o.Clauses = []Agenticvirtualagentaggregatequeryclause{{}} 
     o.Predicates = []Agenticvirtualagentaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentaggregatequeryfilter

    if AgenticvirtualagentaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agenticvirtualagentaggregatequeryclause `json:"clauses"`
        
        Predicates []Agenticvirtualagentaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agenticvirtualagentaggregatequeryclause{{}},
        


        
        Predicates: []Agenticvirtualagentaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

