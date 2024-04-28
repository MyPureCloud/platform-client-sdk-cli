package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationutilizationlabelupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationutilizationlabelupdateDud struct { 
    

}

// Conversationutilizationlabelupdate
type Conversationutilizationlabelupdate struct { 
    // UtilizationLabelId - The utilization label associated with the conversation.
    UtilizationLabelId string `json:"utilizationLabelId"`

}

// String returns a JSON representation of the model
func (o *Conversationutilizationlabelupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationutilizationlabelupdate) MarshalJSON() ([]byte, error) {
    type Alias Conversationutilizationlabelupdate

    if ConversationutilizationlabelupdateMarshalled {
        return []byte("{}"), nil
    }
    ConversationutilizationlabelupdateMarshalled = true

    return json.Marshal(&struct {
        
        UtilizationLabelId string `json:"utilizationLabelId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

