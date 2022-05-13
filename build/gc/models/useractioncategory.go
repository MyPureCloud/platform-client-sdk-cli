package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractioncategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractioncategoryDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Useractioncategory
type Useractioncategory struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Useractioncategory) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractioncategory) MarshalJSON() ([]byte, error) {
    type Alias Useractioncategory

    if UseractioncategoryMarshalled {
        return []byte("{}"), nil
    }
    UseractioncategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

