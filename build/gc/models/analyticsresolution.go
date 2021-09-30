package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsresolutionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsresolutionDud struct { 
    


    


    


    

}

// Analyticsresolution
type Analyticsresolution struct { 
    // EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // QueueId - The ID of the last queue on which the conversation was handled.
    QueueId string `json:"queueId"`


    // UserId - The ID of the last user who handled the conversation.
    UserId string `json:"userId"`


    // NNextContactAvoided
    NNextContactAvoided int `json:"nNextContactAvoided"`

}

// String returns a JSON representation of the model
func (o *Analyticsresolution) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsresolution) MarshalJSON() ([]byte, error) {
    type Alias Analyticsresolution

    if AnalyticsresolutionMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsresolutionMarshalled = true

    return json.Marshal(&struct { 
        EventTime time.Time `json:"eventTime"`
        
        QueueId string `json:"queueId"`
        
        UserId string `json:"userId"`
        
        NNextContactAvoided int `json:"nNextContactAvoided"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

