package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommentcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommentcreateDud struct { 
    

}

// Commentcreate
type Commentcreate struct { 
    // Content - Comment body, which cannot be empty and has a maximum length of 4,096 characters.
    Content string `json:"content"`

}

// String returns a JSON representation of the model
func (o *Commentcreate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commentcreate) MarshalJSON() ([]byte, error) {
    type Alias Commentcreate

    if CommentcreateMarshalled {
        return []byte("{}"), nil
    }
    CommentcreateMarshalled = true

    return json.Marshal(&struct {
        
        Content string `json:"content"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

