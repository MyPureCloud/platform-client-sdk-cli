package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdaterequestDud struct { 
    


    


    


    


    

}

// Learningmodulepreviewupdaterequest - Learning module preview update request
type Learningmodulepreviewupdaterequest struct { 
    // State - The assignment State
    State string `json:"state"`


    // CurrentStep - The assignment current step
    CurrentStep Learningmodulepreviewupdaterequestcurrentstep `json:"currentStep"`


    // Steps - The assignment Steps
    Steps []Learningmodulepreviewupdatestep `json:"steps"`


    // Assessment - The assessment for learning module
    Assessment Learningassessment `json:"assessment"`


    // AssessmentForm - The assessment form for learning module
    AssessmentForm Assessmentform `json:"assessmentForm"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdaterequest) String() string {
    
    
     o.Steps = []Learningmodulepreviewupdatestep{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdaterequest

    if LearningmodulepreviewupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        CurrentStep Learningmodulepreviewupdaterequestcurrentstep `json:"currentStep"`
        
        Steps []Learningmodulepreviewupdatestep `json:"steps"`
        
        Assessment Learningassessment `json:"assessment"`
        
        AssessmentForm Assessmentform `json:"assessmentForm"`
        *Alias
    }{

        


        


        
        Steps: []Learningmodulepreviewupdatestep{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

