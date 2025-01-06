package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatuserrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatuserrefDud struct { 
    


    


    


    


    


    


    

}

// Chatuserref
type Chatuserref struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Jid
    Jid string `json:"jid"`


    // Inactive
    Inactive bool `json:"inactive"`


    // Integrations
    Integrations []Contact `json:"integrations"`


    // Presence
    Presence Chatpresence `json:"presence"`

}

// String returns a JSON representation of the model
func (o *Chatuserref) String() string {
    
    
    
    
    
     o.Integrations = []Contact{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatuserref) MarshalJSON() ([]byte, error) {
    type Alias Chatuserref

    if ChatuserrefMarshalled {
        return []byte("{}"), nil
    }
    ChatuserrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        Jid string `json:"jid"`
        
        Inactive bool `json:"inactive"`
        
        Integrations []Contact `json:"integrations"`
        
        Presence Chatpresence `json:"presence"`
        *Alias
    }{

        


        


        


        


        


        
        Integrations: []Contact{{}},
        


        

        Alias: (*Alias)(u),
    })
}

