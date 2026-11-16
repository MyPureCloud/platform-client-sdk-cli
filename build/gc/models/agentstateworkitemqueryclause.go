package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateworkitemqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateworkitemqueryclauseDud struct { 
    


    

}

// Agentstateworkitemqueryclause
type Agentstateworkitemqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstateworkitemquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstateworkitemqueryclause) String() string {
    
     o.Predicates = []Agentstateworkitemquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateworkitemqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Agentstateworkitemqueryclause

    if AgentstateworkitemqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgentstateworkitemqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agentstateworkitemquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agentstateworkitemquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

