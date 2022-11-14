package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuqueryadherenceexplanationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuqueryadherenceexplanationsrequestDud struct { 
    


    


    


    

}

// Buqueryadherenceexplanationsrequest
type Buqueryadherenceexplanationsrequest struct { 
    // StartDate - The start date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // ManagementUnitIds - A filter for which management units to query. Leave empty or omit entirely for all management units in the business unit
    ManagementUnitIds []string `json:"managementUnitIds"`


    // AgentIds - A filter for which agents within the business unit to query. Leave empty or omit entirely for all agents in the business unit (or management units if specified)
    AgentIds []string `json:"agentIds"`

}

// String returns a JSON representation of the model
func (o *Buqueryadherenceexplanationsrequest) String() string {
    
    
     o.ManagementUnitIds = []string{""} 
     o.AgentIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buqueryadherenceexplanationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buqueryadherenceexplanationsrequest

    if BuqueryadherenceexplanationsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuqueryadherenceexplanationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        AgentIds []string `json:"agentIds"`
        *Alias
    }{

        


        


        
        ManagementUnitIds: []string{""},
        


        
        AgentIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

