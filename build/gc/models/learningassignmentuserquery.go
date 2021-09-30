package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentuserqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentuserqueryDud struct { 
    


    

}

// Learningassignmentuserquery - Learning module users query request model
type Learningassignmentuserquery struct { 
    // Rule - Learning module rule object
    Rule Learningmodulerule `json:"rule"`


    // SearchTerm - The user name to be searched for
    SearchTerm string `json:"searchTerm"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentuserquery) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentuserquery) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentuserquery

    if LearningassignmentuserqueryMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentuserqueryMarshalled = true

    return json.Marshal(&struct { 
        Rule Learningmodulerule `json:"rule"`
        
        SearchTerm string `json:"searchTerm"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

