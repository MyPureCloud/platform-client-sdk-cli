package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SendmessagebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SendmessagebodyDud struct { 
    


    


    

}

// Sendmessagebody
type Sendmessagebody struct { 
    // Message - The body of the message
    Message string `json:"message"`


    // Mentions - user ids to be notified
    Mentions []string `json:"mentions"`


    // ThreadId - The thread id of the message
    ThreadId string `json:"threadId"`

}

// String returns a JSON representation of the model
func (o *Sendmessagebody) String() string {
    
     o.Mentions = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sendmessagebody) MarshalJSON() ([]byte, error) {
    type Alias Sendmessagebody

    if SendmessagebodyMarshalled {
        return []byte("{}"), nil
    }
    SendmessagebodyMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Mentions []string `json:"mentions"`
        
        ThreadId string `json:"threadId"`
        *Alias
    }{

        


        
        Mentions: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

