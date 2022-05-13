package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventmessageDud struct { 
    


    


    


    


    


    

}

// Eventmessage
type Eventmessage struct { 
    // Code
    Code string `json:"code"`


    // Message
    Message string `json:"message"`


    // MessageWithParams
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams
    MessageParams map[string]interface{} `json:"messageParams"`


    // DocumentationUri
    DocumentationUri string `json:"documentationUri"`


    // ResourceURIs
    ResourceURIs []string `json:"resourceURIs"`

}

// String returns a JSON representation of the model
func (o *Eventmessage) String() string {
    
    
    
     o.MessageParams = map[string]interface{}{"": Interface{}} 
    
     o.ResourceURIs = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventmessage) MarshalJSON() ([]byte, error) {
    type Alias Eventmessage

    if EventmessageMarshalled {
        return []byte("{}"), nil
    }
    EventmessageMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]interface{} `json:"messageParams"`
        
        DocumentationUri string `json:"documentationUri"`
        
        ResourceURIs []string `json:"resourceURIs"`
        *Alias
    }{

        


        


        


        
        MessageParams: map[string]interface{}{"": Interface{}},
        


        


        
        ResourceURIs: []string{""},
        

        Alias: (*Alias)(u),
    })
}

