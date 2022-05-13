package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PosttextrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PosttextrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Posttextrequest
type Posttextrequest struct { 
    // BotId - ID of the bot to send the text to.
    BotId string `json:"botId"`


    // BotAlias - Alias/Version of the bot
    BotAlias string `json:"botAlias"`


    // IntegrationId - the integration service id for the bot's credentials
    IntegrationId string `json:"integrationId"`


    // BotSessionId - GUID for this bot's session
    BotSessionId string `json:"botSessionId"`


    // PostTextMessage - Message to send to the bot
    PostTextMessage Posttextmessage `json:"postTextMessage"`


    // LanguageCode - The launguage code the bot will run under
    LanguageCode string `json:"languageCode"`


    // BotSessionTimeoutMinutes - Override timeout for the bot session. This should be greater than 10 minutes.
    BotSessionTimeoutMinutes int `json:"botSessionTimeoutMinutes"`


    // BotChannels - The channels this bot is utilizing
    BotChannels []string `json:"botChannels"`


    // BotCorrelationId - Id for tracking the activity - this will be returned in the response
    BotCorrelationId string `json:"botCorrelationId"`


    // MessagingPlatformType - If the channels list contains a 'Messaging' item and the messaging platform is known, include it here to get accurate analytics
    MessagingPlatformType string `json:"messagingPlatformType"`


    // AmazonLexRequest - Provider specific settings, if any
    AmazonLexRequest Amazonlexrequest `json:"amazonLexRequest"`


    // GoogleDialogflow - Provider specific settings, if any
    GoogleDialogflow Googledialogflowcustomsettings `json:"googleDialogflow"`


    // GenesysBotConnector - Provider specific settings, if any
    GenesysBotConnector Genesysbotconnector `json:"genesysBotConnector"`


    // NuanceMixDlg - Provider specific settings, if any
    NuanceMixDlg Nuancemixdlgsettings `json:"nuanceMixDlg"`

}

// String returns a JSON representation of the model
func (o *Posttextrequest) String() string {
    
    
    
    
    
    
    
     o.BotChannels = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Posttextrequest) MarshalJSON() ([]byte, error) {
    type Alias Posttextrequest

    if PosttextrequestMarshalled {
        return []byte("{}"), nil
    }
    PosttextrequestMarshalled = true

    return json.Marshal(&struct {
        
        BotId string `json:"botId"`
        
        BotAlias string `json:"botAlias"`
        
        IntegrationId string `json:"integrationId"`
        
        BotSessionId string `json:"botSessionId"`
        
        PostTextMessage Posttextmessage `json:"postTextMessage"`
        
        LanguageCode string `json:"languageCode"`
        
        BotSessionTimeoutMinutes int `json:"botSessionTimeoutMinutes"`
        
        BotChannels []string `json:"botChannels"`
        
        BotCorrelationId string `json:"botCorrelationId"`
        
        MessagingPlatformType string `json:"messagingPlatformType"`
        
        AmazonLexRequest Amazonlexrequest `json:"amazonLexRequest"`
        
        GoogleDialogflow Googledialogflowcustomsettings `json:"googleDialogflow"`
        
        GenesysBotConnector Genesysbotconnector `json:"genesysBotConnector"`
        
        NuanceMixDlg Nuancemixdlgsettings `json:"nuanceMixDlg"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        BotChannels: []string{""},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

