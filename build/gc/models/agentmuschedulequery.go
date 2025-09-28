package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmuschedulequeryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmuschedulequeryDud struct { 
    


    

}

// Agentmuschedulequery
type Agentmuschedulequery struct { 
    // StartDate - The start date for the range to query the schedule's for the agent's management unit. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date for the range to query the schedule's for the agent's management unit. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Agentmuschedulequery) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmuschedulequery) MarshalJSON() ([]byte, error) {
    type Alias Agentmuschedulequery

    if AgentmuschedulequeryMarshalled {
        return []byte("{}"), nil
    }
    AgentmuschedulequeryMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

