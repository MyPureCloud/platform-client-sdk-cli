package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsroutingstatusrecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsroutingstatusrecordDud struct { 
    


    


    

}

// Analyticsroutingstatusrecord
type Analyticsroutingstatusrecord struct { 
    // StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // RoutingStatus - The user's ACD routing status
    RoutingStatus string `json:"routingStatus"`

}

// String returns a JSON representation of the model
func (o *Analyticsroutingstatusrecord) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsroutingstatusrecord) MarshalJSON() ([]byte, error) {
    type Alias Analyticsroutingstatusrecord

    if AnalyticsroutingstatusrecordMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsroutingstatusrecordMarshalled = true

    return json.Marshal(&struct { 
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        RoutingStatus string `json:"routingStatus"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

