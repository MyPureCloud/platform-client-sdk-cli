package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableimportjoberrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableimportjoberrorDud struct { 
    


    


    


    


    

}

// Decisiontableimportjoberror - Error details when a decision table import job fails
type Decisiontableimportjoberror struct { 
    // ErrorCode - The error code for this job failure.
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - A human-readable error message.
    ErrorMessage string `json:"errorMessage"`


    // MessageWithParams - Parameterized message template for the aggregate failure (when applicable)
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams - Parameters for messageWithParams
    MessageParams map[string]string `json:"messageParams"`


    // ValidationErrors - Validation failures for individual rows or the file structure
    ValidationErrors []Decisiontablejobvalidationerror `json:"validationErrors"`

}

// String returns a JSON representation of the model
func (o *Decisiontableimportjoberror) String() string {
    
    
    
     o.MessageParams = map[string]string{"": ""} 
     o.ValidationErrors = []Decisiontablejobvalidationerror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableimportjoberror) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableimportjoberror

    if DecisiontableimportjoberrorMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableimportjoberrorMarshalled = true

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

