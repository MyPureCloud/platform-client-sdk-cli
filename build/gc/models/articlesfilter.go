package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArticlesfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArticlesfilterDud struct { 
    


    

}

// Articlesfilter
type Articlesfilter struct { 
    // Labels - The labels filter.
    Labels Labelsfilter `json:"labels"`


    // Categories - The categories filter.
    Categories Categoriesfilter `json:"categories"`

}

// String returns a JSON representation of the model
func (o *Articlesfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Articlesfilter) MarshalJSON() ([]byte, error) {
    type Alias Articlesfilter

    if ArticlesfilterMarshalled {
        return []byte("{}"), nil
    }
    ArticlesfilterMarshalled = true

    return json.Marshal(&struct {
        
        Labels Labelsfilter `json:"labels"`
        
        Categories Categoriesfilter `json:"categories"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

