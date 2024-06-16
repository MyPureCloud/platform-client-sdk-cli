package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentcopilotaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentcopilotaggregatequeryclauseDud struct { 
    


    

}

// Agentcopilotaggregatequeryclause
type Agentcopilotaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agentcopilotaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentcopilotaggregatequeryclause) String() string {
    
     o.Predicates = []Agentcopilotaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentcopilotaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Agentcopilotaggregatequeryclause

    if AgentcopilotaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgentcopilotaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agentcopilotaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agentcopilotaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

