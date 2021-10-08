package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailorganizationpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailorganizationpolicyDud struct { 
    Enabled bool `json:"enabled"`


    


    


    


    


    


    


    


    ModifiedDate time.Time `json:"modifiedDate"`

}

// Voicemailorganizationpolicy
type Voicemailorganizationpolicy struct { 
    


    // AlertTimeoutSeconds - The organization's default number of seconds to ring a user's phone before a call is transferred to voicemail
    AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`


    // PinConfiguration - The configuration for user PINs to access their voicemail from a phone
    PinConfiguration Pinconfiguration `json:"pinConfiguration"`


    // VoicemailExtension - The extension for voicemail retrieval.  The default value is *86.
    VoicemailExtension string `json:"voicemailExtension"`


    // PinRequired - If this is true, a PIN is required when accessing a user's voicemail from a phone.
    PinRequired bool `json:"pinRequired"`


    // InteractiveResponseRequired - Whether user should be prompted with a confirmation prompt when connecting to a Group Ring call
    InteractiveResponseRequired bool `json:"interactiveResponseRequired"`


    // SendEmailNotifications - Whether email notifications are sent for new voicemails in the organization. If false, new voicemail email notifications are not be sent for the organization overriding any user or group setting.
    SendEmailNotifications bool `json:"sendEmailNotifications"`


    // DisableEmailPii - Removes any PII from emails. This overrides any analogous group configuration value. This is always true if HIPAA is enabled or unknown for an organization.
    DisableEmailPii bool `json:"disableEmailPii"`


    

}

// String returns a JSON representation of the model
func (o *Voicemailorganizationpolicy) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailorganizationpolicy) MarshalJSON() ([]byte, error) {
    type Alias Voicemailorganizationpolicy

    if VoicemailorganizationpolicyMarshalled {
        return []byte("{}"), nil
    }
    VoicemailorganizationpolicyMarshalled = true

    return json.Marshal(&struct { 
        
        
        AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`
        
        PinConfiguration Pinconfiguration `json:"pinConfiguration"`
        
        VoicemailExtension string `json:"voicemailExtension"`
        
        PinRequired bool `json:"pinRequired"`
        
        InteractiveResponseRequired bool `json:"interactiveResponseRequired"`
        
        SendEmailNotifications bool `json:"sendEmailNotifications"`
        
        DisableEmailPii bool `json:"disableEmailPii"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

