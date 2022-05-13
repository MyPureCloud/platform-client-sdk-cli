package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingstatusDud struct { 
    


    


    

}

// Routingstatus
type Routingstatus struct { 
    // UserId - The userId of the agent
    UserId string `json:"userId"`


    // Status - Indicates the Routing State of the agent.  A value of OFF_QUEUE will be returned if the specified user does not exist.
    Status string `json:"status"`


    // StartTime - The timestamp when the agent went into this state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`

}

// String returns a JSON representation of the model
func (o *Routingstatus) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingstatus) MarshalJSON() ([]byte, error) {
    type Alias Routingstatus

    if RoutingstatusMarshalled {
        return []byte("{}"), nil
    }
    RoutingstatusMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        Status string `json:"status"`
        
        StartTime time.Time `json:"startTime"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

