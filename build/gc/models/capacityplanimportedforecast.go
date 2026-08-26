package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanimportedforecastMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanimportedforecastDud struct { 
    


    

}

// Capacityplanimportedforecast
type Capacityplanimportedforecast struct { 
    // WeekDate - The week date of the imported forecast, relative to the business unit time zone, in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks in the imported forecast
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Capacityplanimportedforecast) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanimportedforecast) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanimportedforecast

    if CapacityplanimportedforecastMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanimportedforecastMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

