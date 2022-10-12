package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatcheventresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatcheventresponseDud struct { 
    

}

// Batcheventresponse
type Batcheventresponse struct { 
    // Errors - A list of validation or server errors that occurred for posted events.
    Errors []Eventerror `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Batcheventresponse) String() string {
     o.Errors = []Eventerror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batcheventresponse) MarshalJSON() ([]byte, error) {
    type Alias Batcheventresponse

    if BatcheventresponseMarshalled {
        return []byte("{}"), nil
    }
    BatcheventresponseMarshalled = true

    return json.Marshal(&struct {
        
        Errors []Eventerror `json:"errors"`
        *Alias
    }{

        
        Errors: []Eventerror{{}},
        

        Alias: (*Alias)(u),
    })
}

