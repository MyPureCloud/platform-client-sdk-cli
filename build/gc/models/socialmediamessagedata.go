package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediamessagedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediamessagedataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    NormalizedMessage Conversationnormalizedmessage `json:"normalizedMessage"`


    NormalizedReceipts []Conversationnormalizedmessage `json:"normalizedReceipts"`


    


    


    SelfUri string `json:"selfUri"`

}

// Socialmediamessagedata
type Socialmediamessagedata struct { 
    


    // Name
    Name string `json:"name"`


    // ProviderMessageId - The unique identifier of the message from provider
    ProviderMessageId string `json:"providerMessageId"`


    // Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // FromAddress - The sender of the text message.
    FromAddress string `json:"fromAddress"`


    // ToAddress - The recipient of the text message.
    ToAddress string `json:"toAddress"`


    // Direction - The direction of the message.
    Direction string `json:"direction"`


    // MessengerType - Type of text messenger.
    MessengerType string `json:"messengerType"`


    // Status - The status of the message.
    Status string `json:"status"`


    


    


    // CreatedBy - User who sent this message.
    CreatedBy User `json:"createdBy"`


    // ConversationId - The id of the conversation of this message.
    ConversationId string `json:"conversationId"`


    

}

// String returns a JSON representation of the model
func (o *Socialmediamessagedata) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediamessagedata) MarshalJSON() ([]byte, error) {
    type Alias Socialmediamessagedata

    if SocialmediamessagedataMarshalled {
        return []byte("{}"), nil
    }
    SocialmediamessagedataMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ProviderMessageId string `json:"providerMessageId"`
        
        Timestamp time.Time `json:"timestamp"`
        
        FromAddress string `json:"fromAddress"`
        
        ToAddress string `json:"toAddress"`
        
        Direction string `json:"direction"`
        
        MessengerType string `json:"messengerType"`
        
        Status string `json:"status"`
        
        CreatedBy User `json:"createdBy"`
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

