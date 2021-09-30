package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HeadcountintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HeadcountintervalDud struct { 
    


    

}

// Headcountinterval - Headcount interval information for schedule
type Headcountinterval struct { 
    // Interval - The start date-time for this headcount interval in ISO-8601 format, must be within the 8 day schedule
    Interval time.Time `json:"interval"`


    // Value - Headcount value for this interval
    Value float64 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Headcountinterval) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Headcountinterval) MarshalJSON() ([]byte, error) {
    type Alias Headcountinterval

    if HeadcountintervalMarshalled {
        return []byte("{}"), nil
    }
    HeadcountintervalMarshalled = true

    return json.Marshal(&struct { 
        Interval time.Time `json:"interval"`
        
        Value float64 `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

