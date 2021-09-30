package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagestickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagestickerDud struct { 
    


    

}

// Messagesticker
type Messagesticker struct { 
    // Url - The location of the sticker, useful for retrieving it
    Url string `json:"url"`


    // Id - The unique id of the the sticker object.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Messagesticker) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagesticker) MarshalJSON() ([]byte, error) {
    type Alias Messagesticker

    if MessagestickerMarshalled {
        return []byte("{}"), nil
    }
    MessagestickerMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

