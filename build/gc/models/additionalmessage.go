package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdditionalmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdditionalmessageDud struct { 
    


    


    


    

}

// Additionalmessage
type Additionalmessage struct { 
    // TextBody - The body of the text message.  Maximum character counts are: SMS - 765 characters, other channels - 2000 characters.
    TextBody string `json:"textBody"`


    // MediaIds - The media ids associated with the text message. See https://developer.genesys.cloud/api/rest/v2/conversations/messaging-media-upload for example usage.
    MediaIds []string `json:"mediaIds"`


    // StickerIds - The sticker ids associated with the text message.
    StickerIds []string `json:"stickerIds"`


    // MessagingTemplate - The messaging template use to send a predefined canned response with the message
    MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`

}

// String returns a JSON representation of the model
func (o *Additionalmessage) String() string {
    
     o.MediaIds = []string{""} 
     o.StickerIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Additionalmessage) MarshalJSON() ([]byte, error) {
    type Alias Additionalmessage

    if AdditionalmessageMarshalled {
        return []byte("{}"), nil
    }
    AdditionalmessageMarshalled = true

    return json.Marshal(&struct {
        
        TextBody string `json:"textBody"`
        
        MediaIds []string `json:"mediaIds"`
        
        StickerIds []string `json:"stickerIds"`
        
        MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`
        *Alias
    }{

        


        
        MediaIds: []string{""},
        


        
        StickerIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

