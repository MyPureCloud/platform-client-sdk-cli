package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewgetresponsestepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewgetresponsestepDud struct { 
    Id string `json:"id"`


    ModuleStep Learningmoduleinformstep `json:"moduleStep"`


    Structure []Learningmodulepreviewgetscostructure `json:"structure"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    


    PercentageScore float32 `json:"percentageScore"`


    SignedCookie Learningassignmentstepsignedcookie `json:"signedCookie"`

}

// Learningmodulepreviewgetresponsestep - Learning module preview get response assignment step
type Learningmodulepreviewgetresponsestep struct { 
    


    


    


    


    


    // CompletionPercentage - The completion percentage for this step
    CompletionPercentage float32 `json:"completionPercentage"`


    


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponsestep) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewgetresponsestep) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewgetresponsestep

    if LearningmodulepreviewgetresponsestepMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewgetresponsestepMarshalled = true

    return json.Marshal(&struct {
        
        CompletionPercentage float32 `json:"completionPercentage"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

