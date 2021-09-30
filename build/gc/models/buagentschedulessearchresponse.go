package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulessearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulessearchresponseDud struct { 
    


    


    

}

// Buagentschedulessearchresponse
type Buagentschedulessearchresponse struct { 
    // AgentSchedules - The requested agent schedules
    AgentSchedules []Buagentschedulesearchresponse `json:"agentSchedules"`


    // BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
    BusinessUnitTimeZone string `json:"businessUnitTimeZone"`


    // PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
    PublishedSchedules []Buagentschedulepublishedschedulereference `json:"publishedSchedules"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulessearchresponse) String() string {
    
    
     o.AgentSchedules = []Buagentschedulesearchresponse{{}} 
    
    
    
    
    
    
    
     o.PublishedSchedules = []Buagentschedulepublishedschedulereference{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulessearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulessearchresponse

    if BuagentschedulessearchresponseMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulessearchresponseMarshalled = true

    return json.Marshal(&struct { 
        AgentSchedules []Buagentschedulesearchresponse `json:"agentSchedules"`
        
        BusinessUnitTimeZone string `json:"businessUnitTimeZone"`
        
        PublishedSchedules []Buagentschedulepublishedschedulereference `json:"publishedSchedules"`
        
        *Alias
    }{
        

        
        AgentSchedules: []Buagentschedulesearchresponse{{}},
        

        

        

        

        
        PublishedSchedules: []Buagentschedulepublishedschedulereference{{}},
        

        
        Alias: (*Alias)(u),
    })
}

