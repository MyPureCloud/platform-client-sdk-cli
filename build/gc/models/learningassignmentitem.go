package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentitemDud struct { 
    


    

}

// Learningassignmentitem
type Learningassignmentitem struct { 
    // ModuleId - The Learning Module ID associated with this assignment
    ModuleId string `json:"moduleId"`


    // UserId - The User ID associated with this assignment
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentitem) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentitem

    if LearningassignmentitemMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentitemMarshalled = true

    return json.Marshal(&struct {
        
        ModuleId string `json:"moduleId"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

