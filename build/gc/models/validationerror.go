package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationerrorDud struct { 
    


    

}

// Validationerror
type Validationerror struct { 
    // LineNumber - Line number for the error in CSV
    LineNumber int `json:"lineNumber"`


    // Message - Detail of the error in CSV
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Validationerror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationerror) MarshalJSON() ([]byte, error) {
    type Alias Validationerror

    if ValidationerrorMarshalled {
        return []byte("{}"), nil
    }
    ValidationerrorMarshalled = true

    return json.Marshal(&struct {
        
        LineNumber int `json:"lineNumber"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

