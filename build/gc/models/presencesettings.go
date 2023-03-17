package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresencesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresencesettingsDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Presencesettings
type Presencesettings struct { 
    


    // Name
    Name string `json:"name"`


    // RestorePresenceSettings - The settings for the restore presence feature
    RestorePresenceSettings Restorepresencesettings `json:"restorePresenceSettings"`


    

}

// String returns a JSON representation of the model
func (o *Presencesettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presencesettings) MarshalJSON() ([]byte, error) {
    type Alias Presencesettings

    if PresencesettingsMarshalled {
        return []byte("{}"), nil
    }
    PresencesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        RestorePresenceSettings Restorepresencesettings `json:"restorePresenceSettings"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

