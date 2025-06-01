package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutgoingmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutgoingmessagerequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Outgoingmessagerequest - Outgoing Message request
type Outgoingmessagerequest struct { 
    // BotId - The unique id of the bot.
    BotId string `json:"botId"`


    // BotVersion - The version of the bot.
    BotVersion string `json:"botVersion"`


    // BotSessionId - The id of the session. This id will be used for an entire conversation with the bot (a series of back and forth between the bot until the bot has fulfilled its intent).
    BotSessionId string `json:"botSessionId"`


    // BotState - The state of the bot reported
    BotState string `json:"botState"`


    // LanguageCode - The language used for this message. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code).
    LanguageCode string `json:"languageCode"`


    // ReplyMessages - This is a list of messages to send back to the user, this field can be null or an empty list.
    ReplyMessages []Replymessage `json:"replyMessages"`


    // Intent - The name of the intent the bot is either processing or has processed, this will be blank if no intent could be detected.
    Intent string `json:"intent"`


    // Confidence - A value between 0 and 1.0 denoting the confidence of the discovered intent (if found) this is optional and if left null genesys assumes a confidence of 1.0 on success and 0 on fail.
    Confidence float64 `json:"confidence"`


    // ErrorInfo - If the botState is Failed the bot can add this error object with more details about the error.
    ErrorInfo Errorinfo `json:"errorInfo"`


    // Parameters - This is a map of string-string key, value pairs containing optional fields that can be passed from the bot for custom behavior, tracking, etc, which can be used by the flow.
    Parameters map[string]string `json:"parameters"`


    // Entities - A set of entity values that go along with the intent.
    Entities []Botentityvalue `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Outgoingmessagerequest) String() string {
    
    
    
    
    
     o.ReplyMessages = []Replymessage{{}} 
    
    
    
     o.Parameters = map[string]string{"": ""} 
     o.Entities = []Botentityvalue{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outgoingmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Outgoingmessagerequest

    if OutgoingmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    OutgoingmessagerequestMarshalled = true

    return json.Marshal(&struct {
        
        BotId string `json:"botId"`
        
        BotVersion string `json:"botVersion"`
        
        BotSessionId string `json:"botSessionId"`
        
        BotState string `json:"botState"`
        
        LanguageCode string `json:"languageCode"`
        
        ReplyMessages []Replymessage `json:"replyMessages"`
        
        Intent string `json:"intent"`
        
        Confidence float64 `json:"confidence"`
        
        ErrorInfo Errorinfo `json:"errorInfo"`
        
        Parameters map[string]string `json:"parameters"`
        
        Entities []Botentityvalue `json:"entities"`
        *Alias
    }{

        


        


        


        


        


        
        ReplyMessages: []Replymessage{{}},
        


        


        


        


        
        Parameters: map[string]string{"": ""},
        


        
        Entities: []Botentityvalue{{}},
        

        Alias: (*Alias)(u),
    })
}

