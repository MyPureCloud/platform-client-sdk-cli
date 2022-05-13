package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MeteredevaluationassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MeteredevaluationassignmentDud struct { 
    


    


    


    


    


    

}

// Meteredevaluationassignment
type Meteredevaluationassignment struct { 
    // EvaluationContextId
    EvaluationContextId string `json:"evaluationContextId"`


    // Evaluators
    Evaluators []User `json:"evaluators"`


    // MaxNumberEvaluations
    MaxNumberEvaluations int `json:"maxNumberEvaluations"`


    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // AssignToActiveUser
    AssignToActiveUser bool `json:"assignToActiveUser"`


    // TimeInterval
    TimeInterval Timeinterval `json:"timeInterval"`

}

// String returns a JSON representation of the model
func (o *Meteredevaluationassignment) String() string {
    
     o.Evaluators = []User{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Meteredevaluationassignment) MarshalJSON() ([]byte, error) {
    type Alias Meteredevaluationassignment

    if MeteredevaluationassignmentMarshalled {
        return []byte("{}"), nil
    }
    MeteredevaluationassignmentMarshalled = true

    return json.Marshal(&struct {
        
        EvaluationContextId string `json:"evaluationContextId"`
        
        Evaluators []User `json:"evaluators"`
        
        MaxNumberEvaluations int `json:"maxNumberEvaluations"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        AssignToActiveUser bool `json:"assignToActiveUser"`
        
        TimeInterval Timeinterval `json:"timeInterval"`
        *Alias
    }{

        


        
        Evaluators: []User{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

