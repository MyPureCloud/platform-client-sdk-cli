package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserrecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserrecordingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Userrecording
type Userrecording struct { 
    


    // Name
    Name string `json:"name"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // CreatedBy
    CreatedBy Domainentityref `json:"createdBy"`


    // Conversation
    Conversation Conversation `json:"conversation"`


    // ContentLength
    ContentLength int `json:"contentLength"`


    // DurationMilliseconds
    DurationMilliseconds int `json:"durationMilliseconds"`


    // Thumbnails
    Thumbnails []Documentthumbnail `json:"thumbnails"`


    // Read
    Read bool `json:"read"`


    

}

// String returns a JSON representation of the model
func (o *Userrecording) String() string {
    
    
    
    
    
    
    
    
     o.Thumbnails = []Documentthumbnail{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userrecording) MarshalJSON() ([]byte, error) {
    type Alias Userrecording

    if UserrecordingMarshalled {
        return []byte("{}"), nil
    }
    UserrecordingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Workspace Domainentityref `json:"workspace"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        Conversation Conversation `json:"conversation"`
        
        ContentLength int `json:"contentLength"`
        
        DurationMilliseconds int `json:"durationMilliseconds"`
        
        Thumbnails []Documentthumbnail `json:"thumbnails"`
        
        Read bool `json:"read"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Thumbnails: []Documentthumbnail{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

