package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScorablesurveyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScorablesurveyDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Scorablesurvey
type Scorablesurvey struct { 
    


    // Name
    Name string `json:"name"`


    // SurveyForm - Survey form used for this survey.
    SurveyForm Surveyform `json:"surveyForm"`


    // Status
    Status string `json:"status"`


    // Answers
    Answers Surveyscoringset `json:"answers"`


    

}

// String returns a JSON representation of the model
func (o *Scorablesurvey) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scorablesurvey) MarshalJSON() ([]byte, error) {
    type Alias Scorablesurvey

    if ScorablesurveyMarshalled {
        return []byte("{}"), nil
    }
    ScorablesurveyMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SurveyForm Surveyform `json:"surveyForm"`
        
        Status string `json:"status"`
        
        Answers Surveyscoringset `json:"answers"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

