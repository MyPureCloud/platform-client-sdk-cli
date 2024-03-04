package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationcreatecalibrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationcreatecalibrationDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Evaluationcreatecalibration
type Evaluationcreatecalibration struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationcreatecalibration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationcreatecalibration) MarshalJSON() ([]byte, error) {
    type Alias Evaluationcreatecalibration

    if EvaluationcreatecalibrationMarshalled {
        return []byte("{}"), nil
    }
    EvaluationcreatecalibrationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

