package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowasyncaggregatequeryresponseDud struct { 
    


    

}

// Flowasyncaggregatequeryresponse
type Flowasyncaggregatequeryresponse struct { 
    // Results
    Results []Flowaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Flowasyncaggregatequeryresponse) String() string {
     o.Results = []Flowaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowasyncaggregatequeryresponse

    if FlowasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Flowaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Flowaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

