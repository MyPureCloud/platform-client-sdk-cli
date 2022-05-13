package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportscheduleuploadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportscheduleuploadresponseDud struct { 
    


    


    


    

}

// Importscheduleuploadresponse
type Importscheduleuploadresponse struct { 
    // UploadKey - The key to pass to the secondary request to start processing of the upload
    UploadKey string `json:"uploadKey"`


    // Url - The url to which to PUT the upload body
    Url string `json:"url"`


    // Headers - Required headers for the PUT request to the url
    Headers map[string]string `json:"headers"`


    // UploadBodySchema - Always null. Defines the schema of the json body to be PUT to the url. The json body should be gzip encoded before uploading
    UploadBodySchema Importscheduleuploadschema `json:"uploadBodySchema"`

}

// String returns a JSON representation of the model
func (o *Importscheduleuploadresponse) String() string {
    
    
     o.Headers = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importscheduleuploadresponse) MarshalJSON() ([]byte, error) {
    type Alias Importscheduleuploadresponse

    if ImportscheduleuploadresponseMarshalled {
        return []byte("{}"), nil
    }
    ImportscheduleuploadresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        Url string `json:"url"`
        
        Headers map[string]string `json:"headers"`
        
        UploadBodySchema Importscheduleuploadschema `json:"uploadBodySchema"`
        *Alias
    }{

        


        


        
        Headers: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

