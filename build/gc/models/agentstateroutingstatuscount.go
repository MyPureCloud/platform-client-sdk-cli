package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateroutingstatuscountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateroutingstatuscountDud struct { 
    


    

}

// Agentstateroutingstatuscount
type Agentstateroutingstatuscount struct { 
    // RoutingStatus - Routing status
    RoutingStatus string `json:"routingStatus"`


    // Count - Count of users with this routing status
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstateroutingstatuscount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateroutingstatuscount) MarshalJSON() ([]byte, error) {
    type Alias Agentstateroutingstatuscount

    if AgentstateroutingstatuscountMarshalled {
        return []byte("{}"), nil
    }
    AgentstateroutingstatuscountMarshalled = true

    return json.Marshal(&struct {
        
        RoutingStatus string `json:"routingStatus"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

