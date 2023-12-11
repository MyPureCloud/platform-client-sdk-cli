package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionaggregatequeryresponseDud struct { 
    

}

// Flowexecutionaggregatequeryresponse
type Flowexecutionaggregatequeryresponse struct { 
    // Results
    Results []Flowexecutionaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionaggregatequeryresponse) String() string {
     o.Results = []Flowexecutionaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionaggregatequeryresponse

    if FlowexecutionaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Flowexecutionaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Flowexecutionaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

