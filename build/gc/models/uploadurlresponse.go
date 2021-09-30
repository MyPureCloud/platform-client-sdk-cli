package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UploadurlresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UploadurlresponseDud struct { 
    Url string `json:"url"`


    UploadKey string `json:"uploadKey"`


    Headers map[string]string `json:"headers"`

}

// Uploadurlresponse
type Uploadurlresponse struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Uploadurlresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Uploadurlresponse) MarshalJSON() ([]byte, error) {
    type Alias Uploadurlresponse

    if UploadurlresponseMarshalled {
        return []byte("{}"), nil
    }
    UploadurlresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

