package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulevisibilitysettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulevisibilitysettingsrequestDud struct { 
    


    

}

// Schedulevisibilitysettingsrequest
type Schedulevisibilitysettingsrequest struct { 
    // Enabled - Whether schedule visibility controls are enabled for this management unit
    Enabled bool `json:"enabled"`


    // FutureWeeks - The number of weeks into the future that agents can see schedules in this management unit. 0 means current week
    FutureWeeks int `json:"futureWeeks"`

}

// String returns a JSON representation of the model
func (o *Schedulevisibilitysettingsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulevisibilitysettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulevisibilitysettingsrequest

    if SchedulevisibilitysettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulevisibilitysettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        FutureWeeks int `json:"futureWeeks"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

