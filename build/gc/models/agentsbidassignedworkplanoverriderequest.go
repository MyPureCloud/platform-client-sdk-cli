package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentsbidassignedworkplanoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentsbidassignedworkplanoverriderequestDud struct { 
    

}

// Agentsbidassignedworkplanoverriderequest
type Agentsbidassignedworkplanoverriderequest struct { 
    // AgentWorkPlanOverrides - The list of agent work plan overrides
    AgentWorkPlanOverrides []Agentbidworkplanoverriderequest `json:"agentWorkPlanOverrides"`

}

// String returns a JSON representation of the model
func (o *Agentsbidassignedworkplanoverriderequest) String() string {
     o.AgentWorkPlanOverrides = []Agentbidworkplanoverriderequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentsbidassignedworkplanoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Agentsbidassignedworkplanoverriderequest

    if AgentsbidassignedworkplanoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    AgentsbidassignedworkplanoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentWorkPlanOverrides []Agentbidworkplanoverriderequest `json:"agentWorkPlanOverrides"`
        *Alias
    }{

        
        AgentWorkPlanOverrides: []Agentbidworkplanoverriderequest{{}},
        

        Alias: (*Alias)(u),
    })
}

