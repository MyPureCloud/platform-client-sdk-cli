package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyasyncaggregatequeryresponseDud struct { 
    


    

}

// Journeyasyncaggregatequeryresponse
type Journeyasyncaggregatequeryresponse struct { 
    // Results
    Results []Journeyaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Journeyasyncaggregatequeryresponse) String() string {
     o.Results = []Journeyaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Journeyasyncaggregatequeryresponse

    if JourneyasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    JourneyasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Journeyaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Journeyaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

