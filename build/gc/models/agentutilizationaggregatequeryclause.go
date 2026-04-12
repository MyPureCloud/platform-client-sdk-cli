package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentutilizationaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentutilizationaggregatequeryclauseDud struct { 
    


    

}

// Agentutilizationaggregatequeryclause
type Agentutilizationaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agentutilizationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentutilizationaggregatequeryclause) String() string {
    
     o.Predicates = []Agentutilizationaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentutilizationaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Agentutilizationaggregatequeryclause

    if AgentutilizationaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgentutilizationaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agentutilizationaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agentutilizationaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

