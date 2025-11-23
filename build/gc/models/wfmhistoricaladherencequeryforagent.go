package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencequeryforagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencequeryforagentDud struct { 
    


    


    

}

// Wfmhistoricaladherencequeryforagent
type Wfmhistoricaladherencequeryforagent struct { 
    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
    EndDate time.Time `json:"endDate"`


    // TimeZone - The time zone, in olson format, to use in defining days when computing adherence. The results will be returned as UTC timestamps regardless of the time zone input.
    TimeZone string `json:"timeZone"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequeryforagent) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencequeryforagent) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencequeryforagent

    if WfmhistoricaladherencequeryforagentMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencequeryforagentMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        TimeZone string `json:"timeZone"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

