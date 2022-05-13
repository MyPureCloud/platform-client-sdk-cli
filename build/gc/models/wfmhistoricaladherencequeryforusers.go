package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencequeryforusersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencequeryforusersDud struct { 
    


    


    


    


    

}

// Wfmhistoricaladherencequeryforusers
type Wfmhistoricaladherencequeryforusers struct { 
    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
    EndDate time.Time `json:"endDate"`


    // TimeZone - The time zone to use for returned results in olson format. If it is not set, the business unit time zone will be used to compute adherence
    TimeZone string `json:"timeZone"`


    // UserIds - The userIds to report on
    UserIds []string `json:"userIds"`


    // IncludeExceptions - Whether user exceptions should be returned as part of the results
    IncludeExceptions bool `json:"includeExceptions"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequeryforusers) String() string {
    
    
    
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencequeryforusers) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencequeryforusers

    if WfmhistoricaladherencequeryforusersMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencequeryforusersMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        TimeZone string `json:"timeZone"`
        
        UserIds []string `json:"userIds"`
        
        IncludeExceptions bool `json:"includeExceptions"`
        *Alias
    }{

        


        


        


        
        UserIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

