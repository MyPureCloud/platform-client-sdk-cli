package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnifiedcommunicationsintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnifiedcommunicationsintegrationDud struct { 
    Id string `json:"id"`


    


    UcIntegrationKey Addressableentityref `json:"ucIntegrationKey"`


    IntegrationPresenceSource string `json:"integrationPresenceSource"`


    PbxPermission string `json:"pbxPermission"`


    Icon Ucicon `json:"icon"`


    BadgeIcons map[string]Ucicon `json:"badgeIcons"`


    I10n map[string]Uci10n `json:"i10n"`


    PolledPresence bool `json:"polledPresence"`


    PollIntervalSec int `json:"pollIntervalSec"`


    IncludeBadge bool `json:"includeBadge"`


    UserPermissions []string `json:"userPermissions"`


    


    SelfUri string `json:"selfUri"`

}

// Unifiedcommunicationsintegration - UC Integration UI configuration data
type Unifiedcommunicationsintegration struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    // OauthScopes
    OauthScopes []string `json:"oauthScopes"`


    

}

// String returns a JSON representation of the model
func (o *Unifiedcommunicationsintegration) String() string {
    
     o.OauthScopes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unifiedcommunicationsintegration) MarshalJSON() ([]byte, error) {
    type Alias Unifiedcommunicationsintegration

    if UnifiedcommunicationsintegrationMarshalled {
        return []byte("{}"), nil
    }
    UnifiedcommunicationsintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        OauthScopes []string `json:"oauthScopes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        OauthScopes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

