package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FinishwebauthnregistrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FinishwebauthnregistrationrequestDud struct { 
    


    


    

}

// Finishwebauthnregistrationrequest
type Finishwebauthnregistrationrequest struct { 
    // Credential - The credential creation response returned by the authenticator (i.e., the result of navigator.credentials.create()).
    Credential Publickeycredentialcreationresponse `json:"credential"`


    // Name - The user-facing name for this verifier.
    Name string `json:"name"`


    // VarDefault - Indicates whether this should be set as the user's default verifier.
    VarDefault bool `json:"default"`

}

// String returns a JSON representation of the model
func (o *Finishwebauthnregistrationrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Finishwebauthnregistrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Finishwebauthnregistrationrequest

    if FinishwebauthnregistrationrequestMarshalled {
        return []byte("{}"), nil
    }
    FinishwebauthnregistrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Credential Publickeycredentialcreationresponse `json:"credential"`
        
        Name string `json:"name"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

