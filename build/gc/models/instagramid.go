package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramidDud struct { 
    


    


    

}

// Instagramid - User information for an Instagram account
type Instagramid struct { 
    // Ids - The set of scopedIds that this person has. Each scopedId is specific to an Instagram page or app that the user interacts with.
    Ids []Instagramscopedid `json:"ids"`


    // DisplayName - The displayName of the person who owns this Instagram account
    DisplayName string `json:"displayName"`


    // Handle - The handle of the person who owns this Instagram account
    Handle string `json:"handle"`

}

// String returns a JSON representation of the model
func (o *Instagramid) String() string {
     o.Ids = []Instagramscopedid{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramid) MarshalJSON() ([]byte, error) {
    type Alias Instagramid

    if InstagramidMarshalled {
        return []byte("{}"), nil
    }
    InstagramidMarshalled = true

    return json.Marshal(&struct {
        
        Ids []Instagramscopedid `json:"ids"`
        
        DisplayName string `json:"displayName"`
        
        Handle string `json:"handle"`
        *Alias
    }{

        
        Ids: []Instagramscopedid{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

