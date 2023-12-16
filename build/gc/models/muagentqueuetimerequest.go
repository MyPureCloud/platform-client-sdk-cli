package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MuagentqueuetimerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MuagentqueuetimerequestDud struct { 
    


    

}

// Muagentqueuetimerequest
type Muagentqueuetimerequest struct { 
    // ManagementUnitId - ID of the management unit
    ManagementUnitId string `json:"managementUnitId"`


    // AgentOnQueueTimes - List of Agent queue times within the management unit
    AgentOnQueueTimes []Agentqueuetimerequest `json:"agentOnQueueTimes"`

}

// String returns a JSON representation of the model
func (o *Muagentqueuetimerequest) String() string {
    
     o.AgentOnQueueTimes = []Agentqueuetimerequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Muagentqueuetimerequest) MarshalJSON() ([]byte, error) {
    type Alias Muagentqueuetimerequest

    if MuagentqueuetimerequestMarshalled {
        return []byte("{}"), nil
    }
    MuagentqueuetimerequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitId string `json:"managementUnitId"`
        
        AgentOnQueueTimes []Agentqueuetimerequest `json:"agentOnQueueTimes"`
        *Alias
    }{

        


        
        AgentOnQueueTimes: []Agentqueuetimerequest{{}},
        

        Alias: (*Alias)(u),
    })
}

