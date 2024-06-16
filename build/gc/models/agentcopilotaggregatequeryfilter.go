package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentcopilotaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentcopilotaggregatequeryfilterDud struct { 
    


    


    

}

// Agentcopilotaggregatequeryfilter
type Agentcopilotaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agentcopilotaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agentcopilotaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentcopilotaggregatequeryfilter) String() string {
    
     o.Clauses = []Agentcopilotaggregatequeryclause{{}} 
     o.Predicates = []Agentcopilotaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentcopilotaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Agentcopilotaggregatequeryfilter

    if AgentcopilotaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    AgentcopilotaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agentcopilotaggregatequeryclause `json:"clauses"`
        
        Predicates []Agentcopilotaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agentcopilotaggregatequeryclause{{}},
        


        
        Predicates: []Agentcopilotaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

