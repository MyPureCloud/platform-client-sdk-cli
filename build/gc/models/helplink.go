package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HelplinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HelplinkDud struct { 
    Uri string `json:"uri"`


    Title string `json:"title"`


    Description string `json:"description"`

}

// Helplink - Link to a help or support resource
type Helplink struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Helplink) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Helplink) MarshalJSON() ([]byte, error) {
    type Alias Helplink

    if HelplinkMarshalled {
        return []byte("{}"), nil
    }
    HelplinkMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

