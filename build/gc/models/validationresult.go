package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationresultDud struct { 
    


    


    

}

// Validationresult
type Validationresult struct { 
    // SeparatorValid - Separator valid in the upload
    SeparatorValid bool `json:"separatorValid"`


    // FileEncodingValid - File encoding valid for the upload
    FileEncodingValid bool `json:"fileEncodingValid"`


    // Errors - List of errors for the upload
    Errors []Validationerror `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Validationresult) String() string {
    
    
     o.Errors = []Validationerror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationresult) MarshalJSON() ([]byte, error) {
    type Alias Validationresult

    if ValidationresultMarshalled {
        return []byte("{}"), nil
    }
    ValidationresultMarshalled = true

    return json.Marshal(&struct {
        
        SeparatorValid bool `json:"separatorValid"`
        
        FileEncodingValid bool `json:"fileEncodingValid"`
        
        Errors []Validationerror `json:"errors"`
        *Alias
    }{

        


        


        
        Errors: []Validationerror{{}},
        

        Alias: (*Alias)(u),
    })
}

