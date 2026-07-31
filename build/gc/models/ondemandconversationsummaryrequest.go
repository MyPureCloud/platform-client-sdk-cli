package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OndemandconversationsummaryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OndemandconversationsummaryrequestDud struct { 
    

}

// Ondemandconversationsummaryrequest - Request to queue an on-demand summary for a conversation.
type Ondemandconversationsummaryrequest struct { 
    // Locale - Locale for the summary (e.g. en-us).
    Locale string `json:"locale"`

}

// String returns a JSON representation of the model
func (o *Ondemandconversationsummaryrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ondemandconversationsummaryrequest) MarshalJSON() ([]byte, error) {
    type Alias Ondemandconversationsummaryrequest

    if OndemandconversationsummaryrequestMarshalled {
        return []byte("{}"), nil
    }
    OndemandconversationsummaryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Locale string `json:"locale"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

