package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagedataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    NormalizedMessage Conversationnormalizedmessage `json:"normalizedMessage"`


    NormalizedReceipts []Conversationnormalizedmessage `json:"normalizedReceipts"`


    


    


    SelfUri string `json:"selfUri"`

}

// Messagedata
type Messagedata struct { 
    


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


    // TextBody - The body of the text message.
    TextBody string `json:"textBody"`


    // Status - The status of the message.
    Status string `json:"status"`


    // Media - The media details associated to a message.
    Media []Messagemedia `json:"media"`


    // Stickers - The sticker details associated to a message.
    Stickers []Messagesticker `json:"stickers"`


    


    


    // CreatedBy - User who sent this message.
    CreatedBy User `json:"createdBy"`


    // ConversationId - The id of the conversation of this message.
    ConversationId string `json:"conversationId"`


    

}

// String returns a JSON representation of the model
func (o *Messagedata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Media = []Messagemedia{{}} 
    
    
    
     o.Stickers = []Messagesticker{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagedata) MarshalJSON() ([]byte, error) {
    type Alias Messagedata

    if MessagedataMarshalled {
        return []byte("{}"), nil
    }
    MessagedataMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ProviderMessageId string `json:"providerMessageId"`
        
        Timestamp time.Time `json:"timestamp"`
        
        FromAddress string `json:"fromAddress"`
        
        ToAddress string `json:"toAddress"`
        
        Direction string `json:"direction"`
        
        MessengerType string `json:"messengerType"`
        
        TextBody string `json:"textBody"`
        
        Status string `json:"status"`
        
        Media []Messagemedia `json:"media"`
        
        Stickers []Messagesticker `json:"stickers"`
        
        
        
        
        
        CreatedBy User `json:"createdBy"`
        
        ConversationId string `json:"conversationId"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Media: []Messagemedia{{}},
        

        

        
        Stickers: []Messagesticker{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

