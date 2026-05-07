package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentnotificationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentnotificationresponseDud struct { 
    


    


    


    

}

// Conversationcontentnotificationresponse - Inbound response to a notification, such as an Apple Invitations acceptance.
type Conversationcontentnotificationresponse struct { 
    // OriginatingMessageId - Reference to the ID of the original outbound notification message this response is for (e.g. the Apple requestIdentifier).
    OriginatingMessageId string `json:"originatingMessageId"`


    // ReferenceId - The business context reference associated with the notification (e.g. order ID, case ID). May be empty if the provider does not return it.
    ReferenceId string `json:"referenceId"`


    // NotificationStatus - The status of the notification response.
    NotificationStatus string `json:"notificationStatus"`


    // NotificationText - The localized display text of the user's response (e.g. \"Yes\").
    NotificationText string `json:"notificationText"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentnotificationresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentnotificationresponse

    if ConversationcontentnotificationresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentnotificationresponseMarshalled = true

    return json.Marshal(&struct {
        
        OriginatingMessageId string `json:"originatingMessageId"`
        
        ReferenceId string `json:"referenceId"`
        
        NotificationStatus string `json:"notificationStatus"`
        
        NotificationText string `json:"notificationText"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

