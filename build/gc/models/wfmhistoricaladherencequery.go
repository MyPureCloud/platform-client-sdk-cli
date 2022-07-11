package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencequeryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencequeryDud struct { 
    


    


    


    


    


    

}

// Wfmhistoricaladherencequery
type Wfmhistoricaladherencequery struct { 
    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
    EndDate time.Time `json:"endDate"`


    // TimeZone - The time zone to use for returned results in olson format. If it is not set, the business unit time zone will be used to compute adherence
    TimeZone string `json:"timeZone"`


    // UserIds - The userIds to report on. If null or not set, adherence will be computed for all the users in management unit or requested teamIds
    UserIds []string `json:"userIds"`


    // IncludeExceptions - Whether user exceptions should be returned as part of the results
    IncludeExceptions bool `json:"includeExceptions"`


    // TeamIds - The teamIds to report on. If null or not set, adherence will be computed for requested users if applicable or otherwise all users in the management unit. Note: If teamIds is also specified, only adherence for users in the requested teams will be returned
    TeamIds []string `json:"teamIds"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequery) String() string {
    
    
    
     o.UserIds = []string{""} 
    
     o.TeamIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencequery) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencequery

    if WfmhistoricaladherencequeryMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencequeryMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        TimeZone string `json:"timeZone"`
        
        UserIds []string `json:"userIds"`
        
        IncludeExceptions bool `json:"includeExceptions"`
        
        TeamIds []string `json:"teamIds"`
        *Alias
    }{

        


        


        


        
        UserIds: []string{""},
        


        


        
        TeamIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

