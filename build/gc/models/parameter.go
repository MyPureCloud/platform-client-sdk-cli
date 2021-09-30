package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParameterDud struct { 
    


    


    


    

}

// Parameter
type Parameter struct { 
    // Name
    Name string `json:"name"`


    // ParameterType
    ParameterType string `json:"parameterType"`


    // Domain
    Domain string `json:"domain"`


    // Required
    Required bool `json:"required"`

}

// String returns a JSON representation of the model
func (o *Parameter) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Parameter) MarshalJSON() ([]byte, error) {
    type Alias Parameter

    if ParameterMarshalled {
        return []byte("{}"), nil
    }
    ParameterMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        ParameterType string `json:"parameterType"`
        
        Domain string `json:"domain"`
        
        Required bool `json:"required"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

