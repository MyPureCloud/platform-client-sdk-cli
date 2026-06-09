package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulevisibilityrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulevisibilityrangeDud struct { 
    


    

}

// Schedulevisibilityrange
type Schedulevisibilityrange struct { 
    // EndDate - The schedule visibility end time in ISO-8601, the schedule is visible up to (but not including) that exact time)
    EndDate time.Time `json:"endDate"`


    // EndBusinessUnitDate - The schedule visibility end date interpreted in the business unit time zone in yyyy-MM-dd format, the schedule is visible up to (but not including) that exact date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`

}

// String returns a JSON representation of the model
func (o *Schedulevisibilityrange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulevisibilityrange) MarshalJSON() ([]byte, error) {
    type Alias Schedulevisibilityrange

    if SchedulevisibilityrangeMarshalled {
        return []byte("{}"), nil
    }
    SchedulevisibilityrangeMarshalled = true

    return json.Marshal(&struct {
        
        EndDate time.Time `json:"endDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

