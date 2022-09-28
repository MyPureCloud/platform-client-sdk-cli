package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ViproutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ViproutingDud struct { 
    


    


    


    


    

}

// Viprouting
type Viprouting struct { 
    // VipCallMediaSettings - VIP Settings specific to Call media.
    VipCallMediaSettings Vipcallmediasettings `json:"vipCallMediaSettings"`


    // VipEmailMediaSettings - VIP Settings specific to Email media.
    VipEmailMediaSettings Vipmediasettings `json:"vipEmailMediaSettings"`


    // VipMessageMediaSettings - VIP Settings specific to Message media.
    VipMessageMediaSettings Vipmediasettings `json:"vipMessageMediaSettings"`


    // VipMaxOwnershipTimeSeconds - The max amount of time a VIP interaction will wait for the VIP user before going to the selected backup option (in seconds). Allowable range 10 seconds - 864000 seconds (ten days).
    VipMaxOwnershipTimeSeconds int `json:"vipMaxOwnershipTimeSeconds"`


    // VipBackup - VIP queue default VIP Backup settings for unanswered VIP interactions to fallback to. VIP Backup set for a specific VIP user has preference before queue default.
    VipBackup Vipbackup `json:"vipBackup"`

}

// String returns a JSON representation of the model
func (o *Viprouting) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Viprouting) MarshalJSON() ([]byte, error) {
    type Alias Viprouting

    if ViproutingMarshalled {
        return []byte("{}"), nil
    }
    ViproutingMarshalled = true

    return json.Marshal(&struct {
        
        VipCallMediaSettings Vipcallmediasettings `json:"vipCallMediaSettings"`
        
        VipEmailMediaSettings Vipmediasettings `json:"vipEmailMediaSettings"`
        
        VipMessageMediaSettings Vipmediasettings `json:"vipMessageMediaSettings"`
        
        VipMaxOwnershipTimeSeconds int `json:"vipMaxOwnershipTimeSeconds"`
        
        VipBackup Vipbackup `json:"vipBackup"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

