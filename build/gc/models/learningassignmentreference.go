package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Learningassignmentreference
type Learningassignmentreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Learningassignmentreference) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentreference) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentreference

    if LearningassignmentreferenceMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

