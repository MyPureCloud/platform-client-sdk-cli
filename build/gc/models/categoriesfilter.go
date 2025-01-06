package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoriesfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoriesfilterDud struct { 
    

}

// Categoriesfilter
type Categoriesfilter struct { 
    // Entities - A list of categories to filter by. Articles matching any of the specified categories can be accessed.
    Entities []Categoryentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Categoriesfilter) String() string {
     o.Entities = []Categoryentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoriesfilter) MarshalJSON() ([]byte, error) {
    type Alias Categoriesfilter

    if CategoriesfilterMarshalled {
        return []byte("{}"), nil
    }
    CategoriesfilterMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Categoryentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Categoryentity{{}},
        

        Alias: (*Alias)(u),
    })
}

