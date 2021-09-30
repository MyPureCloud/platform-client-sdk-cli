package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttachmentDud struct { 
    


    


    


    


    


    

}

// Attachment
type Attachment struct { 
    // AttachmentId - The unique identifier for the attachment.
    AttachmentId string `json:"attachmentId"`


    // Name - The name of the attachment.
    Name string `json:"name"`


    // ContentUri - The content uri of the attachment. If set, this is commonly a public api download location.
    ContentUri string `json:"contentUri"`


    // ContentType - The type of file the attachment is.
    ContentType string `json:"contentType"`


    // ContentLength - The length of the attachment file.
    ContentLength int `json:"contentLength"`


    // InlineImage - Whether or not the attachment was attached inline.,
    InlineImage bool `json:"inlineImage"`

}

// String returns a JSON representation of the model
func (o *Attachment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attachment) MarshalJSON() ([]byte, error) {
    type Alias Attachment

    if AttachmentMarshalled {
        return []byte("{}"), nil
    }
    AttachmentMarshalled = true

    return json.Marshal(&struct { 
        AttachmentId string `json:"attachmentId"`
        
        Name string `json:"name"`
        
        ContentUri string `json:"contentUri"`
        
        ContentType string `json:"contentType"`
        
        ContentLength int `json:"contentLength"`
        
        InlineImage bool `json:"inlineImage"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

