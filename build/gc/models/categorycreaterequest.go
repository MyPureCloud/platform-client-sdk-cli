package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategorycreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategorycreaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Categorycreaterequest
type Categorycreaterequest struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // ParentCategoryId
    ParentCategoryId string `json:"parentCategoryId"`


    // Description - The description for the category.
    Description string `json:"description"`


    // ExternalId - The external id associated with the category.
    ExternalId string `json:"externalId"`


    

}

// String returns a JSON representation of the model
func (o *Categorycreaterequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categorycreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Categorycreaterequest

    if CategorycreaterequestMarshalled {
        return []byte("{}"), nil
    }
    CategorycreaterequestMarshalled = true

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

