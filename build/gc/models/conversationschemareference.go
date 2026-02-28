package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationschemareferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationschemareferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Conversationschemareference
type Conversationschemareference struct { 
    // Id - The id of the schema.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Conversationschemareference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationschemareference) MarshalJSON() ([]byte, error) {
    type Alias Conversationschemareference

    if ConversationschemareferenceMarshalled {
        return []byte("{}"), nil
    }
    ConversationschemareferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

