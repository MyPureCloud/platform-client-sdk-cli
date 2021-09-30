package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationquestionscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationquestionscoreDud struct { 
    


    


    


    


    


    

}

// Evaluationquestionscore
type Evaluationquestionscore struct { 
    // QuestionId
    QuestionId string `json:"questionId"`


    // AnswerId
    AnswerId string `json:"answerId"`


    // Score - Unweighted score of the question
    Score int `json:"score"`


    // MarkedNA
    MarkedNA bool `json:"markedNA"`


    // FailedKillQuestion - Applicable only on fatal questions. Indicates that the answer selected was not the highest score available for the question
    FailedKillQuestion bool `json:"failedKillQuestion"`


    // Comments - Comments from the evaluator specific to this question
    Comments string `json:"comments"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationquestionscore) MarshalJSON() ([]byte, error) {
    type Alias Evaluationquestionscore

    if EvaluationquestionscoreMarshalled {
        return []byte("{}"), nil
    }
    EvaluationquestionscoreMarshalled = true

    return json.Marshal(&struct { 
        QuestionId string `json:"questionId"`
        
        AnswerId string `json:"answerId"`
        
        Score int `json:"score"`
        
        MarkedNA bool `json:"markedNA"`
        
        FailedKillQuestion bool `json:"failedKillQuestion"`
        
        Comments string `json:"comments"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

