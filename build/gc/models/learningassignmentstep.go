package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentstepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentstepDud struct { 
    Id string `json:"id"`


    ModuleStep Learningmoduleinformstep `json:"moduleStep"`


    Structure []Learningassignmentstepscostructure `json:"structure"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    


    PercentageScore float32 `json:"percentageScore"`


    


    SignedCookie Learningassignmentstepsignedcookie `json:"signedCookie"`


    SelfUri string `json:"selfUri"`

}

// Learningassignmentstep - Learning assignment step
type Learningassignmentstep struct { 
    


    


    


    


    


    // CompletionPercentage - The completion percentage for this step
    CompletionPercentage float32 `json:"completionPercentage"`


    


    // ShareableContentObject - The SCO (Shareable Content Object) data
    ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`


    


    

}

// String returns a JSON representation of the model
func (o *Learningassignmentstep) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentstep) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentstep

    if LearningassignmentstepMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentstepMarshalled = true

    return json.Marshal(&struct {
        
        CompletionPercentage float32 `json:"completionPercentage"`
        
        ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

