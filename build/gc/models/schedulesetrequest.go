package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulesetrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulesetrequestDud struct { 
    


    

}

// Schedulesetrequest
type Schedulesetrequest struct { 
    // ScheduleSetId - The ID of the schedule set
    ScheduleSetId string `json:"scheduleSetId"`


    // OverrideAgentCount - The overridden agent count for the schedule set
    OverrideAgentCount int `json:"overrideAgentCount"`

}

// String returns a JSON representation of the model
func (o *Schedulesetrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulesetrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulesetrequest

    if SchedulesetrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulesetrequestMarshalled = true

    return json.Marshal(&struct {
        
        ScheduleSetId string `json:"scheduleSetId"`
        
        OverrideAgentCount int `json:"overrideAgentCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

