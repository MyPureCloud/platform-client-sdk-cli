package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionuploadrequestDud struct { 
    


    

}

// Functionuploadrequest - Action function URL upload input.
type Functionuploadrequest struct { 
    // FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    FileName string `json:"fileName"`


    // SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 900 seconds
    SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`

}

// String returns a JSON representation of the model
func (o *Functionuploadrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Functionuploadrequest

    if FunctionuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    FunctionuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

