package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdditionalsocialmediamessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdditionalsocialmediamessageDud struct { 
    


    


    

}

// Additionalsocialmediamessage
type Additionalsocialmediamessage struct { 
    // TextBody - The body of the text message.  Maximum character count is 2000 characters.
    TextBody string `json:"textBody"`


    // MediaIds - The media ids associated with the text message. See https://developer.genesys.cloud/api/rest/v2/conversations/messaging-media-upload for example usage.
    MediaIds []string `json:"mediaIds"`


    // InReplyToMessageId - The ID of the message to which this request is replying.
    InReplyToMessageId string `json:"inReplyToMessageId"`

}

// String returns a JSON representation of the model
func (o *Additionalsocialmediamessage) String() string {
    
     o.MediaIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Additionalsocialmediamessage) MarshalJSON() ([]byte, error) {
    type Alias Additionalsocialmediamessage

    if AdditionalsocialmediamessageMarshalled {
        return []byte("{}"), nil
    }
    AdditionalsocialmediamessageMarshalled = true

    return json.Marshal(&struct {
        
        TextBody string `json:"textBody"`
        
        MediaIds []string `json:"mediaIds"`
        
        InReplyToMessageId string `json:"inReplyToMessageId"`
        *Alias
    }{

        


        
        MediaIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

