package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationchannelmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationchannelmetadataDud struct { }

// Conversationchannelmetadata - Information about the channel.
type Conversationchannelmetadata struct { }

// String returns a JSON representation of the model
func (o *Conversationchannelmetadata) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationchannelmetadata) MarshalJSON() ([]byte, error) {
    type Alias Conversationchannelmetadata

    if ConversationchannelmetadataMarshalled {
        return []byte("{}"), nil
    }
    ConversationchannelmetadataMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

