package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyquestiongroupscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyquestiongroupscoreDud struct { 
    


    


    


    


    

}

// Surveyquestiongroupscore
type Surveyquestiongroupscore struct { 
    // QuestionGroupId
    QuestionGroupId string `json:"questionGroupId"`


    // TotalScore - Score of all questions in the group
    TotalScore float32 `json:"totalScore"`


    // MaxTotalScore - Maximum possible score of all questions in the group
    MaxTotalScore float32 `json:"maxTotalScore"`


    // MarkedNA
    MarkedNA bool `json:"markedNA"`


    // QuestionScores
    QuestionScores []Surveyquestionscore `json:"questionScores"`

}

// String returns a JSON representation of the model
func (o *Surveyquestiongroupscore) String() string {
    
    
    
    
     o.QuestionScores = []Surveyquestionscore{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyquestiongroupscore) MarshalJSON() ([]byte, error) {
    type Alias Surveyquestiongroupscore

    if SurveyquestiongroupscoreMarshalled {
        return []byte("{}"), nil
    }
    SurveyquestiongroupscoreMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupId string `json:"questionGroupId"`
        
        TotalScore float32 `json:"totalScore"`
        
        MaxTotalScore float32 `json:"maxTotalScore"`
        
        MarkedNA bool `json:"markedNA"`
        
        QuestionScores []Surveyquestionscore `json:"questionScores"`
        *Alias
    }{

        


        


        


        


        
        QuestionScores: []Surveyquestionscore{{}},
        

        Alias: (*Alias)(u),
    })
}

