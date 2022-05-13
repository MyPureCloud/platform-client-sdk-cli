package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingeventcobrowseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingeventcobrowseDud struct { 
    


    


    

}

// Webmessagingeventcobrowse - A Cobrowse event.
type Webmessagingeventcobrowse struct { 
    // VarType - Describes the type of Cobrowse event.
    VarType string `json:"type"`


    // SessionId - The Cobrowse session ID.
    SessionId string `json:"sessionId"`


    // SessionJoinToken - The Cobrowse session join token.
    SessionJoinToken string `json:"sessionJoinToken"`

}

// String returns a JSON representation of the model
func (o *Webmessagingeventcobrowse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingeventcobrowse) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingeventcobrowse

    if WebmessagingeventcobrowseMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingeventcobrowseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        SessionId string `json:"sessionId"`
        
        SessionJoinToken string `json:"sessionJoinToken"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

