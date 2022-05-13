package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassessmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassessmentDud struct { 
    AssessmentId string `json:"assessmentId"`


    ContextId string `json:"contextId"`


    AssessmentFormId string `json:"assessmentFormId"`


    Status string `json:"status"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DateSubmitted time.Time `json:"dateSubmitted"`

}

// Learningassessment
type Learningassessment struct { 
    


    


    


    


    // Answers - Answers for the assessment
    Answers Assessmentscoringset `json:"answers"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningassessment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassessment) MarshalJSON() ([]byte, error) {
    type Alias Learningassessment

    if LearningassessmentMarshalled {
        return []byte("{}"), nil
    }
    LearningassessmentMarshalled = true

    return json.Marshal(&struct {
        
        Answers Assessmentscoringset `json:"answers"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

