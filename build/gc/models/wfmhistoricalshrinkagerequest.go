package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricalshrinkagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricalshrinkagerequestDud struct { 
    


    


    


    

}

// Wfmhistoricalshrinkagerequest
type Wfmhistoricalshrinkagerequest struct { 
    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
    EndDate time.Time `json:"endDate"`


    // TimeZone - The time zone, in olson format, to use in defining days when computing shrinkage for requested granularity. If it is not set, the business unit time zone will be used. The results will be returned as UTC timestamps regardless of the time zone input.
    TimeZone string `json:"timeZone"`


    // Granularity - Shrinkage aggregation interval granularity.
    Granularity string `json:"granularity"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricalshrinkagerequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricalshrinkagerequest) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricalshrinkagerequest

    if WfmhistoricalshrinkagerequestMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricalshrinkagerequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        TimeZone string `json:"timeZone"`
        
        Granularity string `json:"granularity"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

