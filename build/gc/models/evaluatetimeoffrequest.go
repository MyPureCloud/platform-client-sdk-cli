package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatetimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatetimeoffrequestDud struct { 
    

}

// Evaluatetimeoffrequest
type Evaluatetimeoffrequest struct { 
    // Entities - The list of time-off requests to evaluate
    Entities []Evaluatetimeoffautoapproval `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Evaluatetimeoffrequest) String() string {
     o.Entities = []Evaluatetimeoffautoapproval{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatetimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Evaluatetimeoffrequest

    if EvaluatetimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    EvaluatetimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Evaluatetimeoffautoapproval `json:"entities"`
        *Alias
    }{

        
        Entities: []Evaluatetimeoffautoapproval{{}},
        

        Alias: (*Alias)(u),
    })
}

