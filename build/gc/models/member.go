package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MemberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MemberDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Member - The associated user reference as a member of a performance profile
type Member struct { 
    // Id - The user's id
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Member) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Member) MarshalJSON() ([]byte, error) {
    type Alias Member

    if MemberMarshalled {
        return []byte("{}"), nil
    }
    MemberMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

