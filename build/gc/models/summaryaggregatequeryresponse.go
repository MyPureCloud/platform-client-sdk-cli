package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryaggregatequeryresponseDud struct { 
    

}

// Summaryaggregatequeryresponse
type Summaryaggregatequeryresponse struct { 
    // Results
    Results []Summaryaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Summaryaggregatequeryresponse) String() string {
     o.Results = []Summaryaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Summaryaggregatequeryresponse

    if SummaryaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    SummaryaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Summaryaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Summaryaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

