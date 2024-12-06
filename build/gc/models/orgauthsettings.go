package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrgauthsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrgauthsettingsDud struct { 
    


    


    


    


    


    

}

// Orgauthsettings
type Orgauthsettings struct { 
    // MultifactorAuthenticationRequired - Indicates whether multi-factor authentication is required.
    MultifactorAuthenticationRequired bool `json:"multifactorAuthenticationRequired"`


    // DomainAllowlistEnabled - Indicates whether the domain allowlist is enabled.
    DomainAllowlistEnabled bool `json:"domainAllowlistEnabled"`


    // DomainAllowlist - The list of domains that will be allowed to embed Genesys Cloud applications.
    DomainAllowlist []string `json:"domainAllowlist"`


    // IpAddressAllowlist - The list of IP addresses that will be allowed to authenticate with Genesys Cloud.
    IpAddressAllowlist []string `json:"ipAddressAllowlist"`


    // PasswordRequirements - The password requirements for the organization.
    PasswordRequirements Passwordrequirements `json:"passwordRequirements"`


    // InactivityTimeoutExclusions - The list of exempt apis from inactivity timeout.
    InactivityTimeoutExclusions []string `json:"inactivityTimeoutExclusions"`

}

// String returns a JSON representation of the model
func (o *Orgauthsettings) String() string {
    
    
     o.DomainAllowlist = []string{""} 
     o.IpAddressAllowlist = []string{""} 
    
     o.InactivityTimeoutExclusions = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Orgauthsettings) MarshalJSON() ([]byte, error) {
    type Alias Orgauthsettings

    if OrgauthsettingsMarshalled {
        return []byte("{}"), nil
    }
    OrgauthsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MultifactorAuthenticationRequired bool `json:"multifactorAuthenticationRequired"`
        
        DomainAllowlistEnabled bool `json:"domainAllowlistEnabled"`
        
        DomainAllowlist []string `json:"domainAllowlist"`
        
        IpAddressAllowlist []string `json:"ipAddressAllowlist"`
        
        PasswordRequirements Passwordrequirements `json:"passwordRequirements"`
        
        InactivityTimeoutExclusions []string `json:"inactivityTimeoutExclusions"`
        *Alias
    }{

        


        


        
        DomainAllowlist: []string{""},
        


        
        IpAddressAllowlist: []string{""},
        


        


        
        InactivityTimeoutExclusions: []string{""},
        

        Alias: (*Alias)(u),
    })
}

