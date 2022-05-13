package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentownedroutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentownedroutingDud struct { 
    


    


    

}

// Agentownedrouting
type Agentownedrouting struct { 
    // EnableAgentOwnedCallbacks - Indicates if Agent Owned Callbacks are enabled for the queue
    EnableAgentOwnedCallbacks bool `json:"enableAgentOwnedCallbacks"`


    // MaxOwnedCallbackHours - The max amount of time a callback can be owned (in hours); Allowable range 1 - 168 hour(s) (inclusive)
    MaxOwnedCallbackHours int `json:"maxOwnedCallbackHours"`


    // MaxOwnedCallbackDelayHours - The max amount of time a callback can be scheduled out into the future (in hours); Allowable range 1 - 720 hour(s) (inclusive)
    MaxOwnedCallbackDelayHours int `json:"maxOwnedCallbackDelayHours"`

}

// String returns a JSON representation of the model
func (o *Agentownedrouting) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentownedrouting) MarshalJSON() ([]byte, error) {
    type Alias Agentownedrouting

    if AgentownedroutingMarshalled {
        return []byte("{}"), nil
    }
    AgentownedroutingMarshalled = true

    return json.Marshal(&struct {
        
        EnableAgentOwnedCallbacks bool `json:"enableAgentOwnedCallbacks"`
        
        MaxOwnedCallbackHours int `json:"maxOwnedCallbackHours"`
        
        MaxOwnedCallbackDelayHours int `json:"maxOwnedCallbackDelayHours"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

