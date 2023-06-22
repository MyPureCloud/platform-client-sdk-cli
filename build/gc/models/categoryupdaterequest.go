package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Categoryupdaterequest
type Categoryupdaterequest struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // ParentCategoryId - The category to which this category belongs.
    ParentCategoryId string `json:"parentCategoryId"`


    // Description
    Description string `json:"description"`


    // ExternalId - The external id associated with the category.
    ExternalId string `json:"externalId"`


    

}

// String returns a JSON representation of the model
func (o *Categoryupdaterequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Categoryupdaterequest

    if CategoryupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    CategoryupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ParentCategoryId string `json:"parentCategoryId"`
        
        Description string `json:"description"`
        
        ExternalId string `json:"externalId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

