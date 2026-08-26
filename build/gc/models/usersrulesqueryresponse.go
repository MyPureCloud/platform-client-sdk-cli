package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesqueryresponseDud struct { 
    


    


    


    


    


    

}

// Usersrulesqueryresponse
type Usersrulesqueryresponse struct { 
    // Entities
    Entities []Userreference `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Usersrulesqueryresponse) String() string {
     o.Entities = []Userreference{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesqueryresponse

    if UsersrulesqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userreference `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Userreference{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

