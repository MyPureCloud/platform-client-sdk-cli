package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialercontactidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialercontactidDud struct { 
    


    

}

// Dialercontactid
type Dialercontactid struct { 
    // Id
    Id string `json:"id"`


    // ContactListId
    ContactListId string `json:"contactListId"`

}

// String returns a JSON representation of the model
func (o *Dialercontactid) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialercontactid) MarshalJSON() ([]byte, error) {
    type Alias Dialercontactid

    if DialercontactidMarshalled {
        return []byte("{}"), nil
    }
    DialercontactidMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        ContactListId string `json:"contactListId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

