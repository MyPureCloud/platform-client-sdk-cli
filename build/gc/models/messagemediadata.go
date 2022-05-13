package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediadataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    Status string `json:"status"`


    SelfUri string `json:"selfUri"`

}

// Messagemediadata
type Messagemediadata struct { 
    


    // Name
    Name string `json:"name"`


    // Url - The location of the media, useful for retrieving it
    Url string `json:"url"`


    // MediaType - The detected internet media type of the the media object.  If null then the media type should be dictated by the url.
    MediaType string `json:"mediaType"`


    // ContentLengthBytes - The optional content length of the the media object, in bytes.
    ContentLengthBytes int `json:"contentLengthBytes"`


    // UploadUrl - The URL returned to upload an attachment
    UploadUrl string `json:"uploadUrl"`


    


    

}

// String returns a JSON representation of the model
func (o *Messagemediadata) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediadata) MarshalJSON() ([]byte, error) {
    type Alias Messagemediadata

    if MessagemediadataMarshalled {
        return []byte("{}"), nil
    }
    MessagemediadataMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        
        MediaType string `json:"mediaType"`
        
        ContentLengthBytes int `json:"contentLengthBytes"`
        
        UploadUrl string `json:"uploadUrl"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

