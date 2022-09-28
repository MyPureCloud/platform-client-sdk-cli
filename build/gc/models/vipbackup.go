package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VipbackupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VipbackupDud struct { 
    


    

}

// Vipbackup
type Vipbackup struct { 
    // Id - Unique ID for the selected VIP Backup option. For QUEUE this is the queueId and for USER this is the userId.
    Id string `json:"id"`


    // VarType - The type of VIP Backup to use.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Vipbackup) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Vipbackup) MarshalJSON() ([]byte, error) {
    type Alias Vipbackup

    if VipbackupMarshalled {
        return []byte("{}"), nil
    }
    VipbackupMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

