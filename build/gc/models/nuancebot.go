package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebotDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Nuancebot - Model for a Nuance bot
type Nuancebot struct { 
    // Id - Nuance bot Id
    Id string `json:"id"`


    // Name - Nuance bot name
    Name string `json:"name"`


    // IntegrationId - The Integration Id for this bot
    IntegrationId string `json:"integrationId"`


    // NuanceOrganization - The Nuance Organization for this bot
    NuanceOrganization Nuanceorganization `json:"nuanceOrganization"`


    // Application - The Application for this bot
    Application Nuanceapplication `json:"application"`


    // NuanceEnvironment - The environment of the Nuance bot
    NuanceEnvironment Nuanceenvironment `json:"nuanceEnvironment"`


    // Geography - The Geography of the Nuance bot
    Geography Nuancegeography `json:"geography"`


    // Credentials - client ID/Secret objects for the credentials that execute this Nuance bot
    Credentials []Nuancebotcredentials `json:"credentials"`


    // Variables - List of available variables in this Nuance bot.  When querying, use the 'expand=variables' query param to populate this value
    Variables []Nuancebotvariable `json:"variables"`


    // TransferNodes - List of transferNodes in this Nuance bot.  When querying, use the 'expand=transferNodes' query param to populate this value
    TransferNodes []Nuancebottransfernode `json:"transferNodes"`


    // Locales - List of locales associated with this Nuance bot.  Generally in the ISO format such as 'en-US'
    Locales []string `json:"locales"`


    // Channels - List of channels associated with this Nuance bot.
    Channels []Nuancechannel `json:"channels"`


    

}

// String returns a JSON representation of the model
func (o *Nuancebot) String() string {
    
    
    
    
    
    
    
     o.Credentials = []Nuancebotcredentials{{}} 
     o.Variables = []Nuancebotvariable{{}} 
     o.TransferNodes = []Nuancebottransfernode{{}} 
     o.Locales = []string{""} 
     o.Channels = []Nuancechannel{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebot) MarshalJSON() ([]byte, error) {
    type Alias Nuancebot

    if NuancebotMarshalled {
        return []byte("{}"), nil
    }
    NuancebotMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        IntegrationId string `json:"integrationId"`
        
        NuanceOrganization Nuanceorganization `json:"nuanceOrganization"`
        
        Application Nuanceapplication `json:"application"`
        
        NuanceEnvironment Nuanceenvironment `json:"nuanceEnvironment"`
        
        Geography Nuancegeography `json:"geography"`
        
        Credentials []Nuancebotcredentials `json:"credentials"`
        
        Variables []Nuancebotvariable `json:"variables"`
        
        TransferNodes []Nuancebottransfernode `json:"transferNodes"`
        
        Locales []string `json:"locales"`
        
        Channels []Nuancechannel `json:"channels"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Credentials: []Nuancebotcredentials{{}},
        


        
        Variables: []Nuancebotvariable{{}},
        


        
        TransferNodes: []Nuancebottransfernode{{}},
        


        
        Locales: []string{""},
        


        
        Channels: []Nuancechannel{{}},
        


        

        Alias: (*Alias)(u),
    })
}

