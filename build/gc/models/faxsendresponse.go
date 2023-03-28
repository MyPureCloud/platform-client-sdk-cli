package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxsendresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxsendresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Faxsendresponse
type Faxsendresponse struct { 
    


    // Name
    Name string `json:"name"`


    // UploadDestinationUri
    UploadDestinationUri string `json:"uploadDestinationUri"`


    // UploadMethodType
    UploadMethodType string `json:"uploadMethodType"`


    // Headers
    Headers map[string]string `json:"headers"`


    

}

// String returns a JSON representation of the model
func (o *Faxsendresponse) String() string {
    
    
    
     o.Headers = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxsendresponse) MarshalJSON() ([]byte, error) {
    type Alias Faxsendresponse

    if FaxsendresponseMarshalled {
        return []byte("{}"), nil
    }
    FaxsendresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UploadDestinationUri string `json:"uploadDestinationUri"`
        
        UploadMethodType string `json:"uploadMethodType"`
        
        Headers map[string]string `json:"headers"`
        *Alias
    }{

        


        


        


        


        
        Headers: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

