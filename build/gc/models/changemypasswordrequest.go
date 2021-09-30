package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChangemypasswordrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChangemypasswordrequestDud struct { 
    


    

}

// Changemypasswordrequest
type Changemypasswordrequest struct { 
    // NewPassword - The new password
    NewPassword string `json:"newPassword"`


    // OldPassword - Your current password
    OldPassword string `json:"oldPassword"`

}

// String returns a JSON representation of the model
func (o *Changemypasswordrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Changemypasswordrequest) MarshalJSON() ([]byte, error) {
    type Alias Changemypasswordrequest

    if ChangemypasswordrequestMarshalled {
        return []byte("{}"), nil
    }
    ChangemypasswordrequestMarshalled = true

    return json.Marshal(&struct { 
        NewPassword string `json:"newPassword"`
        
        OldPassword string `json:"oldPassword"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

