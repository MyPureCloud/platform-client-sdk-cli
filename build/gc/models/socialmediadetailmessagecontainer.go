package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediadetailmessagecontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediadetailmessagecontainerDud struct { 
    


    


    

}

// Socialmediadetailmessagecontainer
type Socialmediadetailmessagecontainer struct { 
    // Id
    Id string `json:"id"`


    // NormalizedMessage
    NormalizedMessage Conversationnormalizedmessage `json:"normalizedMessage"`


    // EscalationInfo
    EscalationInfo Socialmediamessageescalationinfo `json:"escalationInfo"`

}

// String returns a JSON representation of the model
func (o *Socialmediadetailmessagecontainer) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediadetailmessagecontainer) MarshalJSON() ([]byte, error) {
    type Alias Socialmediadetailmessagecontainer

    if SocialmediadetailmessagecontainerMarshalled {
        return []byte("{}"), nil
    }
    SocialmediadetailmessagecontainerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        NormalizedMessage Conversationnormalizedmessage `json:"normalizedMessage"`
        
        EscalationInfo Socialmediamessageescalationinfo `json:"escalationInfo"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

