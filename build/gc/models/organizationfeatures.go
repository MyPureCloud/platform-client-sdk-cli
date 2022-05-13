package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationfeaturesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationfeaturesDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Organizationfeatures
type Organizationfeatures struct { 
    // RealtimeCIC
    RealtimeCIC bool `json:"realtimeCIC"`


    // Purecloud
    Purecloud bool `json:"purecloud"`


    // Hipaa
    Hipaa bool `json:"hipaa"`


    // UcEnabled
    UcEnabled bool `json:"ucEnabled"`


    // Pci
    Pci bool `json:"pci"`


    // PurecloudVoice
    PurecloudVoice bool `json:"purecloudVoice"`


    // XmppFederation
    XmppFederation bool `json:"xmppFederation"`


    // Chat
    Chat bool `json:"chat"`


    // InformalPhotos
    InformalPhotos bool `json:"informalPhotos"`


    // Directory
    Directory bool `json:"directory"`


    // ContactCenter
    ContactCenter bool `json:"contactCenter"`


    // UnifiedCommunications
    UnifiedCommunications bool `json:"unifiedCommunications"`


    // Custserv
    Custserv bool `json:"custserv"`

}

// String returns a JSON representation of the model
func (o *Organizationfeatures) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationfeatures) MarshalJSON() ([]byte, error) {
    type Alias Organizationfeatures

    if OrganizationfeaturesMarshalled {
        return []byte("{}"), nil
    }
    OrganizationfeaturesMarshalled = true

    return json.Marshal(&struct {
        
        RealtimeCIC bool `json:"realtimeCIC"`
        
        Purecloud bool `json:"purecloud"`
        
        Hipaa bool `json:"hipaa"`
        
        UcEnabled bool `json:"ucEnabled"`
        
        Pci bool `json:"pci"`
        
        PurecloudVoice bool `json:"purecloudVoice"`
        
        XmppFederation bool `json:"xmppFederation"`
        
        Chat bool `json:"chat"`
        
        InformalPhotos bool `json:"informalPhotos"`
        
        Directory bool `json:"directory"`
        
        ContactCenter bool `json:"contactCenter"`
        
        UnifiedCommunications bool `json:"unifiedCommunications"`
        
        Custserv bool `json:"custserv"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

