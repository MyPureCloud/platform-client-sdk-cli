package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingintervalDud struct { 
    


    

}

// Reportinginterval
type Reportinginterval struct { 
    // IntervalType - The granularity of the reporting interval period
    IntervalType string `json:"intervalType"`


    // IntervalValue - The value of the reporting interval period for a given interval type
    IntervalValue int `json:"intervalValue"`

}

// String returns a JSON representation of the model
func (o *Reportinginterval) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportinginterval) MarshalJSON() ([]byte, error) {
    type Alias Reportinginterval

    if ReportingintervalMarshalled {
        return []byte("{}"), nil
    }
    ReportingintervalMarshalled = true

    return json.Marshal(&struct {
        
        IntervalType string `json:"intervalType"`
        
        IntervalValue int `json:"intervalValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

