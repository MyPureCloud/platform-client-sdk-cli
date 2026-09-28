package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentguardrailssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentguardrailssettingsDud struct { 
    


    


    

}

// Agenticvirtualagentguardrailssettings
type Agenticvirtualagentguardrailssettings struct { 
    // Message - Message the virtual agent should return when this event is handled.
    Message string `json:"message"`


    // ViolationThreshold - Number of guardrail violations allowed before the threshold is crossed.
    ViolationThreshold int `json:"violationThreshold"`


    // ViolationThresholdCrossedMessage - Message the virtual agent should return when the guardrail violation threshold is crossed.
    ViolationThresholdCrossedMessage string `json:"violationThresholdCrossedMessage"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentguardrailssettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentguardrailssettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentguardrailssettings

    if AgenticvirtualagentguardrailssettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentguardrailssettingsMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        ViolationThreshold int `json:"violationThreshold"`
        
        ViolationThresholdCrossedMessage string `json:"violationThresholdCrossedMessage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

