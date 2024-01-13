package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentquestionscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentquestionscoreDud struct { 
    FailedKillQuestion bool `json:"failedKillQuestion"`


    


    


    


    Score int `json:"score"`


    


    


    

}

// Assessmentquestionscore
type Assessmentquestionscore struct { 
    


    // Comments - Comments provided for the answer
    Comments string `json:"comments"`


    // QuestionId - The ID of the question
    QuestionId string `json:"questionId"`


    // AnswerId - The ID of the selected answer
    AnswerId string `json:"answerId"`


    


    // MarkedNA - True if this question was marked as NA
    MarkedNA bool `json:"markedNA"`


    // SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
    SystemMarkedNA bool `json:"systemMarkedNA"`


    // FreeTextAnswer - Answer for free text answer type
    FreeTextAnswer string `json:"freeTextAnswer"`

}

// String returns a JSON representation of the model
func (o *Assessmentquestionscore) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentquestionscore) MarshalJSON() ([]byte, error) {
    type Alias Assessmentquestionscore

    if AssessmentquestionscoreMarshalled {
        return []byte("{}"), nil
    }
    AssessmentquestionscoreMarshalled = true

    return json.Marshal(&struct {
        
        Comments string `json:"comments"`
        
        QuestionId string `json:"questionId"`
        
        AnswerId string `json:"answerId"`
        
        MarkedNA bool `json:"markedNA"`
        
        SystemMarkedNA bool `json:"systemMarkedNA"`
        
        FreeTextAnswer string `json:"freeTextAnswer"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

