package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StacategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StacategoryDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Stacategory
type Stacategory struct { 
    


    // Name
    Name string `json:"name"`


    // Description - The description of the category
    Description string `json:"description"`


    // InteractionType - The type of interaction the category will apply to
    InteractionType string `json:"interactionType"`


    // Criteria - A collection of conditions joined together by logical operation to provide more refined filtering of conversations
    Criteria Operand `json:"criteria"`


    // CreatedBy - The user who created the record
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - The creation date of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // ModifiedBy - The user who last modified the record
    ModifiedBy Addressableentityref `json:"modifiedBy"`


    // DateModified - The last modified date of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Stacategory) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stacategory) MarshalJSON() ([]byte, error) {
    type Alias Stacategory

    if StacategoryMarshalled {
        return []byte("{}"), nil
    }
    StacategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        InteractionType string `json:"interactionType"`
        
        Criteria Operand `json:"criteria"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        ModifiedBy Addressableentityref `json:"modifiedBy"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

