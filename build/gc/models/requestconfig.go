package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestconfigDud struct { 
    


    


    


    


    

}

// Requestconfig - Defines response components of the Action Request.
type Requestconfig struct { 
    // RequestUrlTemplate - URL that may include placeholders for requests to 3rd party service
    RequestUrlTemplate string `json:"requestUrlTemplate"`


    // RequestTemplate - Velocity template to define request body sent to 3rd party service.
    RequestTemplate string `json:"requestTemplate"`


    // RequestTemplateUri - URI to retrieve requestTemplate
    RequestTemplateUri string `json:"requestTemplateUri"`


    // RequestType - HTTP method to use for request
    RequestType string `json:"requestType"`


    // Headers - Headers to include in request in (Header Name, Value) pairs.
    Headers map[string]string `json:"headers"`

}

// String returns a JSON representation of the model
func (o *Requestconfig) String() string {
    
    
    
    
     o.Headers = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestconfig) MarshalJSON() ([]byte, error) {
    type Alias Requestconfig

    if RequestconfigMarshalled {
        return []byte("{}"), nil
    }
    RequestconfigMarshalled = true

    return json.Marshal(&struct {
        
        RequestUrlTemplate string `json:"requestUrlTemplate"`
        
        RequestTemplate string `json:"requestTemplate"`
        
        RequestTemplateUri string `json:"requestTemplateUri"`
        
        RequestType string `json:"requestType"`
        
        Headers map[string]string `json:"headers"`
        *Alias
    }{

        


        


        


        


        
        Headers: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

