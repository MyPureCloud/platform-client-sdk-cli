package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailgrouppolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailgrouppolicyDud struct { 
    


    Group Group `json:"group"`


    


    


    


    


    


    


    


    


    


    


    

}

// Voicemailgrouppolicy
type Voicemailgrouppolicy struct { 
    // Name
    Name string `json:"name"`


    


    // Enabled - Whether voicemail is enabled for the group
    Enabled bool `json:"enabled"`


    // SendEmailNotifications - Whether email notifications are sent to group members when a new voicemail is received
    SendEmailNotifications bool `json:"sendEmailNotifications"`


    // DisableEmailPii - Removes any PII from group emails. This is overridden by the analogous organization configuration value. This is always true if HIPAA is enabled or unknown for an organization.
    DisableEmailPii bool `json:"disableEmailPii"`


    // IncludeEmailTranscriptions - Whether to include the voicemail transcription in a group notification email
    IncludeEmailTranscriptions bool `json:"includeEmailTranscriptions"`


    // LanguagePreference - The language preference for the group.  Used for group voicemail transcription
    LanguagePreference string `json:"languagePreference"`


    // RotateCallsSecs - How many seconds to ring before rotating to the next member in the group
    RotateCallsSecs int `json:"rotateCallsSecs"`


    // StopRingingAfterRotations - How many rotations to go through
    StopRingingAfterRotations int `json:"stopRingingAfterRotations"`


    // OverflowGroupId - A fallback group to contact when all of the members in this group did not answer the call.
    OverflowGroupId string `json:"overflowGroupId"`


    // GroupAlertType - Specifies if the members in this group should be contacted randomly, in a specific order, or by round-robin.
    GroupAlertType string `json:"groupAlertType"`


    // InteractiveResponsePromptId - The prompt to use when connecting a user to a Group Ring call
    InteractiveResponsePromptId string `json:"interactiveResponsePromptId"`


    // InteractiveResponseRequired - Whether user should be prompted with a confirmation prompt when connecting to a Group Ring call
    InteractiveResponseRequired bool `json:"interactiveResponseRequired"`

}

// String returns a JSON representation of the model
func (o *Voicemailgrouppolicy) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailgrouppolicy) MarshalJSON() ([]byte, error) {
    type Alias Voicemailgrouppolicy

    if VoicemailgrouppolicyMarshalled {
        return []byte("{}"), nil
    }
    VoicemailgrouppolicyMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        SendEmailNotifications bool `json:"sendEmailNotifications"`
        
        DisableEmailPii bool `json:"disableEmailPii"`
        
        IncludeEmailTranscriptions bool `json:"includeEmailTranscriptions"`
        
        LanguagePreference string `json:"languagePreference"`
        
        RotateCallsSecs int `json:"rotateCallsSecs"`
        
        StopRingingAfterRotations int `json:"stopRingingAfterRotations"`
        
        OverflowGroupId string `json:"overflowGroupId"`
        
        GroupAlertType string `json:"groupAlertType"`
        
        InteractiveResponsePromptId string `json:"interactiveResponsePromptId"`
        
        InteractiveResponseRequired bool `json:"interactiveResponseRequired"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

