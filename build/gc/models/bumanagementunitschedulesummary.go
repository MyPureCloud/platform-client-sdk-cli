package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BumanagementunitschedulesummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BumanagementunitschedulesummaryDud struct { 
    


    


    


    


    

}

// Bumanagementunitschedulesummary
type Bumanagementunitschedulesummary struct { 
    // ManagementUnit - The management unit to which this summary applies
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // AgentCount - The number of agents from this management unit that are in the schedule
    AgentCount int `json:"agentCount"`


    // StartDate - The start of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // Agents - The agents in the management unit who are part of this schedule, or in schedule change notifications, the agents that were changed. Note this will come back as an empty list unless the appropriate expand query parameter is passed
    Agents []Userreference `json:"agents"`

}

// String returns a JSON representation of the model
func (o *Bumanagementunitschedulesummary) String() string {
    
    
    
    
     o.Agents = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bumanagementunitschedulesummary) MarshalJSON() ([]byte, error) {
    type Alias Bumanagementunitschedulesummary

    if BumanagementunitschedulesummaryMarshalled {
        return []byte("{}"), nil
    }
    BumanagementunitschedulesummaryMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        AgentCount int `json:"agentCount"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Agents []Userreference `json:"agents"`
        *Alias
    }{

        


        


        


        


        
        Agents: []Userreference{{}},
        

        Alias: (*Alias)(u),
    })
}

