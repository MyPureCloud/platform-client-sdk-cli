package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecategoryDud struct { 
    Id string `json:"id"`


    


    


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    LanguageCode string `json:"languageCode"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Knowledgecategory
type Knowledgecategory struct { 
    


    // Name - Category name
    Name string `json:"name"`


    // Description - Category description
    Description string `json:"description"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgecategory) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecategory) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecategory

    if KnowledgecategoryMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

