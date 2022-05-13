package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediaDud struct { 
    


    


    


    


    

}

// Messagemedia
type Messagemedia struct { 
    // Url - The location of the media, useful for retrieving it
    Url string `json:"url"`


    // MediaType - The optional internet media type of the the media object.  If null then the media type should be dictated by the url
    MediaType string `json:"mediaType"`


    // ContentLengthBytes - The optional content length of the the media object, in bytes.
    ContentLengthBytes int `json:"contentLengthBytes"`


    // Name - The optional name of the the media object.
    Name string `json:"name"`


    // Id - The optional id of the the media object.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Messagemedia) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemedia) MarshalJSON() ([]byte, error) {
    type Alias Messagemedia

    if MessagemediaMarshalled {
        return []byte("{}"), nil
    }
    MessagemediaMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        MediaType string `json:"mediaType"`
        
        ContentLengthBytes int `json:"contentLengthBytes"`
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

