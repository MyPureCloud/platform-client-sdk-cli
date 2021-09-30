package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentcreateDud struct { 
    


    


    

}

// Learningassignmentcreate
type Learningassignmentcreate struct { 
    // ModuleId - The Learning module Id associated with this assignment
    ModuleId string `json:"moduleId"`


    // UserId - The User for whom the assignment is assigned
    UserId string `json:"userId"`


    // RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RecommendedCompletionDate time.Time `json:"recommendedCompletionDate"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentcreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentcreate) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentcreate

    if LearningassignmentcreateMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentcreateMarshalled = true

    return json.Marshal(&struct { 
        ModuleId string `json:"moduleId"`
        
        UserId string `json:"userId"`
        
        RecommendedCompletionDate time.Time `json:"recommendedCompletionDate"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

