package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryresponseDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ParentCategory *Categoryreference `json:"parentCategory"`


    DocumentCount int `json:"documentCount"`


    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    SelfUri string `json:"selfUri"`

}

// Categoryresponse
type Categoryresponse struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // Description - The description for the category.
    Description string `json:"description"`


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Categoryresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryresponse) MarshalJSON() ([]byte, error) {
    type Alias Categoryresponse

    if CategoryresponseMarshalled {
        return []byte("{}"), nil
    }
    CategoryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

