package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitvaluesDud struct { 
    


    


    

}

// Butimeofflimitvalues
type Butimeofflimitvalues struct { 
    // StartDate - Start date of the requested date range, in ISO-8601 format. The end date is determined by the size of interval lists
    StartDate time.Time `json:"startDate"`


    // ValuesPerDay - Time-off limit values specified in per day granularity. Set only if granularity is 'Daily'
    ValuesPerDay Timeofflimitvalues `json:"valuesPerDay"`


    // ValuesPerFifteenMinutes - Time-off limit values specified in per fifteen minutes granularity. Set only if granularity is 'FifteenMinutes'
    ValuesPerFifteenMinutes Timeofflimitvalues `json:"valuesPerFifteenMinutes"`

}

// String returns a JSON representation of the model
func (o *Butimeofflimitvalues) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitvalues) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitvalues

    if ButimeofflimitvaluesMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitvaluesMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        ValuesPerDay Timeofflimitvalues `json:"valuesPerDay"`
        
        ValuesPerFifteenMinutes Timeofflimitvalues `json:"valuesPerFifteenMinutes"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

