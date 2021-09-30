package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Widgetclientconfigv1Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Widgetclientconfigv1Dud struct { 
    


    

}

// Widgetclientconfigv1
type Widgetclientconfigv1 struct { 
    // WebChatSkin
    WebChatSkin string `json:"webChatSkin"`


    // AuthenticationUrl
    AuthenticationUrl string `json:"authenticationUrl"`

}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetclientconfigv1) MarshalJSON() ([]byte, error) {
    type Alias Widgetclientconfigv1

    if Widgetclientconfigv1Marshalled {
        return []byte("{}"), nil
    }
    Widgetclientconfigv1Marshalled = true

    return json.Marshal(&struct { 
        WebChatSkin string `json:"webChatSkin"`
        
        AuthenticationUrl string `json:"authenticationUrl"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

