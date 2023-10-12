package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutionasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutionasyncaggregatequeryresponseDud struct { 
    


    

}

// Resolutionasyncaggregatequeryresponse
type Resolutionasyncaggregatequeryresponse struct { 
    // Results
    Results []Resolutionaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Resolutionasyncaggregatequeryresponse) String() string {
     o.Results = []Resolutionaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutionasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Resolutionasyncaggregatequeryresponse

    if ResolutionasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ResolutionasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Resolutionaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Resolutionaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

