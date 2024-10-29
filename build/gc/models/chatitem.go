package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Chatitem
type Chatitem struct { 
    


    // Name
    Name string `json:"name"`


    // Open - If the chat is open
    Open bool `json:"open"`


    // Favorite - The favorite entity for the chat, only appears if the chat is favorited
    Favorite Chatfavorite `json:"favorite"`


    // Images - Avatar images for the chat
    Images []Image `json:"images"`


    // DateLastMessage - The date of the last message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastMessage time.Time `json:"dateLastMessage"`


    // DateClosed - The date the chat was closed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateClosed time.Time `json:"dateClosed"`


    // User - The other 1on1 user
    User Chatuserref `json:"user"`


    // Room - The room of the chat
    Room Room `json:"room"`


    // ChatType - The type of chat
    ChatType string `json:"chatType"`


    

}

// String returns a JSON representation of the model
func (o *Chatitem) String() string {
    
    
    
     o.Images = []Image{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatitem) MarshalJSON() ([]byte, error) {
    type Alias Chatitem

    if ChatitemMarshalled {
        return []byte("{}"), nil
    }
    ChatitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Open bool `json:"open"`
        
        Favorite Chatfavorite `json:"favorite"`
        
        Images []Image `json:"images"`
        
        DateLastMessage time.Time `json:"dateLastMessage"`
        
        DateClosed time.Time `json:"dateClosed"`
        
        User Chatuserref `json:"user"`
        
        Room Room `json:"room"`
        
        ChatType string `json:"chatType"`
        *Alias
    }{

        


        


        


        


        
        Images: []Image{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

