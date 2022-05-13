package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulereassignsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulereassignsummaryDud struct { 
    


    


    


    


    

}

// Learningmodulereassignsummary - Learning module reassign summary data
type Learningmodulereassignsummary struct { 
    // TotalReassigned - The total number of users for whom assignment is reassigned
    TotalReassigned int `json:"totalReassigned"`


    // CompletedCount - The total number of users who have the assignment in Completed state
    CompletedCount int `json:"completedCount"`


    // InProgressCount - The total number of users who have the assignment in InProgress state
    InProgressCount int `json:"inProgressCount"`


    // AssignedCount - The total number of users who have the assignment in Assigned state
    AssignedCount int `json:"assignedCount"`


    // NotCompletedCount - The total number of users who have their assignment overdue
    NotCompletedCount int `json:"notCompletedCount"`

}

// String returns a JSON representation of the model
func (o *Learningmodulereassignsummary) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulereassignsummary) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulereassignsummary

    if LearningmodulereassignsummaryMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulereassignsummaryMarshalled = true

    return json.Marshal(&struct {
        
        TotalReassigned int `json:"totalReassigned"`
        
        CompletedCount int `json:"completedCount"`
        
        InProgressCount int `json:"inProgressCount"`
        
        AssignedCount int `json:"assignedCount"`
        
        NotCompletedCount int `json:"notCompletedCount"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

