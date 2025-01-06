package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatfavoriteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatfavoriteDud struct { 
    


    

}

// Chatfavorite
type Chatfavorite struct { 
    // Id - The guid id of the favorite
    Id string `json:"id"`


    // ExternalId - The external id of the favorite
    ExternalId string `json:"externalId"`

}

// String returns a JSON representation of the model
func (o *Chatfavorite) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatfavorite) MarshalJSON() ([]byte, error) {
    type Alias Chatfavorite

    if ChatfavoriteMarshalled {
        return []byte("{}"), nil
    }
    ChatfavoriteMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ExternalId string `json:"externalId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

