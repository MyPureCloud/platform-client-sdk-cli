package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionaggregatequeryfilterDud struct { 
    


    


    

}

// Flowexecutionaggregatequeryfilter
type Flowexecutionaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Flowexecutionaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowexecutionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionaggregatequeryfilter) String() string {
    
     o.Clauses = []Flowexecutionaggregatequeryclause{{}} 
     o.Predicates = []Flowexecutionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionaggregatequeryfilter

    if FlowexecutionaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Flowexecutionaggregatequeryclause `json:"clauses"`
        
        Predicates []Flowexecutionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Flowexecutionaggregatequeryclause{{}},
        


        
        Predicates: []Flowexecutionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

