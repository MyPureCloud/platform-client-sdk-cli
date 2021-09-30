package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassessmentscoringrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassessmentscoringrequestDud struct { 
    


    

}

// Learningassessmentscoringrequest
type Learningassessmentscoringrequest struct { 
    // AssessmentForm - The assessment form to score against
    AssessmentForm Assessmentform `json:"assessmentForm"`


    // Answers - The answers to score
    Answers Assessmentscoringset `json:"answers"`

}

// String returns a JSON representation of the model
func (o *Learningassessmentscoringrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassessmentscoringrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningassessmentscoringrequest

    if LearningassessmentscoringrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningassessmentscoringrequestMarshalled = true

    return json.Marshal(&struct { 
        AssessmentForm Assessmentform `json:"assessmentForm"`
        
        Answers Assessmentscoringset `json:"answers"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

