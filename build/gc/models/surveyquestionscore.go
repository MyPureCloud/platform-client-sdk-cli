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


    // MarkedNA
    MarkedNA bool `json:"markedNA"`


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
        
        AssistedAnswerId string `json:"assistedAnswerId"`
        
        NpsScore int `json:"npsScore"`
        
        NpsTextAnswer string `json:"npsTextAnswer"`
        
        FreeTextAnswer string `json:"freeTextAnswer"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

