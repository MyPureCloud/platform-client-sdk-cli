package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatDud struct { 
    

}

// Chat
type Chat struct { 
    // JabberId
    JabberId string `json:"jabberId"`

}

// String returns a JSON representation of the model
func (o *Chat) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chat) MarshalJSON() ([]byte, error) {
    type Alias Chat

    if ChatMarshalled {
        return []byte("{}"), nil
    }
    ChatMarshalled = true

    return json.Marshal(&struct { 
        JabberId string `json:"jabberId"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

