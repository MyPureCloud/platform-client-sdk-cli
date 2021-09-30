package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowaggregatequeryresponseDud struct { 
    

}

// Flowaggregatequeryresponse
type Flowaggregatequeryresponse struct { 
    // Results
    Results []Flowaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryresponse) String() string {
    
    
     o.Results = []Flowaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowaggregatequeryresponse

    if FlowaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Flowaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Flowaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

