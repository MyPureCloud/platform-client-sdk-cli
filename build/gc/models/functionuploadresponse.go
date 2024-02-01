package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionuploadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionuploadresponseDud struct { 
    Url string `json:"url"`


    Headers map[string]string `json:"headers"`


    

}

// Functionuploadresponse - Action function create upload URL response.
type Functionuploadresponse struct { 
    


    


    // SignedUrlTimeoutSeconds - Upload time to live in seconds.
    SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`

}

// String returns a JSON representation of the model
func (o *Functionuploadresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionuploadresponse) MarshalJSON() ([]byte, error) {
    type Alias Functionuploadresponse

    if FunctionuploadresponseMarshalled {
        return []byte("{}"), nil
    }
    FunctionuploadresponseMarshalled = true

    return json.Marshal(&struct {
        
        SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

