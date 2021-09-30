package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatewebchatmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatewebchatmessagerequestDud struct { 
    


    

}

// Createwebchatmessagerequest
type Createwebchatmessagerequest struct { 
    // Body - The message body. Note that message bodies are limited to 4,000 characters.
    Body string `json:"body"`


    // BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
    BodyType string `json:"bodyType"`

}

// String returns a JSON representation of the model
func (o *Createwebchatmessagerequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createwebchatmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Createwebchatmessagerequest

    if CreatewebchatmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    CreatewebchatmessagerequestMarshalled = true

    return json.Marshal(&struct { 
        Body string `json:"body"`
        
        BodyType string `json:"bodyType"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

