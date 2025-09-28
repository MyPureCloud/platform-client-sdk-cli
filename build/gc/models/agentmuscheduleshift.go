package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmuscheduleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmuscheduleshiftDud struct { 
    


    


    

}

// Agentmuscheduleshift
type Agentmuscheduleshift struct { 
    // StartOffsetMinutes - The start offset of the shift, relative to referenceStartDate at the top level
    StartOffsetMinutes int `json:"startOffsetMinutes"`


    // LengthMinutes - The length of the shift in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Activities - The activities associated with this shift
    Activities []Agentscheduleshiftactivity `json:"activities"`

}

// String returns a JSON representation of the model
func (o *Agentmuscheduleshift) String() string {
    
    
     o.Activities = []Agentscheduleshiftactivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmuscheduleshift) MarshalJSON() ([]byte, error) {
    type Alias Agentmuscheduleshift

    if AgentmuscheduleshiftMarshalled {
        return []byte("{}"), nil
    }
    AgentmuscheduleshiftMarshalled = true

    return json.Marshal(&struct {
        
        StartOffsetMinutes int `json:"startOffsetMinutes"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Activities []Agentscheduleshiftactivity `json:"activities"`
        *Alias
    }{

        


        


        
        Activities: []Agentscheduleshiftactivity{{}},
        

        Alias: (*Alias)(u),
    })
}

