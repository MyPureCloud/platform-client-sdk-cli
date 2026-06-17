package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryeffectiveschedulesetsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryeffectiveschedulesetsrequestDud struct { 
    


    

}

// Queryeffectiveschedulesetsrequest
type Queryeffectiveschedulesetsrequest struct { 
    // StartDate - The start date for querying effective bids relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // WeekCount - The number of weeks to query for effective bids
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Queryeffectiveschedulesetsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryeffectiveschedulesetsrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryeffectiveschedulesetsrequest

    if QueryeffectiveschedulesetsrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryeffectiveschedulesetsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

