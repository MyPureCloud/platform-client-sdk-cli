package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationrecipientadditionalproviderinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationrecipientadditionalproviderinfoDud struct { }

// Conversationrecipientadditionalproviderinfo
type Conversationrecipientadditionalproviderinfo struct { }

// String returns a JSON representation of the model
func (o *Conversationrecipientadditionalproviderinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationrecipientadditionalproviderinfo) MarshalJSON() ([]byte, error) {
    type Alias Conversationrecipientadditionalproviderinfo

    if ConversationrecipientadditionalproviderinfoMarshalled {
        return []byte("{}"), nil
    }
    ConversationrecipientadditionalproviderinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

