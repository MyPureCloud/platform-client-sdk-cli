package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsesubstitutionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsesubstitutionDud struct { 
    


    


    

}

// Responsesubstitution - Contains information about the substitutions associated with a response.
type Responsesubstitution struct { 
    // Id - Response substitution identifier.
    Id string `json:"id"`


    // Description - Response substitution description.
    Description string `json:"description"`


    // DefaultValue - Response substitution default value.
    DefaultValue string `json:"defaultValue"`

}

// String returns a JSON representation of the model
func (o *Responsesubstitution) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsesubstitution) MarshalJSON() ([]byte, error) {
    type Alias Responsesubstitution

    if ResponsesubstitutionMarshalled {
        return []byte("{}"), nil
    }
    ResponsesubstitutionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Description string `json:"description"`
        
        DefaultValue string `json:"defaultValue"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

