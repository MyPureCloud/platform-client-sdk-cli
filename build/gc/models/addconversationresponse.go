package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddconversationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddconversationresponseDud struct { 
    Conversation Conversationreference `json:"conversation"`


    Appointment Coachingappointmentreference `json:"appointment"`

}

// Addconversationresponse
type Addconversationresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Addconversationresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addconversationresponse) MarshalJSON() ([]byte, error) {
    type Alias Addconversationresponse

    if AddconversationresponseMarshalled {
        return []byte("{}"), nil
    }
    AddconversationresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

