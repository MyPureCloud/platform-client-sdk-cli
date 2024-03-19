package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitrangeDud struct { 
    


    

}

// Butimeofflimitrange
type Butimeofflimitrange struct { 
    // StartDate - Start date of the range. The end date is determined by the size of 'limitMinutesPerDay'. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // LimitMinutesPerDay - The list of time-off limit values in minutes per day. If 'null' is specified, then the day-specific value is cleared. Such a day will have a value of 0
    LimitMinutesPerDay []int `json:"limitMinutesPerDay"`

}

// String returns a JSON representation of the model
func (o *Butimeofflimitrange) String() string {
    
     o.LimitMinutesPerDay = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitrange) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitrange

    if ButimeofflimitrangeMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitrangeMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        LimitMinutesPerDay []int `json:"limitMinutesPerDay"`
        *Alias
    }{

        


        
        LimitMinutesPerDay: []int{0},
        

        Alias: (*Alias)(u),
    })
}

