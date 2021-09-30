package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailattachmentDud struct { 
    


    


    


    


    

}

// Emailattachment
type Emailattachment struct { 
    // Name
    Name string `json:"name"`


    // ContentPath
    ContentPath string `json:"contentPath"`


    // ContentType
    ContentType string `json:"contentType"`


    // AttachmentId
    AttachmentId string `json:"attachmentId"`


    // ContentLength
    ContentLength int `json:"contentLength"`

}

// String returns a JSON representation of the model
func (o *Emailattachment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailattachment) MarshalJSON() ([]byte, error) {
    type Alias Emailattachment

    if EmailattachmentMarshalled {
        return []byte("{}"), nil
    }
    EmailattachmentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        ContentPath string `json:"contentPath"`
        
        ContentType string `json:"contentType"`
        
        AttachmentId string `json:"attachmentId"`
        
        ContentLength int `json:"contentLength"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

