package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeofflineconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeofflineconfigurationDud struct { 
    


    


    


    


    


    

}

// Edgeofflineconfiguration
type Edgeofflineconfiguration struct { 
    // PairingId - The pairingId for your hardware Edge in the format: 00000-00000-00000-00000-00000.
    PairingId string `json:"pairingId"`


    // Network - Network settings for your hardware Edge.
    Network Edgeofflineconfigurationnetwork `json:"network"`


    // UseVerificationCode - Boolean to know if the verification code will be used to provision the Edge. Only used if the Edge is being provisioned.
    UseVerificationCode bool `json:"useVerificationCode"`


    // CertType - The type of Certificate Authority this Edge will use. Defaults to NotRequested if the Edge is already provisioned. PureCloud signed CA is recommended. Public CA signed by a trusted third party. China CA must be used if the Site's Location is in China.
    CertType string `json:"certType"`


    // Site - The Site that will be associated to the Edge. Required if the Edge is being provisioned.
    Site Domainentityref `json:"site"`


    // Proxy - Edge HTTP proxy configuration for the WAN port. The field can be a hostname, FQDN, IPv4 or IPv6 address. If port is not included, port 80 is assumed.
    Proxy string `json:"proxy"`

}

// String returns a JSON representation of the model
func (o *Edgeofflineconfiguration) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeofflineconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Edgeofflineconfiguration

    if EdgeofflineconfigurationMarshalled {
        return []byte("{}"), nil
    }
    EdgeofflineconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        PairingId string `json:"pairingId"`
        
        Network Edgeofflineconfigurationnetwork `json:"network"`
        
        UseVerificationCode bool `json:"useVerificationCode"`
        
        CertType string `json:"certType"`
        
        Site Domainentityref `json:"site"`
        
        Proxy string `json:"proxy"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

