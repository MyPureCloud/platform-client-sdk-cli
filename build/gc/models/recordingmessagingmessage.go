package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingmessagingmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingmessagingmessageDud struct { 
    


    


    


    


    


    


    


    


    

}

// Recordingmessagingmessage
type Recordingmessagingmessage struct { 
    // From
    From string `json:"from"`


    // FromUser
    FromUser User `json:"fromUser"`


    // FromExternalContact
    FromExternalContact Externalcontact `json:"fromExternalContact"`


    // To
    To string `json:"to"`


    // Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Id
    Id string `json:"id"`


    // MessageText
    MessageText string `json:"messageText"`


    // MessageMediaAttachments
    MessageMediaAttachments []Messagemediaattachment `json:"messageMediaAttachments"`


    // MessageStickerAttachments
    MessageStickerAttachments []Messagestickerattachment `json:"messageStickerAttachments"`

}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MessageMediaAttachments = []Messagemediaattachment{{}} 
    
    
    
     o.MessageStickerAttachments = []Messagestickerattachment{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingmessagingmessage) MarshalJSON() ([]byte, error) {
    type Alias Recordingmessagingmessage

    if RecordingmessagingmessageMarshalled {
        return []byte("{}"), nil
    }
    RecordingmessagingmessageMarshalled = true

    return json.Marshal(&struct { 
        From string `json:"from"`
        
        FromUser User `json:"fromUser"`
        
        FromExternalContact Externalcontact `json:"fromExternalContact"`
        
        To string `json:"to"`
        
        Timestamp time.Time `json:"timestamp"`
        
        Id string `json:"id"`
        
        MessageText string `json:"messageText"`
        
        MessageMediaAttachments []Messagemediaattachment `json:"messageMediaAttachments"`
        
        MessageStickerAttachments []Messagestickerattachment `json:"messageStickerAttachments"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        MessageMediaAttachments: []Messagemediaattachment{{}},
        

        

        
        MessageStickerAttachments: []Messagestickerattachment{{}},
        

        
        Alias: (*Alias)(u),
    })
}

