package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TemplateparameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TemplateparameterDud struct { 
    


    

}

// Templateparameter
type Templateparameter struct { 
    // Id - Response substitution identifier
    Id string `json:"id"`


    // Value - Response substitution value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Templateparameter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Templateparameter) MarshalJSON() ([]byte, error) {
    type Alias Templateparameter

    if TemplateparameterMarshalled {
        return []byte("{}"), nil
    }
    TemplateparameterMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

