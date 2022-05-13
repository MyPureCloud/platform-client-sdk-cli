package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeextendedcategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeextendedcategoryDud struct { 
    Id string `json:"id"`


    


    


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    LanguageCode string `json:"languageCode"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Parent Knowledgecategory `json:"parent"`


    Children []Knowledgecategory `json:"children"`


    SelfUri string `json:"selfUri"`

}

// Knowledgeextendedcategory
type Knowledgeextendedcategory struct { 
    


    // Name - Category name
    Name string `json:"name"`


    // Description - Category description
    Description string `json:"description"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeextendedcategory) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeextendedcategory) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeextendedcategory

    if KnowledgeextendedcategoryMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeextendedcategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

