package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomaincertificateauthorityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomaincertificateauthorityDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domaincertificateauthority - A certificate authority represents an organization that has issued a digital certificate for making secure connections with an edge device.
type Domaincertificateauthority struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Certificate - The authorities signed X509 PEM encoded certificate.
    Certificate string `json:"certificate"`


    // VarType - The certificate authorities type.  Managed certificate authorities are generated and maintained by Interactive Intelligence.  These are read-only and not modifiable by clients.  Remote authorities are customer managed.
    VarType string `json:"type"`


    // Services - The service(s) that the authority can be used to authenticate.
    Services []string `json:"services"`


    // CertificateDetails - The details of the parsed certificate(s).
    CertificateDetails []Certificatedetails `json:"certificateDetails"`


    

}

// String returns a JSON representation of the model
func (o *Domaincertificateauthority) String() string {
    
    
    
    
    
    
     o.Services = []string{""} 
     o.CertificateDetails = []Certificatedetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domaincertificateauthority) MarshalJSON() ([]byte, error) {
    type Alias Domaincertificateauthority

    if DomaincertificateauthorityMarshalled {
        return []byte("{}"), nil
    }
    DomaincertificateauthorityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Certificate string `json:"certificate"`
        
        VarType string `json:"type"`
        
        Services []string `json:"services"`
        
        CertificateDetails []Certificatedetails `json:"certificateDetails"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Services: []string{""},
        


        
        CertificateDetails: []Certificatedetails{{}},
        


        

        Alias: (*Alias)(u),
    })
}

