package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttargetdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttargetdefinitionDud struct { 
    


    

}

// Agenticvirtualagenttargetdefinition - Target definition for input or output data type properties
type Agenticvirtualagenttargetdefinition struct { 
    // VarType - The type of target.
    VarType string `json:"type"`


    // Target - The reference target object. Contains information on the Conversation Attributes schema.
    Target Agenticvirtualagenttargetreferencedefinition `json:"target"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttargetdefinition) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttargetdefinition) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttargetdefinition

    if AgenticvirtualagenttargetdefinitionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttargetdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Target Agenticvirtualagenttargetreferencedefinition `json:"target"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

