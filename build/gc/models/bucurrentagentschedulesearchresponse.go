package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BucurrentagentschedulesearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BucurrentagentschedulesearchresponseDud struct { 
    


    


    


    


    


    

}

// Bucurrentagentschedulesearchresponse
type Bucurrentagentschedulesearchresponse struct { 
    // AgentSchedules - The requested agent schedules
    AgentSchedules []Buagentschedulesearchresponse `json:"agentSchedules"`


    // BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
    BusinessUnitTimeZone string `json:"businessUnitTimeZone"`


    // PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
    PublishedSchedules []Buagentschedulepublishedschedulereference `json:"publishedSchedules"`


    // StartDate - The start date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // Updates - The list of updates for the schedule. Only used in notifications
    Updates []Buagentscheduleupdate `json:"updates"`

}

// String returns a JSON representation of the model
func (o *Bucurrentagentschedulesearchresponse) String() string {
    
    
     o.AgentSchedules = []Buagentschedulesearchresponse{{}} 
    
    
    
    
    
    
    
     o.PublishedSchedules = []Buagentschedulepublishedschedulereference{{}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Updates = []Buagentscheduleupdate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bucurrentagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Bucurrentagentschedulesearchresponse

    if BucurrentagentschedulesearchresponseMarshalled {
        return []byte("{}"), nil
    }
    BucurrentagentschedulesearchresponseMarshalled = true

    return json.Marshal(&struct { 
        AgentSchedules []Buagentschedulesearchresponse `json:"agentSchedules"`
        
        BusinessUnitTimeZone string `json:"businessUnitTimeZone"`
        
        PublishedSchedules []Buagentschedulepublishedschedulereference `json:"publishedSchedules"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Updates []Buagentscheduleupdate `json:"updates"`
        
        *Alias
    }{
        

        
        AgentSchedules: []Buagentschedulesearchresponse{{}},
        

        

        

        

        
        PublishedSchedules: []Buagentschedulepublishedschedulereference{{}},
        

        

        

        

        

        

        
        Updates: []Buagentscheduleupdate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

