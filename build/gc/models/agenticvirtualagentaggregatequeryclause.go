package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentaggregatequeryclauseDud struct { 
    


    

}

// Agenticvirtualagentaggregatequeryclause
type Agenticvirtualagentaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Agenticvirtualagentaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentaggregatequeryclause) String() string {
    
     o.Predicates = []Agenticvirtualagentaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentaggregatequeryclause

    if AgenticvirtualagentaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agenticvirtualagentaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agenticvirtualagentaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

