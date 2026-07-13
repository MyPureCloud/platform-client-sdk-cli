package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UploadattachmentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UploadattachmentresponseDud struct { 
    


    


    


    


    

}

// Uploadattachmentresponse
type Uploadattachmentresponse struct { 
    // AttachmentId - The attachment ID
    AttachmentId string `json:"attachmentId"`


    // Name - The name of the attachment file
    Name string `json:"name"`


    // Url - Pre-signed URL to upload the file
    Url string `json:"url"`


    // Headers - Required headers when uploading a file through PUT request to the URL
    Headers map[string]string `json:"headers"`


    // ConversationId - The conversation ID
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Uploadattachmentresponse) String() string {
    
    
    
     o.Headers = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uploadattachmentresponse) MarshalJSON() ([]byte, error) {
    type Alias Uploadattachmentresponse

    if UploadattachmentresponseMarshalled {
        return []byte("{}"), nil
    }
    UploadattachmentresponseMarshalled = true

    return json.Marshal(&struct {
        
        AttachmentId string `json:"attachmentId"`
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        
        Headers map[string]string `json:"headers"`
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        


        


        


        
        Headers: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

