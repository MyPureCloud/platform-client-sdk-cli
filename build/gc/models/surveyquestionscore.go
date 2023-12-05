package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyquestionscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyquestionscoreDud struct { 
    


    


    


    


    


    


    


    


    

}

// Surveyquestionscore
type Surveyquestionscore struct { 
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


    // NpsScore
    NpsScore int `json:"npsScore"`


    // NpsTextAnswer
    NpsTextAnswer string `json:"npsTextAnswer"`


    // FreeTextAnswer
    FreeTextAnswer string `json:"freeTextAnswer"`

}

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyquestionscore) MarshalJSON() ([]byte, error) {
    type Alias Surveyquestionscore

    if SurveyquestionscoreMarshalled {
        return []byte("{}"), nil
    }
    SurveyquestionscoreMarshalled = true

    return json.Marshal(&struct {
        
        QuestionId string `json:"questionId"`
        
        AnswerId string `json:"answerId"`
        
        Score int `json:"score"`
        
        MarkedNA bool `json:"markedNA"`
        
        SystemMarkedNA bool `json:"systemMarkedNA"`
        
        AssistedAnswerId string `json:"assistedAnswerId"`
        
        NpsScore int `json:"npsScore"`
        
        NpsTextAnswer string `json:"npsTextAnswer"`
        
        FreeTextAnswer string `json:"freeTextAnswer"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

