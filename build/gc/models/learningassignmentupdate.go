package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentupdateDud struct { 
    


    

}

// Learningassignmentupdate
type Learningassignmentupdate struct { 
    // State - The Learning Assignment state
    State string `json:"state"`


    // Assessment - An updated Assessment
    Assessment Learningassessment `json:"assessment"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentupdate) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentupdate) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentupdate

    if LearningassignmentupdateMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentupdateMarshalled = true

    return json.Marshal(&struct { 
        State string `json:"state"`
        
        Assessment Learningassessment `json:"assessment"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

