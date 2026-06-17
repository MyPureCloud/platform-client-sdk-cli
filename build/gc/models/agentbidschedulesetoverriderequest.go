package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbidschedulesetoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbidschedulesetoverriderequestDud struct { 
    


    


    

}

// Agentbidschedulesetoverriderequest
type Agentbidschedulesetoverriderequest struct { 
    // AgentId - The ID of the agent
    AgentId string `json:"agentId"`


    // OverrideScheduleSetId - If provided, the schedule set overrides the agent's assigned schedule set
    OverrideScheduleSetId string `json:"overrideScheduleSetId"`


    // OverrideReason - The reason the assigned schedule set has been overridden. This must be null if no override schedule is set
    OverrideReason string `json:"overrideReason"`

}

// String returns a JSON representation of the model
func (o *Agentbidschedulesetoverriderequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbidschedulesetoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Agentbidschedulesetoverriderequest

    if AgentbidschedulesetoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    AgentbidschedulesetoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentId string `json:"agentId"`
        
        OverrideScheduleSetId string `json:"overrideScheduleSetId"`
        
        OverrideReason string `json:"overrideReason"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

