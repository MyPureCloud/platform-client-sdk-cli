package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationeditedinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationeditedinputDud struct { 
    


    

}

// Conversationeditedinput
type Conversationeditedinput struct { 
    // Text - The text of the edited input.
    Text string `json:"text"`


    // DateModified - The modification date of the edited input. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Conversationeditedinput) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationeditedinput) MarshalJSON() ([]byte, error) {
    type Alias Conversationeditedinput

    if ConversationeditedinputMarshalled {
        return []byte("{}"), nil
    }
    ConversationeditedinputMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

