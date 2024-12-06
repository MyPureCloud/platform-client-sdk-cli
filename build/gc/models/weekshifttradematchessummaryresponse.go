package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekshifttradematchessummaryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekshifttradematchessummaryresponseDud struct { 
    


    

}

// Weekshifttradematchessummaryresponse
type Weekshifttradematchessummaryresponse struct { 
    // WeekDate - The schedule week date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // Count - The number of trades in the Matched state with the initiating shift in the given week
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Weekshifttradematchessummaryresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekshifttradematchessummaryresponse) MarshalJSON() ([]byte, error) {
    type Alias Weekshifttradematchessummaryresponse

    if WeekshifttradematchessummaryresponseMarshalled {
        return []byte("{}"), nil
    }
    WeekshifttradematchessummaryresponseMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

