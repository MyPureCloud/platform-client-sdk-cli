package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatconfigDud struct { 
    

}

// Webchatconfig
type Webchatconfig struct { 
    // WebChatSkin - css class to be applied to the web chat widget.
    WebChatSkin string `json:"webChatSkin"`

}

// String returns a JSON representation of the model
func (o *Webchatconfig) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatconfig) MarshalJSON() ([]byte, error) {
    type Alias Webchatconfig

    if WebchatconfigMarshalled {
        return []byte("{}"), nil
    }
    WebchatconfigMarshalled = true

    return json.Marshal(&struct { 
        WebChatSkin string `json:"webChatSkin"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

