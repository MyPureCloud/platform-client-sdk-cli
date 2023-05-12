package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencequeryforteamsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencequeryforteamsDud struct { 
    


    


    


    


    

}

// Wfmhistoricaladherencequeryforteams
type Wfmhistoricaladherencequeryforteams struct { 
    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
    EndDate time.Time `json:"endDate"`


    // TimeZone - The time zone, in olson format, to use in defining days when computing adherence. The results will be returned as UTC timestamps regardless of the time zone input.
    TimeZone string `json:"timeZone"`


    // UserIds - The userIds to report on. If null or not set, adherence will be computed for all the users in management unit or requested teamIds
    UserIds []string `json:"userIds"`


    // IncludeExceptions - Whether user exceptions should be returned as part of the results
    IncludeExceptions bool `json:"includeExceptions"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequeryforteams) String() string {
    
    
    
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencequeryforteams) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencequeryforteams

    if WfmhistoricaladherencequeryforteamsMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencequeryforteamsMarshalled = true

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

