package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IncomingmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IncomingmessagerequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Incomingmessagerequest - Incoming Message request
type Incomingmessagerequest struct { 
    // BotId - The unique id of the bot.
    BotId string `json:"botId"`


    // BotName - Name of the bot
    BotName string `json:"botName"`


    // BotVersion - The version of the bot.
    BotVersion string `json:"botVersion"`


    // IntegrationId - The Integration Id for the bot's configuration
    IntegrationId string `json:"integrationId"`


    // BotSessionId - The id of the session. This id will be used for an entire conversation with the bot (a series of back and forth between the bot until the bot has fulfilled its intent).
    BotSessionId string `json:"botSessionId"`


    // AutomateFlowExecId - Execution ID of the Automate Flow.
    AutomateFlowExecId string `json:"automateFlowExecId"`


    // ConversationId - Genesys conversation ID.
    ConversationId string `json:"conversationId"`


    // LanguageCode - Language code for the conversation (e.g., 'en-US').
    LanguageCode string `json:"languageCode"`


    // InputMessage - Message received from the user to send to the bot
    InputMessage Inputmessage `json:"inputMessage"`


    // MessagingPlatformType - Type of messaging platform being used
    MessagingPlatformType string `json:"messagingPlatformType"`


    // Channels - The channels this bot is utilizing.
    Channels []string `json:"channels"`


    // BotSessionTimeout - Timeout duration for bot session in minutes.
    BotSessionTimeout int `json:"botSessionTimeout"`


    // Parameters - This is a map of string-string key, value pairs containing optional fields that can be passed to the bot for custom behavior, tracking, etc.
    Parameters map[string]string `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Incomingmessagerequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Channels = []string{""} 
    
     o.Parameters = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Incomingmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Incomingmessagerequest

    if IncomingmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    IncomingmessagerequestMarshalled = true

    return json.Marshal(&struct {
        
        BotId string `json:"botId"`
        
        BotName string `json:"botName"`
        
        BotVersion string `json:"botVersion"`
        
        IntegrationId string `json:"integrationId"`
        
        BotSessionId string `json:"botSessionId"`
        
        AutomateFlowExecId string `json:"automateFlowExecId"`
        
        ConversationId string `json:"conversationId"`
        
        LanguageCode string `json:"languageCode"`
        
        InputMessage Inputmessage `json:"inputMessage"`
        
        MessagingPlatformType string `json:"messagingPlatformType"`
        
        Channels []string `json:"channels"`
        
        BotSessionTimeout int `json:"botSessionTimeout"`
        
        Parameters map[string]string `json:"parameters"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        Channels: []string{""},
        


        


        
        Parameters: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

