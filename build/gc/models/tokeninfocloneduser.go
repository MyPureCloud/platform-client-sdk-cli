package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TokeninfocloneduserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TokeninfocloneduserDud struct { 
    Id string `json:"id"`


    Organization Entity `json:"organization"`

}

// Tokeninfocloneduser
type Tokeninfocloneduser struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Tokeninfocloneduser) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tokeninfocloneduser) MarshalJSON() ([]byte, error) {
    type Alias Tokeninfocloneduser

    if TokeninfocloneduserMarshalled {
        return []byte("{}"), nil
    }
    TokeninfocloneduserMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

