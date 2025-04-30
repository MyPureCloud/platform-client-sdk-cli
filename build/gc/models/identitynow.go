package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentitynowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentitynowDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Identitynow
type Identitynow struct { 
    


    // Name
    Name string `json:"name"`


    // Disabled
    Disabled bool `json:"disabled"`


    // IssuerURI
    IssuerURI string `json:"issuerURI"`


    // SsoTargetURI
    SsoTargetURI string `json:"ssoTargetURI"`


    // SloURI
    SloURI string `json:"sloURI"`


    // SloBinding
    SloBinding string `json:"sloBinding"`


    // RelyingPartyIdentifier
    RelyingPartyIdentifier string `json:"relyingPartyIdentifier"`


    // Certificate
    Certificate string `json:"certificate"`


    // Certificates
    Certificates []string `json:"certificates"`


    // LogoImageData
    LogoImageData string `json:"logoImageData"`


    // NameIdentifierFormat
    NameIdentifierFormat string `json:"nameIdentifierFormat"`


    // SsoBinding
    SsoBinding string `json:"ssoBinding"`


    // SignAuthnRequests
    SignAuthnRequests bool `json:"signAuthnRequests"`


    // ProviderName
    ProviderName string `json:"providerName"`


    // DisplayOnLogin
    DisplayOnLogin bool `json:"displayOnLogin"`


    // MetadataURL
    MetadataURL string `json:"metadataURL"`


    

}

// String returns a JSON representation of the model
func (o *Identitynow) String() string {
    
    
    
    
    
    
    
    
     o.Certificates = []string{""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identitynow) MarshalJSON() ([]byte, error) {
    type Alias Identitynow

    if IdentitynowMarshalled {
        return []byte("{}"), nil
    }
    IdentitynowMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Disabled bool `json:"disabled"`
        
        IssuerURI string `json:"issuerURI"`
        
        SsoTargetURI string `json:"ssoTargetURI"`
        
        SloURI string `json:"sloURI"`
        
        SloBinding string `json:"sloBinding"`
        
        RelyingPartyIdentifier string `json:"relyingPartyIdentifier"`
        
        Certificate string `json:"certificate"`
        
        Certificates []string `json:"certificates"`
        
        LogoImageData string `json:"logoImageData"`
        
        NameIdentifierFormat string `json:"nameIdentifierFormat"`
        
        SsoBinding string `json:"ssoBinding"`
        
        SignAuthnRequests bool `json:"signAuthnRequests"`
        
        ProviderName string `json:"providerName"`
        
        DisplayOnLogin bool `json:"displayOnLogin"`
        
        MetadataURL string `json:"metadataURL"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Certificates: []string{""},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

