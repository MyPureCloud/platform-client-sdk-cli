package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotconditionDud struct { 
    


    

}

// Copilotcondition
type Copilotcondition struct { 
    // ConditionType - Type of condition.
    ConditionType string `json:"conditionType"`


    // ConditionValues - Condition values.
    ConditionValues []string `json:"conditionValues"`

}

// String returns a JSON representation of the model
func (o *Copilotcondition) String() string {
    
     o.ConditionValues = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcondition) MarshalJSON() ([]byte, error) {
    type Alias Copilotcondition

    if CopilotconditionMarshalled {
        return []byte("{}"), nil
    }
    CopilotconditionMarshalled = true

    return json.Marshal(&struct {
        
        ConditionType string `json:"conditionType"`
        
        ConditionValues []string `json:"conditionValues"`
        *Alias
    }{

        


        
        ConditionValues: []string{""},
        

        Alias: (*Alias)(u),
    })
}

