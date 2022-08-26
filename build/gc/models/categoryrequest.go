package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryrequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Categoryrequest
type Categoryrequest struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // Description - The description for the category.
    Description string `json:"description"`


    // ParentCategoryId - The category to which this category belongs.
    ParentCategoryId string `json:"parentCategoryId"`


    

}

// String returns a JSON representation of the model
func (o *Categoryrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryrequest) MarshalJSON() ([]byte, error) {
    type Alias Categoryrequest

    if CategoryrequestMarshalled {
        return []byte("{}"), nil
    }
    CategoryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ParentCategoryId string `json:"parentCategoryId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

