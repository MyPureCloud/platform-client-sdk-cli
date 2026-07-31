package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PublickeycredentialcreationoptionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PublickeycredentialcreationoptionsDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Publickeycredentialcreationoptions
type Publickeycredentialcreationoptions struct { 
    // Challenge - Cryptographic challenge from the relying party (base64url-encoded). Must be returned to the relying party in the authenticator's response.
    Challenge string `json:"challenge"`


    // Rp - Information about the relying party.
    Rp Relyingpartyentity `json:"rp"`


    // User - Information about the user being registered.
    User Userentity `json:"user"`


    // PubKeyCredParams - Public key credential parameters acceptable to the relying party, in order of preference.
    PubKeyCredParams []Credentialparameter `json:"pubKeyCredParams"`


    // Timeout - Time in milliseconds the relying party is willing to wait for the registration operation to complete.
    Timeout int `json:"timeout"`


    // ExcludeCredentials - Credentials that should be excluded from registration (e.g., to prevent re-registering an existing authenticator).
    ExcludeCredentials []Credentialdescriptor `json:"excludeCredentials"`


    // AuthenticatorSelection - Constraints on the type of authenticator that can be used.
    AuthenticatorSelection Authenticatorselection `json:"authenticatorSelection"`


    // Hints - Hints about the type of authenticator the user should use (e.g., 'security-key', 'client-device', 'hybrid').
    Hints []string `json:"hints"`


    // Attestation - The relying party's attestation conveyance preference ('none', 'indirect', 'direct', or 'enterprise').
    Attestation string `json:"attestation"`


    // AttestationFormats - Acceptable attestation statement formats, in order of preference.
    AttestationFormats []string `json:"attestationFormats"`


    // Extensions - Inputs to client-side WebAuthn extensions.
    Extensions map[string]interface{} `json:"extensions"`

}

// String returns a JSON representation of the model
func (o *Publickeycredentialcreationoptions) String() string {
    
    
    
     o.PubKeyCredParams = []Credentialparameter{{}} 
    
     o.ExcludeCredentials = []Credentialdescriptor{{}} 
    
     o.Hints = []string{""} 
    
     o.AttestationFormats = []string{""} 
     o.Extensions = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Publickeycredentialcreationoptions) MarshalJSON() ([]byte, error) {
    type Alias Publickeycredentialcreationoptions

    if PublickeycredentialcreationoptionsMarshalled {
        return []byte("{}"), nil
    }
    PublickeycredentialcreationoptionsMarshalled = true

    return json.Marshal(&struct {
        
        Challenge string `json:"challenge"`
        
        Rp Relyingpartyentity `json:"rp"`
        
        User Userentity `json:"user"`
        
        PubKeyCredParams []Credentialparameter `json:"pubKeyCredParams"`
        
        Timeout int `json:"timeout"`
        
        ExcludeCredentials []Credentialdescriptor `json:"excludeCredentials"`
        
        AuthenticatorSelection Authenticatorselection `json:"authenticatorSelection"`
        
        Hints []string `json:"hints"`
        
        Attestation string `json:"attestation"`
        
        AttestationFormats []string `json:"attestationFormats"`
        
        Extensions map[string]interface{} `json:"extensions"`
        *Alias
    }{

        


        


        


        
        PubKeyCredParams: []Credentialparameter{{}},
        


        


        
        ExcludeCredentials: []Credentialdescriptor{{}},
        


        


        
        Hints: []string{""},
        


        


        
        AttestationFormats: []string{""},
        


        
        Extensions: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

