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
    


    


    


    


    


    


    


    


    AiAnswer Aianswer `json:"aiAnswer"`


    

}

// Evaluationquestionscore
type Evaluationquestionscore struct { 
    // QuestionId
    QuestionId string `json:"questionId"`


    // AnswerId
    AnswerId string `json:"answerId"`


    // Score - Unweighted score of the question
    Score int `json:"score"`


    // MarkedNA - True when the evaluation is submitted with a question that does not have an answer. Only allowed when naEnabled is true or if set by the system
    MarkedNA bool `json:"markedNA"`


    // SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
    SystemMarkedNA bool `json:"systemMarkedNA"`


    // AssistedAnswerId - AnswerId found with evaluation assistance conditions
    AssistedAnswerId string `json:"assistedAnswerId"`


    // FailedKillQuestion - Applicable only on fatal questions. Indicates that the answer selected was not the highest score available for the question
    FailedKillQuestion bool `json:"failedKillQuestion"`


    // Comments - Comments from the evaluator specific to this question
    Comments string `json:"comments"`


    


    // MultipleSelectQuestionOptionScores - Only applicable to Multiple Select questions. Scores corresponding to the options of Multiple Select questions.
    MultipleSelectQuestionOptionScores []Evaluationquestionscore `json:"multipleSelectQuestionOptionScores"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
    
    
    
    
    
    
    
    
     o.MultipleSelectQuestionOptionScores = []Evaluationquestionscore{{}} 

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
        
        SystemMarkedNA bool `json:"systemMarkedNA"`
        
        AssistedAnswerId string `json:"assistedAnswerId"`
        
        FailedKillQuestion bool `json:"failedKillQuestion"`
        
        Comments string `json:"comments"`
        
        MultipleSelectQuestionOptionScores []Evaluationquestionscore `json:"multipleSelectQuestionOptionScores"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        MultipleSelectQuestionOptionScores: []Evaluationquestionscore{{}},
        

        Alias: (*Alias)(u),
    })
}

