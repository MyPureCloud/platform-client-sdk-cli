package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanbiddingpreferencerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanbiddingpreferencerequestDud struct { 
    


    

}

// Agentworkplanbiddingpreferencerequest
type Agentworkplanbiddingpreferencerequest struct { 
    // WorkPlanId - The ID of the work plan that belongs to agent's bid group
    WorkPlanId string `json:"workPlanId"`


    // Priority - The agent's priority for this work plan, with 1 being the best priority. Null if priority is not set for the work plan
    Priority int `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Agentworkplanbiddingpreferencerequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanbiddingpreferencerequest) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanbiddingpreferencerequest

    if AgentworkplanbiddingpreferencerequestMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanbiddingpreferencerequestMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanId string `json:"workPlanId"`
        
        Priority int `json:"priority"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

