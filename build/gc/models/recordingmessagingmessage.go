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
    // From - The message sender session id.
    From string `json:"from"`


    // FromUser - The user who sent this message.
    FromUser User `json:"fromUser"`


    // FromExternalContact - The PureCloud external contact sender details.
    FromExternalContact Externalcontact `json:"fromExternalContact"`


    // To - The message recipient.
    To string `json:"to"`


    // Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Purpose - A well known string that specifies the purpose or type of the participant on this communication.
    Purpose string `json:"purpose"`


    // ParticipantId - A globally unique identifier for the participant on this communication.
    ParticipantId string `json:"participantId"`


    // Queue - A globally unique identifier for the queue involved in this communication.
    Queue Addressableentityref `json:"queue"`


    // Workflow - A globally unique identifier for the workflow involved in this communication.
    Workflow Addressableentityref `json:"workflow"`


    // MessageText - The content of this message.
    MessageText string `json:"messageText"`


    // MessageMediaAttachments - List of media objects attached  with this message.
    MessageMediaAttachments []Messagemediaattachment `json:"messageMediaAttachments"`


    // MessageStickerAttachments - List of message stickers attached with this message.
    MessageStickerAttachments []Messagestickerattachment `json:"messageStickerAttachments"`


    // QuickReplies - List of quick reply options offered with this message.
    QuickReplies []Quickreply `json:"quickReplies"`


    // ButtonResponse - Button Response selected by user for this message.
    ButtonResponse Buttonresponse `json:"buttonResponse"`


    // Story - Ephemeral story content.
    Story Recordingcontentstory `json:"story"`


    // Cards - List of cards offered for this message
    Cards []Card `json:"cards"`


    // NotificationTemplate - Template notification content.
    NotificationTemplate Recordingnotificationtemplate `json:"notificationTemplate"`


    // DatePicker - DatePicker content object.
    DatePicker Datepicker `json:"datePicker"`


    // ContentType - Indicates the content type for this message
    ContentType string `json:"contentType"`


    // SocialVisibility - For social media messages, the visibility of the message in the originating social platform
    SocialVisibility string `json:"socialVisibility"`


    // Events - List of event elements
    Events []Conversationmessageevent `json:"events"`

}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.MessageMediaAttachments = []Messagemediaattachment{{}} 
     o.MessageStickerAttachments = []Messagestickerattachment{{}} 
     o.QuickReplies = []Quickreply{{}} 
    
    
     o.Cards = []Card{{}} 
    
    
    
    
     o.Events = []Conversationmessageevent{{}} 

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
        
        Purpose string `json:"purpose"`
        
        ParticipantId string `json:"participantId"`
        
        Queue Addressableentityref `json:"queue"`
        
        Workflow Addressableentityref `json:"workflow"`
        
        MessageText string `json:"messageText"`
        
        MessageMediaAttachments []Messagemediaattachment `json:"messageMediaAttachments"`
        
        MessageStickerAttachments []Messagestickerattachment `json:"messageStickerAttachments"`
        
        QuickReplies []Quickreply `json:"quickReplies"`
        
        ButtonResponse Buttonresponse `json:"buttonResponse"`
        
        Story Recordingcontentstory `json:"story"`
        
        Cards []Card `json:"cards"`
        
        NotificationTemplate Recordingnotificationtemplate `json:"notificationTemplate"`
        
        DatePicker Datepicker `json:"datePicker"`
        
        ContentType string `json:"contentType"`
        
        SocialVisibility string `json:"socialVisibility"`
        
        Events []Conversationmessageevent `json:"events"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        MessageMediaAttachments: []Messagemediaattachment{{}},
        


        
        MessageStickerAttachments: []Messagestickerattachment{{}},
        


        
        QuickReplies: []Quickreply{{}},
        


        


        


        
        Cards: []Card{{}},
        


        


        


        


        


        
        Events: []Conversationmessageevent{{}},
        

        Alias: (*Alias)(u),
    })
}

