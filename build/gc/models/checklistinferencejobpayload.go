package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistinferencejobpayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistinferencejobpayloadDud struct { 
    

}

// Checklistinferencejobpayload
type Checklistinferencejobpayload struct { 
    // ConversationContext - List of conversations on which checklist evaluation is to be done.
    ConversationContext []Conversationcontext `json:"conversationContext"`

}

// String returns a JSON representation of the model
func (o *Checklistinferencejobpayload) String() string {
     o.ConversationContext = []Conversationcontext{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistinferencejobpayload) MarshalJSON() ([]byte, error) {
    type Alias Checklistinferencejobpayload

    if ChecklistinferencejobpayloadMarshalled {
        return []byte("{}"), nil
    }
    ChecklistinferencejobpayloadMarshalled = true

    return json.Marshal(&struct {
        
        ConversationContext []Conversationcontext `json:"conversationContext"`
        *Alias
    }{

        
        ConversationContext: []Conversationcontext{{}},
        

        Alias: (*Alias)(u),
    })
}

