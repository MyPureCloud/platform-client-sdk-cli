package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecategoryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecategoryrequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgecategoryrequest
type Knowledgecategoryrequest struct { 
    


    // Name - Category name
    Name string `json:"name"`


    // Description - Category description
    Description string `json:"description"`


    // Parent - Category parent
    Parent Documentcategoryinput `json:"parent"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgecategoryrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecategoryrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecategoryrequest

    if KnowledgecategoryrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecategoryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Parent Documentcategoryinput `json:"parent"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

