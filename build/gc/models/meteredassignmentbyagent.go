package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MeteredassignmentbyagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MeteredassignmentbyagentDud struct { 
    


    


    


    


    


    

}

// Meteredassignmentbyagent
type Meteredassignmentbyagent struct { 
    // EvaluationContextId
    EvaluationContextId string `json:"evaluationContextId"`


    // Evaluators
    Evaluators []User `json:"evaluators"`


    // MaxNumberEvaluations
    MaxNumberEvaluations int `json:"maxNumberEvaluations"`


    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // TimeInterval
    TimeInterval Timeinterval `json:"timeInterval"`


    // TimeZone
    TimeZone string `json:"timeZone"`

}

// String returns a JSON representation of the model
func (o *Meteredassignmentbyagent) String() string {
    
    
    
    
    
    
     o.Evaluators = []User{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Meteredassignmentbyagent) MarshalJSON() ([]byte, error) {
    type Alias Meteredassignmentbyagent

    if MeteredassignmentbyagentMarshalled {
        return []byte("{}"), nil
    }
    MeteredassignmentbyagentMarshalled = true

    return json.Marshal(&struct { 
        EvaluationContextId string `json:"evaluationContextId"`
        
        Evaluators []User `json:"evaluators"`
        
        MaxNumberEvaluations int `json:"maxNumberEvaluations"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        TimeInterval Timeinterval `json:"timeInterval"`
        
        TimeZone string `json:"timeZone"`
        
        *Alias
    }{
        

        

        

        
        Evaluators: []User{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

