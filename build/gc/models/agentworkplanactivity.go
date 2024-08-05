package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanactivityDud struct { 
    


    

}

// Agentworkplanactivity
type Agentworkplanactivity struct { 
    // LengthMinutes - Length of the activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // CountsAsPaidTime - Whether the activity is paid
    CountsAsPaidTime bool `json:"countsAsPaidTime"`

}

// String returns a JSON representation of the model
func (o *Agentworkplanactivity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanactivity) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanactivity

    if AgentworkplanactivityMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanactivityMarshalled = true

    return json.Marshal(&struct {
        
        LengthMinutes int `json:"lengthMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

