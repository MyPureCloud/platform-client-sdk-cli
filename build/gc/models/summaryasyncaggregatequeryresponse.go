package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryasyncaggregatequeryresponseDud struct { 
    


    

}

// Summaryasyncaggregatequeryresponse
type Summaryasyncaggregatequeryresponse struct { 
    // Results
    Results []Summaryaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Summaryasyncaggregatequeryresponse) String() string {
     o.Results = []Summaryaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Summaryasyncaggregatequeryresponse

    if SummaryasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    SummaryasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Summaryaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Summaryaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

