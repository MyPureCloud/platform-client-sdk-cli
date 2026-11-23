package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KpiattributionaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KpiattributionaggregatequeryresponseDud struct { 
    

}

// Kpiattributionaggregatequeryresponse
type Kpiattributionaggregatequeryresponse struct { 
    // Results
    Results []Kpiattributionaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Kpiattributionaggregatequeryresponse) String() string {
     o.Results = []Kpiattributionaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Kpiattributionaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Kpiattributionaggregatequeryresponse

    if KpiattributionaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    KpiattributionaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Kpiattributionaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Kpiattributionaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

