package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CalibrationassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CalibrationassignmentDud struct { 
    


    


    


    

}

// Calibrationassignment
type Calibrationassignment struct { 
    // Calibrator
    Calibrator User `json:"calibrator"`


    // Evaluators
    Evaluators []User `json:"evaluators"`


    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // ExpertEvaluator
    ExpertEvaluator User `json:"expertEvaluator"`

}

// String returns a JSON representation of the model
func (o *Calibrationassignment) String() string {
    
     o.Evaluators = []User{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Calibrationassignment) MarshalJSON() ([]byte, error) {
    type Alias Calibrationassignment

    if CalibrationassignmentMarshalled {
        return []byte("{}"), nil
    }
    CalibrationassignmentMarshalled = true

    return json.Marshal(&struct {
        
        Calibrator User `json:"calibrator"`
        
        Evaluators []User `json:"evaluators"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        ExpertEvaluator User `json:"expertEvaluator"`
        *Alias
    }{

        


        
        Evaluators: []User{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

