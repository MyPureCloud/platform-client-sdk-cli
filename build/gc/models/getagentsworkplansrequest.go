package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetagentsworkplansrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetagentsworkplansrequestDud struct { 
    


    


    

}

// Getagentsworkplansrequest
type Getagentsworkplansrequest struct { 
    // AgentIds - The list of agent IDs
    AgentIds []string `json:"agentIds"`


    // StartDate - The start of a date in yyyy-MM-dd format. Response contains values rolled back to nearest BU start day of week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // WeekCount - The number of weeks to query
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Getagentsworkplansrequest) String() string {
     o.AgentIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getagentsworkplansrequest) MarshalJSON() ([]byte, error) {
    type Alias Getagentsworkplansrequest

    if GetagentsworkplansrequestMarshalled {
        return []byte("{}"), nil
    }
    GetagentsworkplansrequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentIds []string `json:"agentIds"`
        
        StartDate time.Time `json:"startDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        
        AgentIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

