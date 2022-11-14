package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueryadherenceexplanationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueryadherenceexplanationsrequestDud struct { 
    


    

}

// Agentqueryadherenceexplanationsrequest
type Agentqueryadherenceexplanationsrequest struct { 
    // StartDate - The start date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Agentqueryadherenceexplanationsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueryadherenceexplanationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentqueryadherenceexplanationsrequest

    if AgentqueryadherenceexplanationsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentqueryadherenceexplanationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

