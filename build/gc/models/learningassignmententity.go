package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmententityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmententityDud struct { 
    

}

// Learningassignmententity
type Learningassignmententity struct { 
    // AssignmentId
    AssignmentId string `json:"assignmentId"`

}

// String returns a JSON representation of the model
func (o *Learningassignmententity) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmententity) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmententity

    if LearningassignmententityMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmententityMarshalled = true

    return json.Marshal(&struct { 
        AssignmentId string `json:"assignmentId"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

