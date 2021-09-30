package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsuserpresencerecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsuserpresencerecordDud struct { 
    


    


    


    

}

// Analyticsuserpresencerecord
type Analyticsuserpresencerecord struct { 
    // StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // SystemPresence - The user's system presence
    SystemPresence string `json:"systemPresence"`


    // OrganizationPresenceId - The identifier for the user's organization presence
    OrganizationPresenceId string `json:"organizationPresenceId"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserpresencerecord) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsuserpresencerecord) MarshalJSON() ([]byte, error) {
    type Alias Analyticsuserpresencerecord

    if AnalyticsuserpresencerecordMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsuserpresencerecordMarshalled = true

    return json.Marshal(&struct { 
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        SystemPresence string `json:"systemPresence"`
        
        OrganizationPresenceId string `json:"organizationPresenceId"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

