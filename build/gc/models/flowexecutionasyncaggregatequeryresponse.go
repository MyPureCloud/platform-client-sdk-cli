package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionasyncaggregatequeryresponseDud struct { 
    


    

}

// Flowexecutionasyncaggregatequeryresponse
type Flowexecutionasyncaggregatequeryresponse struct { 
    // Results
    Results []Flowexecutionaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionasyncaggregatequeryresponse) String() string {
     o.Results = []Flowexecutionaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionasyncaggregatequeryresponse

    if FlowexecutionasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Flowexecutionaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Flowexecutionaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

