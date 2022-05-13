package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialtypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialtypeDud struct { 
    Id string `json:"id"`


    


    Properties interface{} `json:"properties"`


    DisplayOrder []string `json:"displayOrder"`


    Required []string `json:"required"`

}

// Credentialtype
type Credentialtype struct { 
    


    // Name
    Name string `json:"name"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Credentialtype) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialtype) MarshalJSON() ([]byte, error) {
    type Alias Credentialtype

    if CredentialtypeMarshalled {
        return []byte("{}"), nil
    }
    CredentialtypeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

