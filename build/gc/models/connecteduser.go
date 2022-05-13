package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnecteduserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnecteduserDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Connecteduser
type Connecteduser struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Connecteduser) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connecteduser) MarshalJSON() ([]byte, error) {
    type Alias Connecteduser

    if ConnecteduserMarshalled {
        return []byte("{}"), nil
    }
    ConnecteduserMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

