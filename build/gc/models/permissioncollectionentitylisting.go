package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PermissioncollectionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PermissioncollectionentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Permissioncollectionentitylisting
type Permissioncollectionentitylisting struct { 
    // Entities
    Entities []Domainpermissioncollection `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Permissioncollectionentitylisting) String() string {
     o.Entities = []Domainpermissioncollection{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Permissioncollectionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Permissioncollectionentitylisting

    if PermissioncollectionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    PermissioncollectionentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domainpermissioncollection `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domainpermissioncollection{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

