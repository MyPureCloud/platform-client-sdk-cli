package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationassignmentDud struct { 
    


    

}

// Evaluationassignment
type Evaluationassignment struct { 
    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // User
    User User `json:"user"`

}

// String returns a JSON representation of the model
func (o *Evaluationassignment) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationassignment) MarshalJSON() ([]byte, error) {
    type Alias Evaluationassignment

    if EvaluationassignmentMarshalled {
        return []byte("{}"), nil
    }
    EvaluationassignmentMarshalled = true

    return json.Marshal(&struct { 
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        User User `json:"user"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

