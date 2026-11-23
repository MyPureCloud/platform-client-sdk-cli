package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatetimeoffjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatetimeoffjobresponseDud struct { 
    

}

// Evaluatetimeoffjobresponse
type Evaluatetimeoffjobresponse struct { 
    // Job - The asynchronous job handling the evaluate request
    Job Evaluatetimeoffjobreference `json:"job"`

}

// String returns a JSON representation of the model
func (o *Evaluatetimeoffjobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatetimeoffjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluatetimeoffjobresponse

    if EvaluatetimeoffjobresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluatetimeoffjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Evaluatetimeoffjobreference `json:"job"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

