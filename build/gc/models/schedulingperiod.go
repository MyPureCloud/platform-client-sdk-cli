package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingperiodMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingperiodDud struct { 
    


    

}

// Schedulingperiod
type Schedulingperiod struct { 
    // EarliestStartDate - The earliest date the associated activity plan can begin, in YYYY-MM-DD format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EarliestStartDate time.Time `json:"earliestStartDate"`


    // LatestEndDate - The latest date the associated activity plan can end, in YYYY-MM-DD format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    LatestEndDate time.Time `json:"latestEndDate"`

}

// String returns a JSON representation of the model
func (o *Schedulingperiod) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingperiod) MarshalJSON() ([]byte, error) {
    type Alias Schedulingperiod

    if SchedulingperiodMarshalled {
        return []byte("{}"), nil
    }
    SchedulingperiodMarshalled = true

    return json.Marshal(&struct {
        
        EarliestStartDate time.Time `json:"earliestStartDate"`
        
        LatestEndDate time.Time `json:"latestEndDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

