package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationpushnotificationmessagelabelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationpushnotificationmessagelabelDud struct { 
    


    

}

// Conversationpushnotificationmessagelabel - A Push notification message label.
type Conversationpushnotificationmessagelabel struct { 
    // Title - Title to use in the push for each language configured in the deploymentId
    Title string `json:"title"`


    // Body - Body to use in the push for each language configured in the deploymentId
    Body string `json:"body"`

}

// String returns a JSON representation of the model
func (o *Conversationpushnotificationmessagelabel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationpushnotificationmessagelabel) MarshalJSON() ([]byte, error) {
    type Alias Conversationpushnotificationmessagelabel

    if ConversationpushnotificationmessagelabelMarshalled {
        return []byte("{}"), nil
    }
    ConversationpushnotificationmessagelabelMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Body string `json:"body"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

