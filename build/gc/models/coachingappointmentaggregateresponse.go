package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentaggregateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentaggregateresponseDud struct { 
    

}

// Coachingappointmentaggregateresponse
type Coachingappointmentaggregateresponse struct { 
    // Results - The results of the query
    Results []Queryresponsegroupeddata `json:"results"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregateresponse) String() string {
    
    
     o.Results = []Queryresponsegroupeddata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentaggregateresponse) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentaggregateresponse

    if CoachingappointmentaggregateresponseMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentaggregateresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Queryresponsegroupeddata `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Queryresponsegroupeddata{{}},
        

        
        Alias: (*Alias)(u),
    })
}

