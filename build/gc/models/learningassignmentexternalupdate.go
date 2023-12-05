package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentexternalupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentexternalupdateDud struct { 
    


    


    

}

// Learningassignmentexternalupdate
type Learningassignmentexternalupdate struct { 
    // State - The Learning Assignment state
    State string `json:"state"`


    // PercentageScore - The score
    PercentageScore float32 `json:"percentageScore"`


    // IsPassed - Was the assignment marked as passed
    IsPassed bool `json:"isPassed"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentexternalupdate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentexternalupdate) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentexternalupdate

    if LearningassignmentexternalupdateMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentexternalupdateMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        PercentageScore float32 `json:"percentageScore"`
        
        IsPassed bool `json:"isPassed"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

