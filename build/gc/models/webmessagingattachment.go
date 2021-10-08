package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingattachmentDud struct { 
    Id string `json:"id"`


    MediaType string `json:"mediaType"`


    Url string `json:"url"`


    Mime string `json:"mime"`


    Text string `json:"text"`


    Sha256 string `json:"sha256"`


    Filename string `json:"filename"`


    FileSize int `json:"fileSize"`

}

// Webmessagingattachment - Attachment object.
type Webmessagingattachment struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Webmessagingattachment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingattachment) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingattachment

    if WebmessagingattachmentMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingattachmentMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

