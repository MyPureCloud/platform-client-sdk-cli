package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeweeksummaryresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeweeksummaryresponseitemDud struct { 
    


    


    

}

// Shifttradeweeksummaryresponseitem
type Shifttradeweeksummaryresponseitem struct { 
    // WeekDate - The schedule week date in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // InitiatingMatchedCount - The number of trades in the 'Matched' state with the initiating shift in the given week
    InitiatingMatchedCount int `json:"initiatingMatchedCount"`


    // CrossWeekReceivingMatchedCount - The number of cross-week trades in the 'Matched' state with the receiving shift for the given week
    CrossWeekReceivingMatchedCount int `json:"crossWeekReceivingMatchedCount"`

}

// String returns a JSON representation of the model
func (o *Shifttradeweeksummaryresponseitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeweeksummaryresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeweeksummaryresponseitem

    if ShifttradeweeksummaryresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeweeksummaryresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        InitiatingMatchedCount int `json:"initiatingMatchedCount"`
        
        CrossWeekReceivingMatchedCount int `json:"crossWeekReceivingMatchedCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

