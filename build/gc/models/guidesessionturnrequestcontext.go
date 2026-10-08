package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturnrequestcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturnrequestcontextDud struct { 
    


    


    

}

// Guidesessionturnrequestcontext - Context provided to a guide session turn, including conversation custom attributes and prior conversation history.
type Guidesessionturnrequestcontext struct { 
    // CustomConversationAttributes - The Conversation Custom Attributes schemas and records available for this turn.
    CustomConversationAttributes []Customconversationattributeinput `json:"customConversationAttributes"`


    // Messages - The conversation history that occurred before this guide session.
    Messages []Guidesessionmessage `json:"messages"`


    // KnowledgeQueryDetected - Whether a knowledge query was detected in the previous conversation turns.
    KnowledgeQueryDetected bool `json:"knowledgeQueryDetected"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturnrequestcontext) String() string {
     o.CustomConversationAttributes = []Customconversationattributeinput{{}} 
     o.Messages = []Guidesessionmessage{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturnrequestcontext) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturnrequestcontext

    if GuidesessionturnrequestcontextMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturnrequestcontextMarshalled = true

    return json.Marshal(&struct {
        
        CustomConversationAttributes []Customconversationattributeinput `json:"customConversationAttributes"`
        
        Messages []Guidesessionmessage `json:"messages"`
        
        KnowledgeQueryDetected bool `json:"knowledgeQueryDetected"`
        *Alias
    }{

        
        CustomConversationAttributes: []Customconversationattributeinput{{}},
        


        
        Messages: []Guidesessionmessage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

