package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentattachmentDud struct { 
    


    


    


    


    


    


    

}

// Conversationcontentattachment - Attachment object.
type Conversationcontentattachment struct { 
    // Id - Provider specific ID for attachment. For example, a LINE sticker ID.
    Id string `json:"id"`


    // MediaType - The type of attachment this instance represents.
    MediaType string `json:"mediaType"`


    // Url - URL of the attachment.
    Url string `json:"url"`


    // Mime - Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
    Mime string `json:"mime"`


    // Text - Text associated with attachment such as an image caption.
    Text string `json:"text"`


    // Sha256 - Secure hash of the attachment content.
    Sha256 string `json:"sha256"`


    // Filename - Suggested file name for attachment.
    Filename string `json:"filename"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentattachment) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentattachment) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentattachment

    if ConversationcontentattachmentMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentattachmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        MediaType string `json:"mediaType"`
        
        Url string `json:"url"`
        
        Mime string `json:"mime"`
        
        Text string `json:"text"`
        
        Sha256 string `json:"sha256"`
        
        Filename string `json:"filename"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

