package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingattachmentDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Recordingattachment
type Recordingattachment struct { 
    


    // MediaType - The type of attachment this instance represents.
    MediaType string `json:"mediaType"`


    // Url - URL of the attachment.
    Url string `json:"url"`


    // Mime - Attachment mime type.
    Mime string `json:"mime"`


    // Text - Text associated with attachment such as an image caption.
    Text string `json:"text"`


    // FileName - Suggested file name for attachment.
    FileName string `json:"fileName"`

}

// String returns a JSON representation of the model
func (o *Recordingattachment) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingattachment) MarshalJSON() ([]byte, error) {
    type Alias Recordingattachment

    if RecordingattachmentMarshalled {
        return []byte("{}"), nil
    }
    RecordingattachmentMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        
        Url string `json:"url"`
        
        Mime string `json:"mime"`
        
        Text string `json:"text"`
        
        FileName string `json:"fileName"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

