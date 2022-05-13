package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregateresponseDud struct { 
    

}

// Developmentactivityaggregateresponse
type Developmentactivityaggregateresponse struct { 
    // Results - The results of the query
    Results []Developmentactivityaggregatequeryresponsegroupeddata `json:"results"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregateresponse) String() string {
     o.Results = []Developmentactivityaggregatequeryresponsegroupeddata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregateresponse) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregateresponse

    if DevelopmentactivityaggregateresponseMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Developmentactivityaggregatequeryresponsegroupeddata `json:"results"`
        *Alias
    }{

        
        Results: []Developmentactivityaggregatequeryresponsegroupeddata{{}},
        

        Alias: (*Alias)(u),
    })
}

