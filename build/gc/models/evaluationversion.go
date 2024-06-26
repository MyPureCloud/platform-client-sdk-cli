package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationversionDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Evaluationversion
type Evaluationversion struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationversion) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationversion) MarshalJSON() ([]byte, error) {
    type Alias Evaluationversion

    if EvaluationversionMarshalled {
        return []byte("{}"), nil
    }
    EvaluationversionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

