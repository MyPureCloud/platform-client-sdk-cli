package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulebiddingpreferencepriorityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulebiddingpreferencepriorityDud struct { 
    


    

}

// Agentschedulebiddingpreferencepriority
type Agentschedulebiddingpreferencepriority struct { 
    // ScheduleSetId - The ID of the schedule set that belongs to agent's bid group
    ScheduleSetId string `json:"scheduleSetId"`


    // Priority - The agent's priority for this schedule set. Lower numbers indicate higher priority, with 1 being the highest priority. Minimum value is 1. Null if priority is not set for the schedule set
    Priority int `json:"priority"`

}

// String returns a JSON representation of the model
func (o *Agentschedulebiddingpreferencepriority) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulebiddingpreferencepriority) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulebiddingpreferencepriority

    if AgentschedulebiddingpreferencepriorityMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulebiddingpreferencepriorityMarshalled = true

    return json.Marshal(&struct {
        
        ScheduleSetId string `json:"scheduleSetId"`
        
        Priority int `json:"priority"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

