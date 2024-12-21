package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateuserfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateuserfilterDud struct { 
    


    


    

}

// Agentstateuserfilter
type Agentstateuserfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agentstateagentqueryclause `json:"clauses"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstateagentquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstateuserfilter) String() string {
    
     o.Clauses = []Agentstateagentqueryclause{{}} 
     o.Predicates = []Agentstateagentquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateuserfilter) MarshalJSON() ([]byte, error) {
    type Alias Agentstateuserfilter

    if AgentstateuserfilterMarshalled {
        return []byte("{}"), nil
    }
    AgentstateuserfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agentstateagentqueryclause `json:"clauses"`
        
        Predicates []Agentstateagentquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agentstateagentqueryclause{{}},
        


        
        Predicates: []Agentstateagentquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

