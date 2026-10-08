package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturnresponsecontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturnresponsecontextDud struct { 
    

}

// Guidesessionturnresponsecontext - Context returned from a guide session turn, including conversation custom attribute updates.
type Guidesessionturnresponsecontext struct { 
    // CustomConversationAttributes - The Conversation Custom Attributes updates made during this turn.
    CustomConversationAttributes []Customconversationattributeoutput `json:"customConversationAttributes"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturnresponsecontext) String() string {
     o.CustomConversationAttributes = []Customconversationattributeoutput{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturnresponsecontext) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturnresponsecontext

    if GuidesessionturnresponsecontextMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturnresponsecontextMarshalled = true

    return json.Marshal(&struct {
        
        CustomConversationAttributes []Customconversationattributeoutput `json:"customConversationAttributes"`
        *Alias
    }{

        
        CustomConversationAttributes: []Customconversationattributeoutput{{}},
        

        Alias: (*Alias)(u),
    })
}

