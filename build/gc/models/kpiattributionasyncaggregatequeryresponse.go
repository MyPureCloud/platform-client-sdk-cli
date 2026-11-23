package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KpiattributionasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KpiattributionasyncaggregatequeryresponseDud struct { 
    


    

}

// Kpiattributionasyncaggregatequeryresponse
type Kpiattributionasyncaggregatequeryresponse struct { 
    // Results
    Results []Kpiattributionaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Kpiattributionasyncaggregatequeryresponse) String() string {
     o.Results = []Kpiattributionaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Kpiattributionasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Kpiattributionasyncaggregatequeryresponse

    if KpiattributionasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    KpiattributionasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Kpiattributionaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Kpiattributionaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

