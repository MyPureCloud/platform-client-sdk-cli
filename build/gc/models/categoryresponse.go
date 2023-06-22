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


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Categoryresponse
type Categoryresponse struct { 
    


    // Name - The name of the category.
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // ExternalId
    ExternalId string `json:"externalId"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ParentCategory - The reference to category to which this category belongs to.
    ParentCategory *Categoryreference `json:"parentCategory"`


    // DocumentCount - Number of documents assigned to this category.
    DocumentCount int `json:"documentCount"`


    // KnowledgeBase - The reference to knowledge base to which the category belongs to.
    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    

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
        
        ExternalId string `json:"externalId"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ParentCategory *Categoryreference `json:"parentCategory"`
        
        DocumentCount int `json:"documentCount"`
        
        KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

