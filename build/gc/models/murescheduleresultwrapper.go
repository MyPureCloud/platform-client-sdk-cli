package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MurescheduleresultwrapperMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MurescheduleresultwrapperDud struct { 
    

}

// Murescheduleresultwrapper
type Murescheduleresultwrapper struct { 
    // AgentSchedules - The list of agent schedules
    AgentSchedules []Buagentschedulerescheduleresponse `json:"agentSchedules"`

}

// String returns a JSON representation of the model
func (o *Murescheduleresultwrapper) String() string {
     o.AgentSchedules = []Buagentschedulerescheduleresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Murescheduleresultwrapper) MarshalJSON() ([]byte, error) {
    type Alias Murescheduleresultwrapper

    if MurescheduleresultwrapperMarshalled {
        return []byte("{}"), nil
    }
    MurescheduleresultwrapperMarshalled = true

    return json.Marshal(&struct {
        
        AgentSchedules []Buagentschedulerescheduleresponse `json:"agentSchedules"`
        *Alias
    }{

        
        AgentSchedules: []Buagentschedulerescheduleresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

