package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowobservationqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowobservationqueryresponseDud struct { 
    

}

// Flowobservationqueryresponse
type Flowobservationqueryresponse struct { 
    // Results
    Results []Flowobservationdatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryresponse) String() string {
    
    
     o.Results = []Flowobservationdatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowobservationqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowobservationqueryresponse

    if FlowobservationqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowobservationqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Flowobservationdatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Flowobservationdatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

