package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdatestepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdatestepDud struct { 
    Id string `json:"id"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    CompletionPercentage float32 `json:"completionPercentage"`


    PercentageScore float32 `json:"percentageScore"`


    Structure []Learningmodulepreviewupdatescostructure `json:"structure"`

}

// Learningmodulepreviewupdatestep - Learning module preview update assignment step
type Learningmodulepreviewupdatestep struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdatestep) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdatestep) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdatestep

    if LearningmodulepreviewupdatestepMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdatestepMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

