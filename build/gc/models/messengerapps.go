package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessengerappsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessengerappsDud struct { 
    


    

}

// Messengerapps - The apps embedded in the messenger
type Messengerapps struct { 
    // Conversations - The conversation settings that handles chats within the messenger
    Conversations Conversationappsettings `json:"conversations"`


    // Knowledge - The knowledge base config for messenger
    Knowledge Knowledge `json:"knowledge"`

}

// String returns a JSON representation of the model
func (o *Messengerapps) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messengerapps) MarshalJSON() ([]byte, error) {
    type Alias Messengerapps

    if MessengerappsMarshalled {
        return []byte("{}"), nil
    }
    MessengerappsMarshalled = true

    return json.Marshal(&struct {
        
        Conversations Conversationappsettings `json:"conversations"`
        
        Knowledge Knowledge `json:"knowledge"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

