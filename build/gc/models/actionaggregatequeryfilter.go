package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionaggregatequeryfilterDud struct { 
    


    


    

}

// Actionaggregatequeryfilter
type Actionaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Actionaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Actionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Actionaggregatequeryfilter) String() string {
    
     o.Clauses = []Actionaggregatequeryclause{{}} 
     o.Predicates = []Actionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Actionaggregatequeryfilter

    if ActionaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    ActionaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Actionaggregatequeryclause `json:"clauses"`
        
        Predicates []Actionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Actionaggregatequeryclause{{}},
        


        
        Predicates: []Actionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

