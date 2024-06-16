package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdateresponseassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdateresponseassignmentDud struct { 
    


    PercentageScore float32 `json:"percentageScore"`


    CompletionPercentage float32 `json:"completionPercentage"`


    AssessmentPercentageScore float32 `json:"assessmentPercentageScore"`


    AssessmentCompletionPercentage float32 `json:"assessmentCompletionPercentage"`


    IsPassed bool `json:"isPassed"`


    CurrentStep Learningmodulepreviewupdateresponsecurrentstep `json:"currentStep"`


    Steps []Learningmodulepreviewupdatestep `json:"steps"`

}

// Learningmodulepreviewupdateresponseassignment - Learning module preview update response assignment
type Learningmodulepreviewupdateresponseassignment struct { 
    // State - The Learning Assignment state
    State string `json:"state"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdateresponseassignment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdateresponseassignment) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdateresponseassignment

    if LearningmodulepreviewupdateresponseassignmentMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdateresponseassignmentMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

