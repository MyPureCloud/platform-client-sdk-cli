package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Widgetclientconfigv1httpMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Widgetclientconfigv1httpDud struct { 
    


    

}

// Widgetclientconfigv1http
type Widgetclientconfigv1http struct { 
    // WebChatSkin
    WebChatSkin string `json:"webChatSkin"`


    // AuthenticationUrl
    AuthenticationUrl string `json:"authenticationUrl"`

}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1http) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetclientconfigv1http) MarshalJSON() ([]byte, error) {
    type Alias Widgetclientconfigv1http

    if Widgetclientconfigv1httpMarshalled {
        return []byte("{}"), nil
    }
    Widgetclientconfigv1httpMarshalled = true

    return json.Marshal(&struct { 
        WebChatSkin string `json:"webChatSkin"`
        
        AuthenticationUrl string `json:"authenticationUrl"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

