package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentevaluatoractivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentevaluatoractivityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentevaluatoractivity
type Agentevaluatoractivity struct { 
    


    // Name
    Name string `json:"name"`


    // Agent
    Agent User `json:"agent"`


    // Evaluator
    Evaluator User `json:"evaluator"`


    // NumEvaluations
    NumEvaluations int `json:"numEvaluations"`


    // AverageEvaluationScore
    AverageEvaluationScore int `json:"averageEvaluationScore"`


    // NumEvaluationsWithoutViewPermission
    NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`


    

}

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentevaluatoractivity) MarshalJSON() ([]byte, error) {
    type Alias Agentevaluatoractivity

    if AgentevaluatoractivityMarshalled {
        return []byte("{}"), nil
    }
    AgentevaluatoractivityMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Agent User `json:"agent"`
        
        Evaluator User `json:"evaluator"`
        
        NumEvaluations int `json:"numEvaluations"`
        
        AverageEvaluationScore int `json:"averageEvaluationScore"`
        
        NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

