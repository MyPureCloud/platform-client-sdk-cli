package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmessageuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmessageuserDud struct { 
    


    


    


    


    

}

// Chatmessageuser
type Chatmessageuser struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // DisplayName
    DisplayName string `json:"displayName"`


    // Username
    Username string `json:"username"`


    // Images
    Images []Image `json:"images"`

}

// String returns a JSON representation of the model
func (o *Chatmessageuser) String() string {
    
    
    
    
     o.Images = []Image{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmessageuser) MarshalJSON() ([]byte, error) {
    type Alias Chatmessageuser

    if ChatmessageuserMarshalled {
        return []byte("{}"), nil
    }
    ChatmessageuserMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        DisplayName string `json:"displayName"`
        
        Username string `json:"username"`
        
        Images []Image `json:"images"`
        *Alias
    }{

        


        


        


        


        
        Images: []Image{{}},
        

        Alias: (*Alias)(u),
    })
}

