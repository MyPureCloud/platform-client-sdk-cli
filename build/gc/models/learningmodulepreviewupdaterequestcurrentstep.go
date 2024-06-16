package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdaterequestcurrentstepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdaterequestcurrentstepDud struct { 
    


    


    

}

// Learningmodulepreviewupdaterequestcurrentstep - Learning module preview update request current step
type Learningmodulepreviewupdaterequestcurrentstep struct { 
    // Id - The id of this step
    Id string `json:"id"`


    // CompletionPercentage - The completion percentage for this step
    CompletionPercentage float32 `json:"completionPercentage"`


    // ShareableContentObject - The SCO (Shareable Content Object) data
    ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdaterequestcurrentstep) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdaterequestcurrentstep) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdaterequestcurrentstep

    if LearningmodulepreviewupdaterequestcurrentstepMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdaterequestcurrentstepMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        CompletionPercentage float32 `json:"completionPercentage"`
        
        ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

