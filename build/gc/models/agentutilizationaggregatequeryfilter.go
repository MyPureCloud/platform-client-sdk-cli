package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentutilizationaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentutilizationaggregatequeryfilterDud struct { 
    


    


    

}

// Agentutilizationaggregatequeryfilter
type Agentutilizationaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agentutilizationaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agentutilizationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentutilizationaggregatequeryfilter) String() string {
    
     o.Clauses = []Agentutilizationaggregatequeryclause{{}} 
     o.Predicates = []Agentutilizationaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentutilizationaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Agentutilizationaggregatequeryfilter

    if AgentutilizationaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    AgentutilizationaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agentutilizationaggregatequeryclause `json:"clauses"`
        
        Predicates []Agentutilizationaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agentutilizationaggregatequeryclause{{}},
        


        
        Predicates: []Agentutilizationaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

