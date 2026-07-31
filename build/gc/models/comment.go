package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Comment
type Comment struct { 
    


    // Content - The comment body.
    Content string `json:"content"`


    // User - The User who authored the comment.
    User Userreference `json:"user"`


    // ModifiedBy - The User who last modified the comment.
    ModifiedBy Userreference `json:"modifiedBy"`


    // DateCreated - The date the comment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date the comment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Comment) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Comment) MarshalJSON() ([]byte, error) {
    type Alias Comment

    if CommentMarshalled {
        return []byte("{}"), nil
    }
    CommentMarshalled = true

    return json.Marshal(&struct {
        
        Content string `json:"content"`
        
        User Userreference `json:"user"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

