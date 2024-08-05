package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbidworkplanoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbidworkplanoverriderequestDud struct { 
    


    


    

}

// Agentbidworkplanoverriderequest
type Agentbidworkplanoverriderequest struct { 
    // AgentId - The ID of agent
    AgentId string `json:"agentId"`


    // OverrideWorkPlanId - The ID of the work plan that overrides the assigned work plan for the agent
    OverrideWorkPlanId string `json:"overrideWorkPlanId"`


    // OverrideReason - The reason for overriding the assigned work plan. This must be null if overrideWorkPlanId is not specified
    OverrideReason string `json:"overrideReason"`

}

// String returns a JSON representation of the model
func (o *Agentbidworkplanoverriderequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbidworkplanoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Agentbidworkplanoverriderequest

    if AgentbidworkplanoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    AgentbidworkplanoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentId string `json:"agentId"`
        
        OverrideWorkPlanId string `json:"overrideWorkPlanId"`
        
        OverrideReason string `json:"overrideReason"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

