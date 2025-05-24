package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentpushMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentpushDud struct { 
    


    


    


    


    


    


    

}

// Conversationcontentpush - A Push object
type Conversationcontentpush struct { 
    // DeviceType - The device type used to send the push notification
    DeviceType string `json:"deviceType"`


    // DeviceTokenId - Unique Id of the device token
    DeviceTokenId string `json:"deviceTokenId"`


    // DeviceToken - device token from the notification provider
    DeviceToken string `json:"deviceToken"`


    // FailedMessages - MessageIds failed to be sent which trigger the push event
    FailedMessages []Conversationpushfailedmessagereferences `json:"failedMessages"`


    // NotificationMessage - Title and body localized according to deployment
    NotificationMessage Conversationpushnotificationmessagelabel `json:"notificationMessage"`


    // PushProviderIntegration - Push provider integrations details configured on the deployment
    PushProviderIntegration Conversationpushproviderintegration `json:"pushProviderIntegration"`


    // Expiration - The time to live of the pushed message
    Expiration int `json:"expiration"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentpush) String() string {
    
    
    
     o.FailedMessages = []Conversationpushfailedmessagereferences{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentpush) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentpush

    if ConversationcontentpushMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentpushMarshalled = true

    return json.Marshal(&struct {
        
        DeviceType string `json:"deviceType"`
        
        DeviceTokenId string `json:"deviceTokenId"`
        
        DeviceToken string `json:"deviceToken"`
        
        FailedMessages []Conversationpushfailedmessagereferences `json:"failedMessages"`
        
        NotificationMessage Conversationpushnotificationmessagelabel `json:"notificationMessage"`
        
        PushProviderIntegration Conversationpushproviderintegration `json:"pushProviderIntegration"`
        
        Expiration int `json:"expiration"`
        *Alias
    }{

        


        


        


        
        FailedMessages: []Conversationpushfailedmessagereferences{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

