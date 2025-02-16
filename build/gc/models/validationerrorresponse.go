package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationerrorresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationerrorresponseDud struct { 
    


    


    

}

// Validationerrorresponse
type Validationerrorresponse struct { 
    // Message - Message string of the validation error.
    Message string `json:"message"`


    // ErrorType - Type of validation errror.
    ErrorType string `json:"errorType"`


    // Arguments - Map of argument names to corresponding values. Used for localization.
    Arguments map[string]string `json:"arguments"`

}

// String returns a JSON representation of the model
func (o *Validationerrorresponse) String() string {
    
    
     o.Arguments = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationerrorresponse) MarshalJSON() ([]byte, error) {
    type Alias Validationerrorresponse

    if ValidationerrorresponseMarshalled {
        return []byte("{}"), nil
    }
    ValidationerrorresponseMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        ErrorType string `json:"errorType"`
        
        Arguments map[string]string `json:"arguments"`
        *Alias
    }{

        


        


        
        Arguments: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

