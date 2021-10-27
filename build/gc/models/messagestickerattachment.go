package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagestickerattachmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagestickerattachmentDud struct { 
    


    

}

// Messagestickerattachment
type Messagestickerattachment struct { 
    // Url - The location of the media, useful for retrieving it
    Url string `json:"url"`


    // Id - A globally unique identifier for the media object.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Messagestickerattachment) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagestickerattachment) MarshalJSON() ([]byte, error) {
    type Alias Messagestickerattachment

    if MessagestickerattachmentMarshalled {
        return []byte("{}"), nil
    }
    MessagestickerattachmentMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

