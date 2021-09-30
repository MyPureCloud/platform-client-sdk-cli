package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediaattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediaattachmentDud struct { 
    


    


    


    


    

}

// Messagemediaattachment
type Messagemediaattachment struct { 
    // Url - The location of the media, useful for retrieving it
    Url string `json:"url"`


    // MediaType - The optional internet media type of the the media object.If null then the media type should be dictated by the url.
    MediaType string `json:"mediaType"`


    // ContentLength - The optional content length of the the media object, in bytes.
    ContentLength int `json:"contentLength"`


    // Name
    Name string `json:"name"`


    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Messagemediaattachment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediaattachment) MarshalJSON() ([]byte, error) {
    type Alias Messagemediaattachment

    if MessagemediaattachmentMarshalled {
        return []byte("{}"), nil
    }
    MessagemediaattachmentMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        MediaType string `json:"mediaType"`
        
        ContentLength int `json:"contentLength"`
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

