package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Evaluationreference
type Evaluationreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Evaluationreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationreference) MarshalJSON() ([]byte, error) {
    type Alias Evaluationreference

    if EvaluationreferenceMarshalled {
        return []byte("{}"), nil
    }
    EvaluationreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

