package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CobrowseconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CobrowseconversationDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Cobrowseconversation
type Cobrowseconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Cobrowsemediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    

}

// String returns a JSON representation of the model
func (o *Cobrowseconversation) String() string {
    
    
    
    
    
    
    
    
     o.Participants = []Cobrowsemediaparticipant{{}} 
    
    
    
     o.OtherMediaUris = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cobrowseconversation) MarshalJSON() ([]byte, error) {
    type Alias Cobrowseconversation

    if CobrowseconversationMarshalled {
        return []byte("{}"), nil
    }
    CobrowseconversationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Participants []Cobrowsemediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Participants: []Cobrowsemediaparticipant{{}},
        

        

        
        OtherMediaUris: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

