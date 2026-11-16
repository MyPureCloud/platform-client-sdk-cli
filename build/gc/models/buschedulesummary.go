package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulesummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulesummaryDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Buschedulesummary
type Buschedulesummary struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks spanned by this schedule
    WeekCount int `json:"weekCount"`


    // Published - Whether this schedule is published
    Published bool `json:"published"`


    

}

// String returns a JSON representation of the model
func (o *Buschedulesummary) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulesummary) MarshalJSON() ([]byte, error) {
    type Alias Buschedulesummary

    if BuschedulesummaryMarshalled {
        return []byte("{}"), nil
    }
    BuschedulesummaryMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        
        Published bool `json:"published"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

