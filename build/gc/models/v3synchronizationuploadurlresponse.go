package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationuploadurlresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationuploadurlresponseDud struct { 
    


    

}

// V3synchronizationuploadurlresponse
type V3synchronizationuploadurlresponse struct { 
    // Url - Pre-signed URL to PUT the file to.
    Url string `json:"url"`


    // Headers - Required headers when uploading a file through PUT request to the URL.
    Headers map[string]string `json:"headers"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadurlresponse) String() string {
    
     o.Headers = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationuploadurlresponse) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationuploadurlresponse

    if V3synchronizationuploadurlresponseMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationuploadurlresponseMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Headers map[string]string `json:"headers"`
        *Alias
    }{

        


        
        Headers: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

