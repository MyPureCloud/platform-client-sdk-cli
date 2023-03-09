package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentpossibleworkshiftsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentpossibleworkshiftsrequestDud struct { 
    


    

}

// Agentpossibleworkshiftsrequest
type Agentpossibleworkshiftsrequest struct { 
    // WeekStartDate - Start date of requested effective work plan, day of week will be in line with business unit start day of week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekStartDate time.Time `json:"weekStartDate"`


    // WeekCount - Number of weeks for which to return possible work shifts
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Agentpossibleworkshiftsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentpossibleworkshiftsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentpossibleworkshiftsrequest

    if AgentpossibleworkshiftsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentpossibleworkshiftsrequestMarshalled = true

    return json.Marshal(&struct {
        
        WeekStartDate time.Time `json:"weekStartDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

