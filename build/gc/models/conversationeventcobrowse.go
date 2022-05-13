package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationeventcobrowseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationeventcobrowseDud struct { 
    


    


    

}

// Conversationeventcobrowse - A CoBrowse event.
type Conversationeventcobrowse struct { 
    // VarType - Describes the type of CoBrowse event.
    VarType string `json:"type"`


    // SessionId - The CoBrowse session ID.
    SessionId string `json:"sessionId"`


    // SessionJoinToken - The CoBrowse session join token.
    SessionJoinToken string `json:"sessionJoinToken"`

}

// String returns a JSON representation of the model
func (o *Conversationeventcobrowse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationeventcobrowse) MarshalJSON() ([]byte, error) {
    type Alias Conversationeventcobrowse

    if ConversationeventcobrowseMarshalled {
        return []byte("{}"), nil
    }
    ConversationeventcobrowseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        SessionId string `json:"sessionId"`
        
        SessionJoinToken string `json:"sessionJoinToken"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

