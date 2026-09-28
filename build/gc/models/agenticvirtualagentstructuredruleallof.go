package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentstructuredruleallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentstructuredruleallofDud struct { 
    


    


    

}

// Agenticvirtualagentstructuredruleallof - A single structured input validation rule.
type Agenticvirtualagentstructuredruleallof struct { 
    // Name - Target name of the tool input this rule applies to.
    Name string `json:"name"`


    // Operator - Operator to apply to the input value.
    Operator string `json:"operator"`


    // Value - Value to compare against. May be a string, integer, number, boolean, or null. Not required for 'IsNull' or 'IsNotNull' operators.
    Value interface{} `json:"value"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentstructuredruleallof) String() string {
    
    
     o.Value = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentstructuredruleallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentstructuredruleallof

    if AgenticvirtualagentstructuredruleallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentstructuredruleallofMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Value interface{} `json:"value"`
        *Alias
    }{

        


        


        
        Value: Interface{},
        

        Alias: (*Alias)(u),
    })
}

