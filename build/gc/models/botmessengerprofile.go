package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotmessengerprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotmessengerprofileDud struct { 
    


    

}

// Botmessengerprofile
type Botmessengerprofile struct { 
    // Name - Name of the Bot
    Name string `json:"name"`


    // AvatarUrl - Avatar for Bot
    AvatarUrl string `json:"avatarUrl"`

}

// String returns a JSON representation of the model
func (o *Botmessengerprofile) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botmessengerprofile) MarshalJSON() ([]byte, error) {
    type Alias Botmessengerprofile

    if BotmessengerprofileMarshalled {
        return []byte("{}"), nil
    }
    BotmessengerprofileMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        AvatarUrl string `json:"avatarUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

