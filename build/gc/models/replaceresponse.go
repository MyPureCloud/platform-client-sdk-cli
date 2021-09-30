package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReplaceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReplaceresponseDud struct { 
    


    


    


    


    


    

}

// Replaceresponse
type Replaceresponse struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // ChangeNumber
    ChangeNumber int `json:"changeNumber"`


    // UploadStatus
    UploadStatus Domainentityref `json:"uploadStatus"`


    // UploadDestinationUri
    UploadDestinationUri string `json:"uploadDestinationUri"`


    // UploadMethod
    UploadMethod string `json:"uploadMethod"`

}

// String returns a JSON representation of the model
func (o *Replaceresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Replaceresponse) MarshalJSON() ([]byte, error) {
    type Alias Replaceresponse

    if ReplaceresponseMarshalled {
        return []byte("{}"), nil
    }
    ReplaceresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ChangeNumber int `json:"changeNumber"`
        
        UploadStatus Domainentityref `json:"uploadStatus"`
        
        UploadDestinationUri string `json:"uploadDestinationUri"`
        
        UploadMethod string `json:"uploadMethod"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

