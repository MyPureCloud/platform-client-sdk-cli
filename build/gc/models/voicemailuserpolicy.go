package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailuserpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailuserpolicyDud struct { 
    Enabled bool `json:"enabled"`


    


    


    ModifiedDate time.Time `json:"modifiedDate"`


    

}

// Voicemailuserpolicy
type Voicemailuserpolicy struct { 
    


    // AlertTimeoutSeconds - The number of seconds to ring the user's phone before a call is transfered to voicemail
    AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`


    // Pin - The user's PIN to access their voicemail. This property is only used for updates and never provided otherwise to ensure security
    Pin string `json:"pin"`


    


    // SendEmailNotifications - Whether email notifications are sent to the user when a new voicemail is received
    SendEmailNotifications bool `json:"sendEmailNotifications"`

}

// String returns a JSON representation of the model
func (o *Voicemailuserpolicy) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailuserpolicy) MarshalJSON() ([]byte, error) {
    type Alias Voicemailuserpolicy

    if VoicemailuserpolicyMarshalled {
        return []byte("{}"), nil
    }
    VoicemailuserpolicyMarshalled = true

    return json.Marshal(&struct {
        
        AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`
        
        Pin string `json:"pin"`
        
        SendEmailNotifications bool `json:"sendEmailNotifications"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

