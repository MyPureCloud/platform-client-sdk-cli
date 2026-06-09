package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexportjoberrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexportjoberrorDud struct { 
    


    


    


    


    

}

// Decisiontableexportjoberror - Error details when a decision table export job fails
type Decisiontableexportjoberror struct { 
    // ErrorCode - The error code for this job failure.
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - A human-readable error message.
    ErrorMessage string `json:"errorMessage"`


    // MessageWithParams - Parameterized message template for the aggregate failure (when applicable)
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams - Parameters for messageWithParams
    MessageParams map[string]string `json:"messageParams"`


    // ValidationErrors - Validation failures for the export job
    ValidationErrors []Decisiontablejobvalidationerror `json:"validationErrors"`

}

// String returns a JSON representation of the model
func (o *Decisiontableexportjoberror) String() string {
    
    
    
     o.MessageParams = map[string]string{"": ""} 
     o.ValidationErrors = []Decisiontablejobvalidationerror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexportjoberror) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexportjoberror

    if DecisiontableexportjoberrorMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexportjoberrorMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]string `json:"messageParams"`
        
        ValidationErrors []Decisiontablejobvalidationerror `json:"validationErrors"`
        *Alias
    }{

        


        


        


        
        MessageParams: map[string]string{"": ""},
        


        
        ValidationErrors: []Decisiontablejobvalidationerror{{}},
        

        Alias: (*Alias)(u),
    })
}

