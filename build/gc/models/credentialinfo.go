package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialinfoDud struct { 
    Id string `json:"id"`


    


    CreatedDate time.Time `json:"createdDate"`


    ModifiedDate time.Time `json:"modifiedDate"`


    


    SelfUri string `json:"selfUri"`

}

// Credentialinfo
type Credentialinfo struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // VarType - Type of the credentials.
    VarType Credentialtype `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Credentialinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialinfo) MarshalJSON() ([]byte, error) {
    type Alias Credentialinfo

    if CredentialinfoMarshalled {
        return []byte("{}"), nil
    }
    CredentialinfoMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType Credentialtype `json:"type"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

