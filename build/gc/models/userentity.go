package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserentityDud struct { 
    


    


    

}

// Userentity
type Userentity struct { 
    // Id - The user handle for the account being registered (base64url-encoded binary).
    Id string `json:"id"`


    // Name - A human-palatable identifier for the account (e.g., username or email).
    Name string `json:"name"`


    // DisplayName - A human-friendly display name for the account.
    DisplayName string `json:"displayName"`

}

// String returns a JSON representation of the model
func (o *Userentity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userentity) MarshalJSON() ([]byte, error) {
    type Alias Userentity

    if UserentityMarshalled {
        return []byte("{}"), nil
    }
    UserentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        DisplayName string `json:"displayName"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

