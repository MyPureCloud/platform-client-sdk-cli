package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailretentionpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailretentionpolicyDud struct { 
    


    

}

// Voicemailretentionpolicy - Governs how the voicemail is retained
type Voicemailretentionpolicy struct { 
    // VoicemailRetentionPolicyType - The retention policy type
    VoicemailRetentionPolicyType string `json:"voicemailRetentionPolicyType"`


    // NumberOfDays - If retentionPolicyType == RETAIN_WITH_TTL, then this value represents the number of days for the TTL
    NumberOfDays int `json:"numberOfDays"`

}

// String returns a JSON representation of the model
func (o *Voicemailretentionpolicy) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailretentionpolicy) MarshalJSON() ([]byte, error) {
    type Alias Voicemailretentionpolicy

    if VoicemailretentionpolicyMarshalled {
        return []byte("{}"), nil
    }
    VoicemailretentionpolicyMarshalled = true

    return json.Marshal(&struct { 
        VoicemailRetentionPolicyType string `json:"voicemailRetentionPolicyType"`
        
        NumberOfDays int `json:"numberOfDays"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

