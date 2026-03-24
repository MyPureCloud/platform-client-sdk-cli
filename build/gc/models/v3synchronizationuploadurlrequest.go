package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationuploadurlrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationuploadurlrequestDud struct { 
    


    


    


    


    

}

// V3synchronizationuploadurlrequest
type V3synchronizationuploadurlrequest struct { 
    // FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    FileName string `json:"fileName"`


    // ContentMd5 - Content MD5 of the file to upload encoded in base64, example: \"f8VicOenD6gaWTW3Lqy+KQ==\". Not the hexadecimal representation as \"7fc56270e7a70fa81a5935b72eacbe29\".
    ContentMd5 string `json:"contentMd5"`


    // ContentType - The content type of the file to upload
    ContentType string `json:"contentType"`


    // ContentLength - The length of the file to upload in bytes
    ContentLength int `json:"contentLength"`


    // Metadata - The metadata of the file to upload
    Metadata V3synchronizationuploadmetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadurlrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationuploadurlrequest) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationuploadurlrequest

    if V3synchronizationuploadurlrequestMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationuploadurlrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        ContentMd5 string `json:"contentMd5"`
        
        ContentType string `json:"contentType"`
        
        ContentLength int `json:"contentLength"`
        
        Metadata V3synchronizationuploadmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

