package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpencontentthumbnailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpencontentthumbnailDud struct { 
    


    


    


    

}

// Opencontentthumbnail - Thumbnail image metadata for the attachment content.
type Opencontentthumbnail struct { 
    // Url - URL of the thumbnail image.
    Url string `json:"url"`


    // Mime - Thumbnail mime type (e.g. image/jpeg).
    Mime string `json:"mime"`


    // Sha256 - Secure hash of the thumbnail content.
    Sha256 string `json:"sha256"`


    // ContentSizeBytes - Size in bytes of the thumbnail content.
    ContentSizeBytes int `json:"contentSizeBytes"`

}

// String returns a JSON representation of the model
func (o *Opencontentthumbnail) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opencontentthumbnail) MarshalJSON() ([]byte, error) {
    type Alias Opencontentthumbnail

    if OpencontentthumbnailMarshalled {
        return []byte("{}"), nil
    }
    OpencontentthumbnailMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Mime string `json:"mime"`
        
        Sha256 string `json:"sha256"`
        
        ContentSizeBytes int `json:"contentSizeBytes"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

