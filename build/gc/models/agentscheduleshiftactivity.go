package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentscheduleshiftactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentscheduleshiftactivityDud struct { 
    


    


    

}

// Agentscheduleshiftactivity
type Agentscheduleshiftactivity struct { 
    // ActivityCategory - The activity category for which the agent is scheduled
    ActivityCategory string `json:"activityCategory"`


    // StartOffsetMinutes - The start offset of the activity, relative to referenceStartDate at the top level
    StartOffsetMinutes int `json:"startOffsetMinutes"`


    // LengthMinutes - The length of the activity in minutes
    LengthMinutes int `json:"lengthMinutes"`

}

// String returns a JSON representation of the model
func (o *Agentscheduleshiftactivity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentscheduleshiftactivity) MarshalJSON() ([]byte, error) {
    type Alias Agentscheduleshiftactivity

    if AgentscheduleshiftactivityMarshalled {
        return []byte("{}"), nil
    }
    AgentscheduleshiftactivityMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCategory string `json:"activityCategory"`
        
        StartOffsetMinutes int `json:"startOffsetMinutes"`
        
        LengthMinutes int `json:"lengthMinutes"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

