package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyasyncaggregatequeryresponseDud struct { 
    


    

}

// Surveyasyncaggregatequeryresponse
type Surveyasyncaggregatequeryresponse struct { 
    // Results
    Results []Surveyaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Surveyasyncaggregatequeryresponse) String() string {
     o.Results = []Surveyaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Surveyasyncaggregatequeryresponse

    if SurveyasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    SurveyasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Surveyaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Surveyaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

