package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentsbidassignedschedulesetoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentsbidassignedschedulesetoverriderequestDud struct { 
    

}

// Agentsbidassignedschedulesetoverriderequest
type Agentsbidassignedschedulesetoverriderequest struct { 
    // AgentScheduleSetOverrides - The agent schedule set overrides
    AgentScheduleSetOverrides []Agentbidschedulesetoverriderequest `json:"agentScheduleSetOverrides"`

}

// String returns a JSON representation of the model
func (o *Agentsbidassignedschedulesetoverriderequest) String() string {
     o.AgentScheduleSetOverrides = []Agentbidschedulesetoverriderequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentsbidassignedschedulesetoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Agentsbidassignedschedulesetoverriderequest

    if AgentsbidassignedschedulesetoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    AgentsbidassignedschedulesetoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentScheduleSetOverrides []Agentbidschedulesetoverriderequest `json:"agentScheduleSetOverrides"`
        *Alias
    }{

        
        AgentScheduleSetOverrides: []Agentbidschedulesetoverriderequest{{}},
        

        Alias: (*Alias)(u),
    })
}

