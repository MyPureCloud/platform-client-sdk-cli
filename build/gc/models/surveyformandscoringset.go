package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyformandscoringsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyformandscoringsetDud struct { 
    


    

}

// Surveyformandscoringset
type Surveyformandscoringset struct { 
    // SurveyForm
    SurveyForm Surveyform `json:"surveyForm"`


    // Answers
    Answers Surveyscoringset `json:"answers"`

}

// String returns a JSON representation of the model
func (o *Surveyformandscoringset) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyformandscoringset) MarshalJSON() ([]byte, error) {
    type Alias Surveyformandscoringset

    if SurveyformandscoringsetMarshalled {
        return []byte("{}"), nil
    }
    SurveyformandscoringsetMarshalled = true

    return json.Marshal(&struct { 
        SurveyForm Surveyform `json:"surveyForm"`
        
        Answers Surveyscoringset `json:"answers"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

