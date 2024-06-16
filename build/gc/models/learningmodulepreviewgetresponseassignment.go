package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewgetresponseassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewgetresponseassignmentDud struct { 
    


    PercentageScore float32 `json:"percentageScore"`


    AssessmentPercentageScore float32 `json:"assessmentPercentageScore"`


    IsPassed bool `json:"isPassed"`


    AssessmentCompletionPercentage float32 `json:"assessmentCompletionPercentage"`


    CompletionPercentage float32 `json:"completionPercentage"`


    Steps []Learningmodulepreviewgetresponsestep `json:"steps"`

}

// Learningmodulepreviewgetresponseassignment - Learning module preview get response assignment
type Learningmodulepreviewgetresponseassignment struct { 
    // State - The Learning Assignment state
    State string `json:"state"`


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponseassignment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewgetresponseassignment) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewgetresponseassignment

    if LearningmodulepreviewgetresponseassignmentMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewgetresponseassignmentMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

