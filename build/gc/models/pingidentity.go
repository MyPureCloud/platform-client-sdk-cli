package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PingidentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PingidentityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Pingidentity
type Pingidentity struct { 
    


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


    

}

// String returns a JSON representation of the model
func (o *Pingidentity) String() string {
    
    
    
    
    
    
    
    
     o.Certificates = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pingidentity) MarshalJSON() ([]byte, error) {
    type Alias Pingidentity

    if PingidentityMarshalled {
        return []byte("{}"), nil
    }
    PingidentityMarshalled = true

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
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Certificates: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

