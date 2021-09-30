package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationDud struct { 
    Id string `json:"id"`


    


    


    


    ThirdPartyOrgName string `json:"thirdPartyOrgName"`


    


    


    


    


    


    


    


    ProductPlatform string `json:"productPlatform"`


    SelfUri string `json:"selfUri"`


    Features map[string]bool `json:"features"`

}

// Organization
type Organization struct { 
    


    // Name
    Name string `json:"name"`


    // DefaultLanguage - The default language for this organization. Example: 'en'
    DefaultLanguage string `json:"defaultLanguage"`


    // DefaultCountryCode - The default country code for this organization. Example: 'US'
    DefaultCountryCode string `json:"defaultCountryCode"`


    


    // ThirdPartyURI
    ThirdPartyURI string `json:"thirdPartyURI"`


    // Domain
    Domain string `json:"domain"`


    // Version - The current version of the organization.
    Version int `json:"version"`


    // State - The current state. Examples are active, inactive, deleted.
    State string `json:"state"`


    // DefaultSiteId
    DefaultSiteId string `json:"defaultSiteId"`


    // SupportURI - Email address where support tickets are sent to.
    SupportURI string `json:"supportURI"`


    // VoicemailEnabled
    VoicemailEnabled bool `json:"voicemailEnabled"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Organization) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organization) MarshalJSON() ([]byte, error) {
    type Alias Organization

    if OrganizationMarshalled {
        return []byte("{}"), nil
    }
    OrganizationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        DefaultLanguage string `json:"defaultLanguage"`
        
        DefaultCountryCode string `json:"defaultCountryCode"`
        
        
        
        ThirdPartyURI string `json:"thirdPartyURI"`
        
        Domain string `json:"domain"`
        
        Version int `json:"version"`
        
        State string `json:"state"`
        
        DefaultSiteId string `json:"defaultSiteId"`
        
        SupportURI string `json:"supportURI"`
        
        VoicemailEnabled bool `json:"voicemailEnabled"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

