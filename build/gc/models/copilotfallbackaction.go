package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotfallbackactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotfallbackactionDud struct { 
    


    

}

// Copilotfallbackaction
type Copilotfallbackaction struct { 
    // ActionType - Type of action.
    ActionType string `json:"actionType"`


    // Attributes - Action specific attributes, if any. Maximum 100 of string key-value pair allowed.
    Attributes map[string]string `json:"attributes"`

}

// String returns a JSON representation of the model
func (o *Copilotfallbackaction) String() string {
    
     o.Attributes = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotfallbackaction) MarshalJSON() ([]byte, error) {
    type Alias Copilotfallbackaction

    if CopilotfallbackactionMarshalled {
        return []byte("{}"), nil
    }
    CopilotfallbackactionMarshalled = true

    return json.Marshal(&struct {
        
        ActionType string `json:"actionType"`
        
        Attributes map[string]string `json:"attributes"`
        *Alias
    }{

        


        
        Attributes: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

