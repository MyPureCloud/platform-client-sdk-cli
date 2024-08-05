package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastweekreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastweekreferenceDud struct { 
    


    


    Description string `json:"description"`


    


    SelfUri string `json:"selfUri"`

}

// Bushorttermforecastweekreference
type Bushorttermforecastweekreference struct { 
    // Id - Forecast id used in this work plan bid
    Id string `json:"id"`


    // WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    


    // WeekNumber - The week number used for this bid. First week starts with number 1
    WeekNumber int `json:"weekNumber"`


    

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastweekreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecastweekreference) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecastweekreference

    if BushorttermforecastweekreferenceMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastweekreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekNumber int `json:"weekNumber"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

