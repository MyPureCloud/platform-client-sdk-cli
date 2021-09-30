package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookidDud struct { 
    


    

}

// Facebookid - User information for a Facebook user interacting with a page or app
type Facebookid struct { 
    // Ids - The set of scopedIds that this person has. Each scopedId is specific to a page or app that the user interacts with.
    Ids []Facebookscopedid `json:"ids"`


    // DisplayName - The displayName of this person's Facebook account. Roughly translates to user.first_name + ' ' + user.last_name in the Facebook API.
    DisplayName string `json:"displayName"`

}

// String returns a JSON representation of the model
func (o *Facebookid) String() string {
    
    
     o.Ids = []Facebookscopedid{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookid) MarshalJSON() ([]byte, error) {
    type Alias Facebookid

    if FacebookidMarshalled {
        return []byte("{}"), nil
    }
    FacebookidMarshalled = true

    return json.Marshal(&struct { 
        Ids []Facebookscopedid `json:"ids"`
        
        DisplayName string `json:"displayName"`
        
        *Alias
    }{
        

        
        Ids: []Facebookscopedid{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

