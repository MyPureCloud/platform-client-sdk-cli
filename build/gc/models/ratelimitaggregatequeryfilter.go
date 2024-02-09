package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregatequeryfilterDud struct { 
    


    


    

}

// Ratelimitaggregatequeryfilter
type Ratelimitaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Ratelimitaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Ratelimitaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregatequeryfilter) String() string {
    
     o.Clauses = []Ratelimitaggregatequeryclause{{}} 
     o.Predicates = []Ratelimitaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregatequeryfilter

    if RatelimitaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Ratelimitaggregatequeryclause `json:"clauses"`
        
        Predicates []Ratelimitaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Ratelimitaggregatequeryclause{{}},
        


        
        Predicates: []Ratelimitaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

