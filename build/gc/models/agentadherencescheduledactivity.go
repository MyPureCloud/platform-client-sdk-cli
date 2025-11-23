package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentadherencescheduledactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentadherencescheduledactivityDud struct { 
    


    


    

}

// Agentadherencescheduledactivity
type Agentadherencescheduledactivity struct { 
    // ActivityCodeId - The ID of the activity code from the schedule
    ActivityCodeId string `json:"activityCodeId"`


    // StartOffsetSeconds - Start offset in seconds relative to query start time
    StartOffsetSeconds int `json:"startOffsetSeconds"`


    // EndOffsetSeconds - End offset in seconds relative to query start time
    EndOffsetSeconds int `json:"endOffsetSeconds"`

}

// String returns a JSON representation of the model
func (o *Agentadherencescheduledactivity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentadherencescheduledactivity) MarshalJSON() ([]byte, error) {
    type Alias Agentadherencescheduledactivity

    if AgentadherencescheduledactivityMarshalled {
        return []byte("{}"), nil
    }
    AgentadherencescheduledactivityMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        StartOffsetSeconds int `json:"startOffsetSeconds"`
        
        EndOffsetSeconds int `json:"endOffsetSeconds"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

