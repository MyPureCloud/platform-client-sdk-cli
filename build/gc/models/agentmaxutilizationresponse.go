package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmaxutilizationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmaxutilizationresponseDud struct { 
    


    

}

// Agentmaxutilizationresponse
type Agentmaxutilizationresponse struct { 
    // Utilization - Map of media type to utilization settings.
    Utilization map[string]Mediautilization `json:"utilization"`


    // Level
    Level string `json:"level"`

}

// String returns a JSON representation of the model
func (o *Agentmaxutilizationresponse) String() string {
     o.Utilization = map[string]Mediautilization{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmaxutilizationresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentmaxutilizationresponse

    if AgentmaxutilizationresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentmaxutilizationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Utilization map[string]Mediautilization `json:"utilization"`
        
        Level string `json:"level"`
        *Alias
    }{

        
        Utilization: map[string]Mediautilization{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

