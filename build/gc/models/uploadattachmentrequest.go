package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UploadattachmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UploadattachmentrequestDud struct { 
    


    


    


    

}

// Uploadattachmentrequest
type Uploadattachmentrequest struct { 
    // Name - Name of the attachment file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    Name string `json:"name"`


    // ContentLengthBytes - The length of the file to upload in bytes
    ContentLengthBytes int `json:"contentLengthBytes"`


    // ContentMd5 - Content MD5 of the file to upload
    ContentMd5 string `json:"contentMd5"`


    // InlineImage - Whether or not the attachment should be attached inline
    InlineImage bool `json:"inlineImage"`

}

// String returns a JSON representation of the model
func (o *Uploadattachmentrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uploadattachmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Uploadattachmentrequest

    if UploadattachmentrequestMarshalled {
        return []byte("{}"), nil
    }
    UploadattachmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ContentLengthBytes int `json:"contentLengthBytes"`
        
        ContentMd5 string `json:"contentMd5"`
        
        InlineImage bool `json:"inlineImage"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

