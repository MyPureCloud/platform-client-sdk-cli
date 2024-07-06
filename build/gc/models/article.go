package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArticleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArticleDud struct { 
    Title string `json:"title"`


    Uri string `json:"uri"`


    Snippets []string `json:"snippets"`


    Confidence float32 `json:"confidence"`


    Metadata map[string]Metadataattribute `json:"metadata"`


    Version Addressableentityref `json:"version"`


    Variations []Addressableentityref `json:"variations"`

}

// Article
type Article struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Article) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Article) MarshalJSON() ([]byte, error) {
    type Alias Article

    if ArticleMarshalled {
        return []byte("{}"), nil
    }
    ArticleMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

