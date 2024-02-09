package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregatequeryresponseDud struct { 
    

}

// Ratelimitaggregatequeryresponse
type Ratelimitaggregatequeryresponse struct { 
    // Results
    Results []Ratelimitaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregatequeryresponse) String() string {
     o.Results = []Ratelimitaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregatequeryresponse

    if RatelimitaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Ratelimitaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Ratelimitaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

