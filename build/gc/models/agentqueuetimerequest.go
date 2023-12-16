package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueuetimerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueuetimerequestDud struct { 
    


    


    

}

// Agentqueuetimerequest
type Agentqueuetimerequest struct { 
    // AgentId - ID of the agent
    AgentId string `json:"agentId"`


    // StartOffsetMinutes - List of offsets in minutes from calculationStartDate
    StartOffsetMinutes []int `json:"startOffsetMinutes"`


    // OnQueueLengthMinutesPerInterval - List of on queue time lengths in minutes per interval of elements in startOffsetMinutes
    OnQueueLengthMinutesPerInterval []int `json:"onQueueLengthMinutesPerInterval"`

}

// String returns a JSON representation of the model
func (o *Agentqueuetimerequest) String() string {
    
     o.StartOffsetMinutes = []int{0} 
     o.OnQueueLengthMinutesPerInterval = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueuetimerequest) MarshalJSON() ([]byte, error) {
    type Alias Agentqueuetimerequest

    if AgentqueuetimerequestMarshalled {
        return []byte("{}"), nil
    }
    AgentqueuetimerequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentId string `json:"agentId"`
        
        StartOffsetMinutes []int `json:"startOffsetMinutes"`
        
        OnQueueLengthMinutesPerInterval []int `json:"onQueueLengthMinutesPerInterval"`
        *Alias
    }{

        


        
        StartOffsetMinutes: []int{0},
        


        
        OnQueueLengthMinutesPerInterval: []int{0},
        

        Alias: (*Alias)(u),
    })
}

