package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuestcategoryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuestcategoryresponseDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ParentCategory Guestcategoryreference `json:"parentCategory"`


    SelfUri string `json:"selfUri"`

}

// Guestcategoryresponse
type Guestcategoryresponse struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // Description - The description for the category.
    Description string `json:"description"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Guestcategoryresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guestcategoryresponse) MarshalJSON() ([]byte, error) {
    type Alias Guestcategoryresponse

    if GuestcategoryresponseMarshalled {
        return []byte("{}"), nil
    }
    GuestcategoryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

