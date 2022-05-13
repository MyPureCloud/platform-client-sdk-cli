package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatoractivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatoractivityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Evaluatoractivity
type Evaluatoractivity struct { 
    


    // Name
    Name string `json:"name"`


    // Evaluator
    Evaluator User `json:"evaluator"`


    // NumEvaluationsAssigned
    NumEvaluationsAssigned int `json:"numEvaluationsAssigned"`


    // NumEvaluationsStarted
    NumEvaluationsStarted int `json:"numEvaluationsStarted"`


    // NumEvaluationsCompleted
    NumEvaluationsCompleted int `json:"numEvaluationsCompleted"`


    // NumCalibrationsAssigned
    NumCalibrationsAssigned int `json:"numCalibrationsAssigned"`


    // NumCalibrationsStarted
    NumCalibrationsStarted int `json:"numCalibrationsStarted"`


    // NumCalibrationsCompleted
    NumCalibrationsCompleted int `json:"numCalibrationsCompleted"`


    // NumEvaluationsWithoutViewPermission
    NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`


    

}

// String returns a JSON representation of the model
func (o *Evaluatoractivity) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatoractivity) MarshalJSON() ([]byte, error) {
    type Alias Evaluatoractivity

    if EvaluatoractivityMarshalled {
        return []byte("{}"), nil
    }
    EvaluatoractivityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Evaluator User `json:"evaluator"`
        
        NumEvaluationsAssigned int `json:"numEvaluationsAssigned"`
        
        NumEvaluationsStarted int `json:"numEvaluationsStarted"`
        
        NumEvaluationsCompleted int `json:"numEvaluationsCompleted"`
        
        NumCalibrationsAssigned int `json:"numCalibrationsAssigned"`
        
        NumCalibrationsStarted int `json:"numCalibrationsStarted"`
        
        NumCalibrationsCompleted int `json:"numCalibrationsCompleted"`
        
        NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

