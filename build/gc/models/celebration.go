package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CelebrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CelebrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Celebration
type Celebration struct { 
    


    // Recipient - The Recipient of the celebration
    Recipient Userreference `json:"recipient"`


    // CreatedBy - The creator of the celebration
    CreatedBy Userreference `json:"createdBy"`


    // DateCreated - The date the celebration was created on. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // VarType - The Celebration Type
    VarType string `json:"type"`


    // Title - The Celebration title
    Title string `json:"title"`


    // Note - The Celebration note
    Note string `json:"note"`


    // SourceEntity - The celebration's source entity
    SourceEntity Sourceentity `json:"sourceEntity"`


    

}

// String returns a JSON representation of the model
func (o *Celebration) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Celebration) MarshalJSON() ([]byte, error) {
    type Alias Celebration

    if CelebrationMarshalled {
        return []byte("{}"), nil
    }
    CelebrationMarshalled = true

    return json.Marshal(&struct {
        
        Recipient Userreference `json:"recipient"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        VarType string `json:"type"`
        
        Title string `json:"title"`
        
        Note string `json:"note"`
        
        SourceEntity Sourceentity `json:"sourceEntity"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

