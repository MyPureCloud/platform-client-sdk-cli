package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentreactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentreactionDud struct { 
    


    

}

// Conversationcontentreaction - User reaction to public message.
type Conversationcontentreaction struct { 
    // ReactionType - Type of reaction.
    ReactionType string `json:"reactionType"`


    // Count - Number of users that reacted this way to the message.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentreaction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentreaction) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentreaction

    if ConversationcontentreactionMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentreactionMarshalled = true

    return json.Marshal(&struct {
        
        ReactionType string `json:"reactionType"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

