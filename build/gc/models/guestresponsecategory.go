package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuestresponsecategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuestresponsecategoryDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    ParentCategory Guestcategoryreference `json:"parentCategory"`


    SelfUri string `json:"selfUri"`

}

// Guestresponsecategory
type Guestresponsecategory struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // ExternalId
    ExternalId string `json:"externalId"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    


    

}

// String returns a JSON representation of the model
func (o *Guestresponsecategory) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guestresponsecategory) MarshalJSON() ([]byte, error) {
    type Alias Guestresponsecategory

    if GuestresponsecategoryMarshalled {
        return []byte("{}"), nil
    }
    GuestresponsecategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ExternalId string `json:"externalId"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

