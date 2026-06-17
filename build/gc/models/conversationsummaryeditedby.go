package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummaryeditedbyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummaryeditedbyDud struct { 
    


    

}

// Conversationsummaryeditedby
type Conversationsummaryeditedby struct { 
    // User - The user that edited the summary.
    User Addressableentityref `json:"user"`


    // Purpose - The participant purpose of the user.
    Purpose string `json:"purpose"`

}

// String returns a JSON representation of the model
func (o *Conversationsummaryeditedby) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummaryeditedby) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummaryeditedby

    if ConversationsummaryeditedbyMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummaryeditedbyMarshalled = true

    return json.Marshal(&struct {
        
        User Addressableentityref `json:"user"`
        
        Purpose string `json:"purpose"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

