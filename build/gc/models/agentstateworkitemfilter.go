package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateworkitemfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateworkitemfilterDud struct { 
    


    


    

}

// Agentstateworkitemfilter
type Agentstateworkitemfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agentstateworkitemqueryclause `json:"clauses"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstateworkitemquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstateworkitemfilter) String() string {
    
     o.Clauses = []Agentstateworkitemqueryclause{{}} 
     o.Predicates = []Agentstateworkitemquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateworkitemfilter) MarshalJSON() ([]byte, error) {
    type Alias Agentstateworkitemfilter

    if AgentstateworkitemfilterMarshalled {
        return []byte("{}"), nil
    }
    AgentstateworkitemfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agentstateworkitemqueryclause `json:"clauses"`
        
        Predicates []Agentstateworkitemquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agentstateworkitemqueryclause{{}},
        


        
        Predicates: []Agentstateworkitemquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

