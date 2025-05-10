package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablerowparametervalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablerowparametervalueDud struct { 
    

}

// Decisiontablerowparametervalue
type Decisiontablerowparametervalue struct { 
    // Literal - A literal parameter value
    Literal Literal `json:"literal"`

}

// String returns a JSON representation of the model
func (o *Decisiontablerowparametervalue) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablerowparametervalue) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablerowparametervalue

    if DecisiontablerowparametervalueMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablerowparametervalueMarshalled = true

    return json.Marshal(&struct {
        
        Literal Literal `json:"literal"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

