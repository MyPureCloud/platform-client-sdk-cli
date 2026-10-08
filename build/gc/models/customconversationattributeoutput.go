package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomconversationattributeoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomconversationattributeoutputDud struct { 
    Schema Conversationattributeschema `json:"schema"`


    


    

}

// Customconversationattributeoutput - Conversation Custom Attribute updates applied to a single record during a guide session turn.
type Customconversationattributeoutput struct { 
    


    // RecordId - The ID of the record that was updated.
    RecordId string `json:"recordId"`


    // Updates - The attribute updates made to this record during the turn.
    Updates []Customconversationattributeupdate `json:"updates"`

}

// String returns a JSON representation of the model
func (o *Customconversationattributeoutput) String() string {
    
     o.Updates = []Customconversationattributeupdate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customconversationattributeoutput) MarshalJSON() ([]byte, error) {
    type Alias Customconversationattributeoutput

    if CustomconversationattributeoutputMarshalled {
        return []byte("{}"), nil
    }
    CustomconversationattributeoutputMarshalled = true

    return json.Marshal(&struct {
        
        RecordId string `json:"recordId"`
        
        Updates []Customconversationattributeupdate `json:"updates"`
        *Alias
    }{

        


        


        
        Updates: []Customconversationattributeupdate{{}},
        

        Alias: (*Alias)(u),
    })
}

