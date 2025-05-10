package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserchatsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserchatsettingsDud struct { 
    


    


    

}

// Userchatsettings
type Userchatsettings struct { 
    // Muted - Whether or not to enable muting notifications
    Muted bool `json:"muted"`


    // MentionsOnly - Whether or not to enable notifications for mentions only
    MentionsOnly bool `json:"mentionsOnly"`


    // NotifyOnReactions - Whether or not to enable notifications for reactions on a user's own messages
    NotifyOnReactions bool `json:"notifyOnReactions"`

}

// String returns a JSON representation of the model
func (o *Userchatsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userchatsettings) MarshalJSON() ([]byte, error) {
    type Alias Userchatsettings

    if UserchatsettingsMarshalled {
        return []byte("{}"), nil
    }
    UserchatsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Muted bool `json:"muted"`
        
        MentionsOnly bool `json:"mentionsOnly"`
        
        NotifyOnReactions bool `json:"notifyOnReactions"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

