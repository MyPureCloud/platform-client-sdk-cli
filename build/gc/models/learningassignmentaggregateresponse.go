package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregateresponseDud struct { 
    

}

// Learningassignmentaggregateresponse
type Learningassignmentaggregateresponse struct { 
    // Results - The results of the query
    Results []Learningassignmentaggregatequeryresponsegroupeddata `json:"results"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregateresponse) String() string {
    
    
     o.Results = []Learningassignmentaggregatequeryresponsegroupeddata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregateresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregateresponse

    if LearningassignmentaggregateresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregateresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Learningassignmentaggregatequeryresponsegroupeddata `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Learningassignmentaggregatequeryresponsegroupeddata{{}},
        

        
        Alias: (*Alias)(u),
    })
}

