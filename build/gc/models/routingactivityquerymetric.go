package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivityquerymetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivityquerymetricDud struct { 
    


    

}

// Routingactivityquerymetric
type Routingactivityquerymetric struct { 
    // Metric - The requested metric
    Metric string `json:"metric"`


    // Details - Flag for including observation details for this metric in the response
    Details bool `json:"details"`

}

// String returns a JSON representation of the model
func (o *Routingactivityquerymetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivityquerymetric) MarshalJSON() ([]byte, error) {
    type Alias Routingactivityquerymetric

    if RoutingactivityquerymetricMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivityquerymetricMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Details bool `json:"details"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

