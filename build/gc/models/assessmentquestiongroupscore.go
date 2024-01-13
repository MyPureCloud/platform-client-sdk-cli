package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentquestiongroupscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentquestiongroupscoreDud struct { 
    


    TotalScore float32 `json:"totalScore"`


    MaxTotalScore float32 `json:"maxTotalScore"`


    


    


    TotalCriticalScore float32 `json:"totalCriticalScore"`


    MaxTotalCriticalScore float32 `json:"maxTotalCriticalScore"`


    TotalNonCriticalScore float32 `json:"totalNonCriticalScore"`


    MaxTotalNonCriticalScore float32 `json:"maxTotalNonCriticalScore"`


    TotalScoreUnweighted float32 `json:"totalScoreUnweighted"`


    MaxTotalScoreUnweighted float32 `json:"maxTotalScoreUnweighted"`


    TotalCriticalScoreUnweighted float32 `json:"totalCriticalScoreUnweighted"`


    MaxTotalCriticalScoreUnweighted float32 `json:"maxTotalCriticalScoreUnweighted"`


    TotalNonCriticalScoreUnweighted float32 `json:"totalNonCriticalScoreUnweighted"`


    MaxTotalNonCriticalScoreUnweighted float32 `json:"maxTotalNonCriticalScoreUnweighted"`


    

}

// Assessmentquestiongroupscore
type Assessmentquestiongroupscore struct { 
    // QuestionGroupId - The ID of the question group
    QuestionGroupId string `json:"questionGroupId"`


    


    


    // MarkedNA - True if this question group is marked NA
    MarkedNA bool `json:"markedNA"`


    // SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
    SystemMarkedNA bool `json:"systemMarkedNA"`


    


    


    


    


    


    


    


    


    


    


    // QuestionScores - The individual question scores
    QuestionScores []Assessmentquestionscore `json:"questionScores"`

}

// String returns a JSON representation of the model
func (o *Assessmentquestiongroupscore) String() string {
    
    
    
     o.QuestionScores = []Assessmentquestionscore{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentquestiongroupscore) MarshalJSON() ([]byte, error) {
    type Alias Assessmentquestiongroupscore

    if AssessmentquestiongroupscoreMarshalled {
        return []byte("{}"), nil
    }
    AssessmentquestiongroupscoreMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupId string `json:"questionGroupId"`
        
        MarkedNA bool `json:"markedNA"`
        
        SystemMarkedNA bool `json:"systemMarkedNA"`
        
        QuestionScores []Assessmentquestionscore `json:"questionScores"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        QuestionScores: []Assessmentquestionscore{{}},
        

        Alias: (*Alias)(u),
    })
}

