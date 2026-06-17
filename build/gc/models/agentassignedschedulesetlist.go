package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentassignedschedulesetlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentassignedschedulesetlistDud struct { 
    

}

// Agentassignedschedulesetlist
type Agentassignedschedulesetlist struct { 
    // AgentAssignedShiftSets - The shift sets, along with the assigned agents
    AgentAssignedShiftSets []Agentassignedshiftset `json:"agentAssignedShiftSets"`

}

// String returns a JSON representation of the model
func (o *Agentassignedschedulesetlist) String() string {
     o.AgentAssignedShiftSets = []Agentassignedshiftset{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentassignedschedulesetlist) MarshalJSON() ([]byte, error) {
    type Alias Agentassignedschedulesetlist

    if AgentassignedschedulesetlistMarshalled {
        return []byte("{}"), nil
    }
    AgentassignedschedulesetlistMarshalled = true

    return json.Marshal(&struct {
        
        AgentAssignedShiftSets []Agentassignedshiftset `json:"agentAssignedShiftSets"`
        *Alias
    }{

        
        AgentAssignedShiftSets: []Agentassignedshiftset{{}},
        

        Alias: (*Alias)(u),
    })
}

