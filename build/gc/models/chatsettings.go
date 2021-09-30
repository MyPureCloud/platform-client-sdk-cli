package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatsettingsDud struct { 
    

}

// Chatsettings
type Chatsettings struct { 
    // MessageRetentionPeriodDays - Retention time for messages in days
    MessageRetentionPeriodDays int `json:"messageRetentionPeriodDays"`

}

// String returns a JSON representation of the model
func (o *Chatsettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatsettings) MarshalJSON() ([]byte, error) {
    type Alias Chatsettings

    if ChatsettingsMarshalled {
        return []byte("{}"), nil
    }
    ChatsettingsMarshalled = true

    return json.Marshal(&struct { 
        MessageRetentionPeriodDays int `json:"messageRetentionPeriodDays"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

