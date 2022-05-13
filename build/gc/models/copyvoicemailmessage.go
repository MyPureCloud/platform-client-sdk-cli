package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyvoicemailmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyvoicemailmessageDud struct { 
    


    


    

}

// Copyvoicemailmessage - Used to copy a VoicemailMessage to either a User or a Group
type Copyvoicemailmessage struct { 
    // VoicemailMessageId - The id of the VoicemailMessage to copy
    VoicemailMessageId string `json:"voicemailMessageId"`


    // UserId - The id of the User to copy the VoicemailMessage to
    UserId string `json:"userId"`


    // GroupId - The id of the Group to copy the VoicemailMessage to
    GroupId string `json:"groupId"`

}

// String returns a JSON representation of the model
func (o *Copyvoicemailmessage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyvoicemailmessage) MarshalJSON() ([]byte, error) {
    type Alias Copyvoicemailmessage

    if CopyvoicemailmessageMarshalled {
        return []byte("{}"), nil
    }
    CopyvoicemailmessageMarshalled = true

    return json.Marshal(&struct {
        
        VoicemailMessageId string `json:"voicemailMessageId"`
        
        UserId string `json:"userId"`
        
        GroupId string `json:"groupId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

