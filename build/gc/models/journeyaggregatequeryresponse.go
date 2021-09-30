package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyaggregatequeryresponseDud struct { 
    

}

// Journeyaggregatequeryresponse
type Journeyaggregatequeryresponse struct { 
    // Results
    Results []Journeyaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryresponse) String() string {
    
    
     o.Results = []Journeyaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Journeyaggregatequeryresponse

    if JourneyaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    JourneyaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Journeyaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Journeyaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

