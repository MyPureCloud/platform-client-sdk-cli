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
    


    


    


    


    


    

}

// Learningmodulepreviewupdatestep - Learning module preview update assignment step
type Learningmodulepreviewupdatestep struct { 
    // Id - The id of the step
    Id string `json:"id"`


    // SuccessStatus - The success status of the step
    SuccessStatus string `json:"successStatus"`


    // CompletionStatus - The completion status of the step
    CompletionStatus string `json:"completionStatus"`


    // CompletionPercentage - The completion percentage of the step
    CompletionPercentage float32 `json:"completionPercentage"`


    // PercentageScore - Percentage Score
    PercentageScore float32 `json:"percentageScore"`


    // Structure - The structure for any SCO associated with this step
    Structure []Learningmodulepreviewupdatescostructure `json:"structure"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdatestep) String() string {
    
    
    
    
    
     o.Structure = []Learningmodulepreviewupdatescostructure{{}} 

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
        
        Id string `json:"id"`
        
        SuccessStatus string `json:"successStatus"`
        
        CompletionStatus string `json:"completionStatus"`
        
        CompletionPercentage float32 `json:"completionPercentage"`
        
        PercentageScore float32 `json:"percentageScore"`
        
        Structure []Learningmodulepreviewupdatescostructure `json:"structure"`
        *Alias
    }{

        


        


        


        


        


        
        Structure: []Learningmodulepreviewupdatescostructure{{}},
        

        Alias: (*Alias)(u),
    })
}

