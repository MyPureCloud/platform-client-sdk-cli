package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdayperiodMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdayperiodDud struct { 
    


    

}

// Workdayperiod
type Workdayperiod struct { 
    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`

}

// String returns a JSON representation of the model
func (o *Workdayperiod) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdayperiod) MarshalJSON() ([]byte, error) {
    type Alias Workdayperiod

    if WorkdayperiodMarshalled {
        return []byte("{}"), nil
    }
    WorkdayperiodMarshalled = true

    return json.Marshal(&struct {
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

