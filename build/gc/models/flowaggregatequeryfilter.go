package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowaggregatequeryfilterDud struct { 
    


    


    

}

// Flowaggregatequeryfilter
type Flowaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Flowaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryfilter) String() string {
    
     o.Clauses = []Flowaggregatequeryclause{{}} 
     o.Predicates = []Flowaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Flowaggregatequeryfilter

    if FlowaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    FlowaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Flowaggregatequeryclause `json:"clauses"`
        
        Predicates []Flowaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Flowaggregatequeryclause{{}},
        


        
        Predicates: []Flowaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

