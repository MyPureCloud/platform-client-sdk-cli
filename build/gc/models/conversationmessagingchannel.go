package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagingchannelDud struct { 
    Id string `json:"id"`


    Platform string `json:"platform"`


    


    MessageId string `json:"messageId"`


    To Conversationmessagingtorecipient `json:"to"`


    From Conversationmessagingfromrecipient `json:"from"`


    Time time.Time `json:"time"`


    DateModified time.Time `json:"dateModified"`


    DateDeleted time.Time `json:"dateDeleted"`


    

}

// Conversationmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Conversationmessagingchannel struct { 
    


    


    // VarType - Specifies if this message is part of a private or public conversation.
    VarType string `json:"type"`


    


    


    


    


    


    


    // PublicMetadata - Information about a public message.
    PublicMetadata Conversationpublicmetadata `json:"publicMetadata"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagingchannel

    if ConversationmessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagingchannelMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        PublicMetadata Conversationpublicmetadata `json:"publicMetadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

