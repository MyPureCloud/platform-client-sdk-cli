package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsercursorentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsercursorentitylistingDud struct { 
    


    


    


    


    

}

// Usercursorentitylisting
type Usercursorentitylisting struct { 
    // Entities
    Entities []User `json:"entities"`


    // NextUri - URI to next page of users
    NextUri string `json:"nextUri"`


    // SelfUri - URI of current page of users
    SelfUri string `json:"selfUri"`


    // Results - Number of users in response
    Results int `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Usercursorentitylisting) String() string {
     o.Entities = []User{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usercursorentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Usercursorentitylisting

    if UsercursorentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UsercursorentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []User `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        Results int `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Entities: []User{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

