package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationchannellocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationchannellocationDud struct { }

// Conversationchannellocation
type Conversationchannellocation struct { }

// String returns a JSON representation of the model
func (o *Conversationchannellocation) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationchannellocation) MarshalJSON() ([]byte, error) {
    type Alias Conversationchannellocation

    if ConversationchannellocationMarshalled {
        return []byte("{}"), nil
    }
    ConversationchannellocationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

