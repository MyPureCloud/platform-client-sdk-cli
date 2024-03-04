package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationcreateevalformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationcreateevalformDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Evaluationcreateevalform
type Evaluationcreateevalform struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationcreateevalform) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationcreateevalform) MarshalJSON() ([]byte, error) {
    type Alias Evaluationcreateevalform

    if EvaluationcreateevalformMarshalled {
        return []byte("{}"), nil
    }
    EvaluationcreateevalformMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

