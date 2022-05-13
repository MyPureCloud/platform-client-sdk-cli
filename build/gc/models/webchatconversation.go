package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatconversationDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Webchatconversation
type Webchatconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Member - Chat Member
    Member Webchatmemberinfo `json:"member"`


    

}

// String returns a JSON representation of the model
func (o *Webchatconversation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatconversation) MarshalJSON() ([]byte, error) {
    type Alias Webchatconversation

    if WebchatconversationMarshalled {
        return []byte("{}"), nil
    }
    WebchatconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Member Webchatmemberinfo `json:"member"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

