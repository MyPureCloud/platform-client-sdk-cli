package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UploadmediarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UploadmediarequestDud struct { 
    


    


    

}

// Uploadmediarequest
type Uploadmediarequest struct { 
    // FileName - Name of the media file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    FileName string `json:"fileName"`


    // ContentLengthBytes - The length of the file to upload in bytes
    ContentLengthBytes int `json:"contentLengthBytes"`


    // ContentMd5 - Content MD5 of the file to upload
    ContentMd5 string `json:"contentMd5"`

}

// String returns a JSON representation of the model
func (o *Uploadmediarequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uploadmediarequest) MarshalJSON() ([]byte, error) {
    type Alias Uploadmediarequest

    if UploadmediarequestMarshalled {
        return []byte("{}"), nil
    }
    UploadmediarequestMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        ContentLengthBytes int `json:"contentLengthBytes"`
        
        ContentMd5 string `json:"contentMd5"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

