package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GamificationcontestprizeimageuploadurlrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GamificationcontestprizeimageuploadurlrequestDud struct { 
    


    


    


    


    


    

}

// Gamificationcontestprizeimageuploadurlrequest
type Gamificationcontestprizeimageuploadurlrequest struct { 
    // FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    FileName string `json:"fileName"`


    // ContentMd5 - Content MD5 of the file to upload
    ContentMd5 string `json:"contentMd5"`


    // SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
    SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`


    // ContentType - The content type of the file to upload.
    ContentType string `json:"contentType"`


    // ContentLength - The size of the file to upload.
    ContentLength int `json:"contentLength"`


    // ServerSideEncryption
    ServerSideEncryption string `json:"serverSideEncryption"`

}

// String returns a JSON representation of the model
func (o *Gamificationcontestprizeimageuploadurlrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gamificationcontestprizeimageuploadurlrequest) MarshalJSON() ([]byte, error) {
    type Alias Gamificationcontestprizeimageuploadurlrequest

    if GamificationcontestprizeimageuploadurlrequestMarshalled {
        return []byte("{}"), nil
    }
    GamificationcontestprizeimageuploadurlrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        ContentMd5 string `json:"contentMd5"`
        
        SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`
        
        ContentType string `json:"contentType"`
        
        ContentLength int `json:"contentLength"`
        
        ServerSideEncryption string `json:"serverSideEncryption"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

