package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagingchannelmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagingchannelmetadataDud struct { }

// Conversationmessagingchannelmetadata - Information about the channel.
type Conversationmessagingchannelmetadata struct { }

// String returns a JSON representation of the model
func (o *Conversationmessagingchannelmetadata) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagingchannelmetadata) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagingchannelmetadata

    if ConversationmessagingchannelmetadataMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagingchannelmetadataMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

