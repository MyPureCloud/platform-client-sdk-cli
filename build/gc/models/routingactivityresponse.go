package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivityresponseDud struct { 
    


    

}

// Routingactivityresponse
type Routingactivityresponse struct { 
    // Results - Query results
    Results []Routingactivitydata `json:"results"`


    // EntityIdDimension - Dimension that is used as an entityId
    EntityIdDimension string `json:"entityIdDimension"`

}

// String returns a JSON representation of the model
func (o *Routingactivityresponse) String() string {
     o.Results = []Routingactivitydata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivityresponse) MarshalJSON() ([]byte, error) {
    type Alias Routingactivityresponse

    if RoutingactivityresponseMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivityresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Routingactivitydata `json:"results"`
        
        EntityIdDimension string `json:"entityIdDimension"`
        *Alias
    }{

        
        Results: []Routingactivitydata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

