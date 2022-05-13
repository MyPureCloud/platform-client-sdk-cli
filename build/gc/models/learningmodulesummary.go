package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulesummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulesummaryDud struct { 
    


    


    


    

}

// Learningmodulesummary - Learning module summary data
type Learningmodulesummary struct { 
    // AssignedCount - The total number of assignments assigned for a learning module
    AssignedCount int `json:"assignedCount"`


    // CompletedCount - The number of assignments completed for a learning module
    CompletedCount int `json:"completedCount"`


    // PassedCount - The number of assignments passed for a learning module
    PassedCount int `json:"passedCount"`


    // CompletedSum - The sum of assignment scores for a learning module
    CompletedSum float32 `json:"completedSum"`

}

// String returns a JSON representation of the model
func (o *Learningmodulesummary) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulesummary) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulesummary

    if LearningmodulesummaryMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulesummaryMarshalled = true

    return json.Marshal(&struct {
        
        AssignedCount int `json:"assignedCount"`
        
        CompletedCount int `json:"completedCount"`
        
        PassedCount int `json:"passedCount"`
        
        CompletedSum float32 `json:"completedSum"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

