package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryreferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Categoryreference
type Categoryreference struct { 
    // Id - The globally unique identifier for the category.
    Id string `json:"id"`


    // Name - Category name.
    Name string `json:"name"`


    // ParentCategory - The reference to category to which this category belongs to.
    ParentCategory *Categoryreference `json:"parentCategory"`


    

}

// String returns a JSON representation of the model
func (o *Categoryreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryreference) MarshalJSON() ([]byte, error) {
    type Alias Categoryreference

    if CategoryreferenceMarshalled {
        return []byte("{}"), nil
    }
    CategoryreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ParentCategory *Categoryreference `json:"parentCategory"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

