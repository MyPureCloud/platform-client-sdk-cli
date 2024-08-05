package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanbiddingpreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanbiddingpreferenceDud struct { 
    


    

}

// Agentworkplanbiddingpreference
type Agentworkplanbiddingpreference struct { 
    // WorkPlan - The work plan that belongs to the agent's bid group
    WorkPlan Workplanreference `json:"workPlan"`


    // Priority - The agent's priority for this work plan, with 1 being the best priority. Null if priority is not set for the work plan
    Priority int `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Agentworkplanbiddingpreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanbiddingpreference) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanbiddingpreference

    if AgentworkplanbiddingpreferenceMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanbiddingpreferenceMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlan Workplanreference `json:"workPlan"`
        
        Priority int `json:"priority"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

