package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdateresponsecurrentstepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdateresponsecurrentstepDud struct { 
    

}

// Learningmodulepreviewupdateresponsecurrentstep - Learning module preview update response current step
type Learningmodulepreviewupdateresponsecurrentstep struct { 
    // ShareableContentObject - The SCO (Shareable Content Object) data
    ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdateresponsecurrentstep) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdateresponsecurrentstep) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdateresponsecurrentstep

    if LearningmodulepreviewupdateresponsecurrentstepMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdateresponsecurrentstepMarshalled = true

    return json.Marshal(&struct {
        
        ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

