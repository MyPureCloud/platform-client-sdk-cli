package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationquestiongroupscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationquestiongroupscoreDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Evaluationquestiongroupscore
type Evaluationquestiongroupscore struct { 
    // QuestionGroupId
    QuestionGroupId string `json:"questionGroupId"`


    // TotalScore - Score of all questions in the group
    TotalScore float32 `json:"totalScore"`


    // MaxTotalScore - Maximum possible score of all questions in the group
    MaxTotalScore float32 `json:"maxTotalScore"`


    // MarkedNA - True when the evaluation is submitted with a question group that does not have any answers. Only allowed when naEnabled is true or if set by the system
    MarkedNA bool `json:"markedNA"`


    // SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
    SystemMarkedNA bool `json:"systemMarkedNA"`


    // TotalCriticalScore - Score of only the critical questions in the group
    TotalCriticalScore float32 `json:"totalCriticalScore"`


    // MaxTotalCriticalScore - Maximum possible score of only the critical questions in the group
    MaxTotalCriticalScore float32 `json:"maxTotalCriticalScore"`


    // TotalNonCriticalScore - Score of only the non critical questions in the group
    TotalNonCriticalScore float32 `json:"totalNonCriticalScore"`


    // MaxTotalNonCriticalScore - Maximum possible score of only the non critical questions in the group
    MaxTotalNonCriticalScore float32 `json:"maxTotalNonCriticalScore"`


    // TotalScoreUnweighted - Unweighted score of all questions in the group
    TotalScoreUnweighted float32 `json:"totalScoreUnweighted"`


    // MaxTotalScoreUnweighted - Maximum possible unweighted score of all questions in the group
    MaxTotalScoreUnweighted float32 `json:"maxTotalScoreUnweighted"`


    // TotalCriticalScoreUnweighted - Unweighted score of only the critical questions in the group
    TotalCriticalScoreUnweighted float32 `json:"totalCriticalScoreUnweighted"`


    // MaxTotalCriticalScoreUnweighted - Maximum possible unweighted score of only the critical questions in the group
    MaxTotalCriticalScoreUnweighted float32 `json:"maxTotalCriticalScoreUnweighted"`


    // TotalNonCriticalScoreUnweighted - Unweighted score of only the non critical questions in the group
    TotalNonCriticalScoreUnweighted float32 `json:"totalNonCriticalScoreUnweighted"`


    // MaxTotalNonCriticalScoreUnweighted - Maximum possible unweighted score of only the non critical questions in the group
    MaxTotalNonCriticalScoreUnweighted float32 `json:"maxTotalNonCriticalScoreUnweighted"`


    // QuestionScores
    QuestionScores []Evaluationquestionscore `json:"questionScores"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroupscore) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.QuestionScores = []Evaluationquestionscore{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationquestiongroupscore) MarshalJSON() ([]byte, error) {
    type Alias Evaluationquestiongroupscore

    if EvaluationquestiongroupscoreMarshalled {
        return []byte("{}"), nil
    }
    EvaluationquestiongroupscoreMarshalled = true

    return json.Marshal(&struct {
        
        QuestionGroupId string `json:"questionGroupId"`
        
        TotalScore float32 `json:"totalScore"`
        
        MaxTotalScore float32 `json:"maxTotalScore"`
        
        MarkedNA bool `json:"markedNA"`
        
        SystemMarkedNA bool `json:"systemMarkedNA"`
        
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
        
        QuestionScores []Evaluationquestionscore `json:"questionScores"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        QuestionScores: []Evaluationquestionscore{{}},
        

        Alias: (*Alias)(u),
    })
}

