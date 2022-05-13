package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArticlecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArticlecontentDud struct { 
    Body Articlecontentbody `json:"body"`

}

// Articlecontent
type Articlecontent struct { 
    

}

// String returns a JSON representation of the model
func (o *Articlecontent) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Articlecontent) MarshalJSON() ([]byte, error) {
    type Alias Articlecontent

    if ArticlecontentMarshalled {
        return []byte("{}"), nil
    }
    ArticlecontentMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

