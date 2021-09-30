package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailmailboxinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailmailboxinfoDud struct { 
    UsageSizeBytes int `json:"usageSizeBytes"`


    TotalCount int `json:"totalCount"`


    UnreadCount int `json:"unreadCount"`


    DeletedCount int `json:"deletedCount"`


    CreatedDate time.Time `json:"createdDate"`


    ModifiedDate time.Time `json:"modifiedDate"`

}

// Voicemailmailboxinfo
type Voicemailmailboxinfo struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailmailboxinfo) MarshalJSON() ([]byte, error) {
    type Alias Voicemailmailboxinfo

    if VoicemailmailboxinfoMarshalled {
        return []byte("{}"), nil
    }
    VoicemailmailboxinfoMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

