package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundstructuredresponsenormalizedmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundstructuredresponsenormalizedmessagecontentDud struct { 
    


    

}

// Openinboundstructuredresponsenormalizedmessagecontent - Message content element.
type Openinboundstructuredresponsenormalizedmessagecontent struct { 
    // ContentType - Type of this content element
    ContentType string `json:"contentType"`


    // ButtonResponse - Button response content.
    ButtonResponse Contentbuttonresponse `json:"buttonResponse"`

}

// String returns a JSON representation of the model
func (o *Openinboundstructuredresponsenormalizedmessagecontent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundstructuredresponsenormalizedmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Openinboundstructuredresponsenormalizedmessagecontent

    if OpeninboundstructuredresponsenormalizedmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundstructuredresponsenormalizedmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        ButtonResponse Contentbuttonresponse `json:"buttonResponse"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

