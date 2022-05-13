package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackconversationDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Callbackconversation
type Callbackconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Callbackmediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    

}

// String returns a JSON representation of the model
func (o *Callbackconversation) String() string {
    
     o.Participants = []Callbackmediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callbackconversation) MarshalJSON() ([]byte, error) {
    type Alias Callbackconversation

    if CallbackconversationMarshalled {
        return []byte("{}"), nil
    }
    CallbackconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Callbackmediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        *Alias
    }{

        


        


        
        Participants: []Callbackmediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

