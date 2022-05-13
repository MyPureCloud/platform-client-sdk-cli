package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentreactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentreactionDud struct { 
    


    

}

// Contentreaction - User reaction to public message.
type Contentreaction struct { 
    // ReactionType - Type of reaction.
    ReactionType string `json:"reactionType"`


    // Count - Number of users that reacted this way to the message.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Contentreaction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentreaction) MarshalJSON() ([]byte, error) {
    type Alias Contentreaction

    if ContentreactionMarshalled {
        return []byte("{}"), nil
    }
    ContentreactionMarshalled = true

    return json.Marshal(&struct {
        
        ReactionType string `json:"reactionType"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

