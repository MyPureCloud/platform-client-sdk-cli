package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentguardrailssettingsallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentguardrailssettingsallofDud struct { 
    


    


    

}

// Agenticvirtualagentguardrailssettingsallof - Guardrails event response settings.
type Agenticvirtualagentguardrailssettingsallof struct { 
    // Message - Message the virtual agent should return when this event is handled.
    Message string `json:"message"`


    // ViolationThreshold - Number of guardrail violations allowed before the threshold is crossed.
    ViolationThreshold int `json:"violationThreshold"`


    // ViolationThresholdCrossedMessage - Message the virtual agent should return when the guardrail violation threshold is crossed.
    ViolationThresholdCrossedMessage string `json:"violationThresholdCrossedMessage"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentguardrailssettingsallof) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentguardrailssettingsallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentguardrailssettingsallof

    if AgenticvirtualagentguardrailssettingsallofMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentguardrailssettingsallofMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        ViolationThreshold int `json:"violationThreshold"`
        
        ViolationThresholdCrossedMessage string `json:"violationThresholdCrossedMessage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

