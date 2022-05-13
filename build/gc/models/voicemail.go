package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailDud struct { 
    


    

}

// Voicemail
type Voicemail struct { 
    // Id - The voicemail id
    Id string `json:"id"`


    // UploadStatus - current state of the voicemail upload
    UploadStatus string `json:"uploadStatus"`

}

// String returns a JSON representation of the model
func (o *Voicemail) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemail) MarshalJSON() ([]byte, error) {
    type Alias Voicemail

    if VoicemailMarshalled {
        return []byte("{}"), nil
    }
    VoicemailMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        UploadStatus string `json:"uploadStatus"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

