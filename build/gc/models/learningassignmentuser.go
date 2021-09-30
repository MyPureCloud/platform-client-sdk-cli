package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentuserDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Learningassignmentuser
type Learningassignmentuser struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Learningassignmentuser) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentuser) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentuser

    if LearningassignmentuserMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentuserMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

