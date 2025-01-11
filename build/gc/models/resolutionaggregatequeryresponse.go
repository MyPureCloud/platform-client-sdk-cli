package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutionaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutionaggregatequeryresponseDud struct { 
    

}

// Resolutionaggregatequeryresponse
type Resolutionaggregatequeryresponse struct { 
    // Results
    Results []Resolutionaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Resolutionaggregatequeryresponse) String() string {
     o.Results = []Resolutionaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutionaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Resolutionaggregatequeryresponse

    if ResolutionaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ResolutionaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Resolutionaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Resolutionaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

