package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArticlecontentbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArticlecontentbodyDud struct { 
    LocationUrl string `json:"locationUrl"`

}

// Articlecontentbody
type Articlecontentbody struct { 
    

}

// String returns a JSON representation of the model
func (o *Articlecontentbody) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Articlecontentbody) MarshalJSON() ([]byte, error) {
    type Alias Articlecontentbody

    if ArticlecontentbodyMarshalled {
        return []byte("{}"), nil
    }
    ArticlecontentbodyMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

