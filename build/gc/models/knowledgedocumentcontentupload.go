package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentcontentuploadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentcontentuploadDud struct { 
    Id string `json:"id"`


    


    


    Status string `json:"status"`


    UploadKey string `json:"uploadKey"`


    Url string `json:"url"`


    Headers map[string]string `json:"headers"`


    Document Addressableentityref `json:"document"`


    ErrorMessage string `json:"errorMessage"`


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentcontentupload
type Knowledgedocumentcontentupload struct { 
    


    // ContentType - Type of Article Content.
    ContentType string `json:"contentType"`


    // FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    FileName string `json:"fileName"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentcontentupload) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentcontentupload) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentcontentupload

    if KnowledgedocumentcontentuploadMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentcontentuploadMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        FileName string `json:"fileName"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

