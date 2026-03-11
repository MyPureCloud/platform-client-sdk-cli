package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingperiodbaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingperiodbaseDud struct { 
    

}

// Schedulingperiodbase
type Schedulingperiodbase struct { 
    // LatestEndDate - The latest date the associated activity plan can end, in YYYY-MM-DD format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    LatestEndDate time.Time `json:"latestEndDate"`

}

// String returns a JSON representation of the model
func (o *Schedulingperiodbase) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingperiodbase) MarshalJSON() ([]byte, error) {
    type Alias Schedulingperiodbase

    if SchedulingperiodbaseMarshalled {
        return []byte("{}"), nil
    }
    SchedulingperiodbaseMarshalled = true

    return json.Marshal(&struct {
        
        LatestEndDate time.Time `json:"latestEndDate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

