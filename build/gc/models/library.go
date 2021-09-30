package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LibraryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LibraryDud struct { 
    Id string `json:"id"`


    


    Version int `json:"version"`


    CreatedBy User `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    


    SelfUri string `json:"selfUri"`

}

// Library
type Library struct { 
    


    // Name - The library name.
    Name string `json:"name"`


    


    


    


    // ResponseType - This value is deprecated. Responses representing message templates may be added to any library.
    ResponseType string `json:"responseType"`


    

}

// String returns a JSON representation of the model
func (o *Library) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Library) MarshalJSON() ([]byte, error) {
    type Alias Library

    if LibraryMarshalled {
        return []byte("{}"), nil
    }
    LibraryMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        
        
        ResponseType string `json:"responseType"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

