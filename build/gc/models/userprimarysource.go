package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserprimarysourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserprimarysourceDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Userprimarysource
type Userprimarysource struct { 
    


    // Name
    Name string `json:"name"`


    // SourceId - The id of the source
    SourceId string `json:"sourceId"`


    // Registered - Whether or not the source is registered
    Registered bool `json:"registered"`


    

}

// String returns a JSON representation of the model
func (o *Userprimarysource) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userprimarysource) MarshalJSON() ([]byte, error) {
    type Alias Userprimarysource

    if UserprimarysourceMarshalled {
        return []byte("{}"), nil
    }
    UserprimarysourceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SourceId string `json:"sourceId"`
        
        Registered bool `json:"registered"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

