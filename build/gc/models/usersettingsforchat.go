package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersettingsforchatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersettingsforchatDud struct { 
    


    


    


    

}

// Usersettingsforchat
type Usersettingsforchat struct { 
    // Muted - Whether or not to enable muting notifications
    Muted bool `json:"muted"`


    // MentionsOnly - Whether or not to enable notifications for mentions only
    MentionsOnly bool `json:"mentionsOnly"`


    // NotifyOnReactions - Whether or not to enable notifications for reactions on a user's own messages
    NotifyOnReactions bool `json:"notifyOnReactions"`


    // Mobile - Settings for mobile devices
    Mobile Mobilesettings `json:"mobile"`

}

// String returns a JSON representation of the model
func (o *Usersettingsforchat) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersettingsforchat) MarshalJSON() ([]byte, error) {
    type Alias Usersettingsforchat

    if UsersettingsforchatMarshalled {
        return []byte("{}"), nil
    }
    UsersettingsforchatMarshalled = true

    return json.Marshal(&struct {
        
        Muted bool `json:"muted"`
        
        MentionsOnly bool `json:"mentionsOnly"`
        
        NotifyOnReactions bool `json:"notifyOnReactions"`
        
        Mobile Mobilesettings `json:"mobile"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

