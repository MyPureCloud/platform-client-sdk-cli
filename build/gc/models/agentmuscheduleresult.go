package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmuscheduleresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmuscheduleresultDud struct { 
    


    

}

// Agentmuscheduleresult
type Agentmuscheduleresult struct { 
    // ReferenceStartDate - The reference start date to use when calculating offsets for shifts and activities. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`


    // AgentSchedules - The list of agent schedules for the management unit
    AgentSchedules []Agentmuscheduleitem `json:"agentSchedules"`

}

// String returns a JSON representation of the model
func (o *Agentmuscheduleresult) String() string {
    
     o.AgentSchedules = []Agentmuscheduleitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmuscheduleresult) MarshalJSON() ([]byte, error) {
    type Alias Agentmuscheduleresult

    if AgentmuscheduleresultMarshalled {
        return []byte("{}"), nil
    }
    AgentmuscheduleresultMarshalled = true

    return json.Marshal(&struct {
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        
        AgentSchedules []Agentmuscheduleitem `json:"agentSchedules"`
        *Alias
    }{

        


        
        AgentSchedules: []Agentmuscheduleitem{{}},
        

        Alias: (*Alias)(u),
    })
}

