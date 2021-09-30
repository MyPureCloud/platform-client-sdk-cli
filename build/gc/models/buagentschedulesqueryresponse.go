package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulesqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulesqueryresponseDud struct { 
    


    

}

// Buagentschedulesqueryresponse
type Buagentschedulesqueryresponse struct { 
    // AgentSchedules - The requested agent schedules
    AgentSchedules []Buagentschedulequeryresponse `json:"agentSchedules"`


    // BusinessUnitTimeZone - The time zone configured for the business unit to which these schedules apply
    BusinessUnitTimeZone string `json:"businessUnitTimeZone"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulesqueryresponse) String() string {
    
    
     o.AgentSchedules = []Buagentschedulequeryresponse{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulesqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulesqueryresponse

    if BuagentschedulesqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulesqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        AgentSchedules []Buagentschedulequeryresponse `json:"agentSchedules"`
        
        BusinessUnitTimeZone string `json:"businessUnitTimeZone"`
        
        *Alias
    }{
        

        
        AgentSchedules: []Buagentschedulequeryresponse{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

