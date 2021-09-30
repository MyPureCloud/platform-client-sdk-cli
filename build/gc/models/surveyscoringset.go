package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyscoringsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyscoringsetDud struct { 
    


    


    

}

// Surveyscoringset
type Surveyscoringset struct { 
    // TotalScore
    TotalScore float32 `json:"totalScore"`


    // NpsScore
    NpsScore int `json:"npsScore"`


    // QuestionGroupScores
    QuestionGroupScores []Surveyquestiongroupscore `json:"questionGroupScores"`

}

// String returns a JSON representation of the model
func (o *Surveyscoringset) String() string {
    
    
    
    
    
    
    
    
    
    
     o.QuestionGroupScores = []Surveyquestiongroupscore{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyscoringset) MarshalJSON() ([]byte, error) {
    type Alias Surveyscoringset

    if SurveyscoringsetMarshalled {
        return []byte("{}"), nil
    }
    SurveyscoringsetMarshalled = true

    return json.Marshal(&struct { 
        TotalScore float32 `json:"totalScore"`
        
        NpsScore int `json:"npsScore"`
        
        QuestionGroupScores []Surveyquestiongroupscore `json:"questionGroupScores"`
        
        *Alias
    }{
        

        

        

        

        

        
        QuestionGroupScores: []Surveyquestiongroupscore{{}},
        

        
        Alias: (*Alias)(u),
    })
}

