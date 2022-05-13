package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitrangeDud struct { 
    


    


    

}

// Timeofflimitrange - A single range filled with time off limit interval values
type Timeofflimitrange struct { 
    // StartDate - Start date of the range. The end date is determined by 'granularity' and the size of 'limitMinutesPerInterval'. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // Granularity - Granularity choice for the time off limit
    Granularity string `json:"granularity"`


    // LimitMinutesPerInterval - The list of time off limit values in minutes per granularity interval. If 'null' is specified, then interval specific value is cleared. Such interval will have 'defaultLimitMinutes' value
    LimitMinutesPerInterval []int `json:"limitMinutesPerInterval"`

}

// String returns a JSON representation of the model
func (o *Timeofflimitrange) String() string {
    
    
     o.LimitMinutesPerInterval = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimitrange) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimitrange

    if TimeofflimitrangeMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitrangeMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        Granularity string `json:"granularity"`
        
        LimitMinutesPerInterval []int `json:"limitMinutesPerInterval"`
        *Alias
    }{

        


        


        
        LimitMinutesPerInterval: []int{0},
        

        Alias: (*Alias)(u),
    })
}

