package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageconversationDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messageconversation
type Messageconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Messagemediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    

}

// String returns a JSON representation of the model
func (o *Messageconversation) String() string {
    
     o.Participants = []Messagemediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messageconversation) MarshalJSON() ([]byte, error) {
    type Alias Messageconversation

    if MessageconversationMarshalled {
        return []byte("{}"), nil
    }
    MessageconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Messagemediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        *Alias
    }{

        


        


        
        Participants: []Messagemediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

