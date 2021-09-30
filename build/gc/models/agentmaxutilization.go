package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmaxutilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmaxutilizationDud struct { 
    


    

}

// Agentmaxutilization
type Agentmaxutilization struct { 
    // Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
    Utilization map[string]Mediautilization `json:"utilization"`


    // Level
    Level string `json:"level"`

}

// String returns a JSON representation of the model
func (o *Agentmaxutilization) String() string {
    
    
     o.Utilization = map[string]Mediautilization{"": {}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmaxutilization) MarshalJSON() ([]byte, error) {
    type Alias Agentmaxutilization

    if AgentmaxutilizationMarshalled {
        return []byte("{}"), nil
    }
    AgentmaxutilizationMarshalled = true

    return json.Marshal(&struct { 
        Utilization map[string]Mediautilization `json:"utilization"`
        
        Level string `json:"level"`
        
        *Alias
    }{
        

        
        Utilization: map[string]Mediautilization{"": {}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

