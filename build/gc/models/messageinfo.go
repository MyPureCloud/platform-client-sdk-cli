package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageinfoDud struct { 
    


    


    


    

}

// Messageinfo
type Messageinfo struct { 
    // LocalizableMessageCode - Key that can be used to localize the message.
    LocalizableMessageCode string `json:"localizableMessageCode"`


    // Message - Description of the message.
    Message string `json:"message"`


    // MessageWithParams - Message with template fields for variable replacement.
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams - Map with fields for variable replacement.
    MessageParams map[string]string `json:"messageParams"`

}

// String returns a JSON representation of the model
func (o *Messageinfo) String() string {
    
    
    
     o.MessageParams = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messageinfo) MarshalJSON() ([]byte, error) {
    type Alias Messageinfo

    if MessageinfoMarshalled {
        return []byte("{}"), nil
    }
    MessageinfoMarshalled = true

    return json.Marshal(&struct {
        
        LocalizableMessageCode string `json:"localizableMessageCode"`
        
        Message string `json:"message"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]string `json:"messageParams"`
        *Alias
    }{

        


        


        


        
        MessageParams: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

