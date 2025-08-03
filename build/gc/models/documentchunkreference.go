package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentchunkreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentchunkreferenceDud struct { 
    


    

}

// Documentchunkreference
type Documentchunkreference struct { 
    // Id - The globally unique identifier for the document.
    Id string `json:"id"`


    // Title - The title of the document.
    Title string `json:"title"`

}

// String returns a JSON representation of the model
func (o *Documentchunkreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentchunkreference) MarshalJSON() ([]byte, error) {
    type Alias Documentchunkreference

    if DocumentchunkreferenceMarshalled {
        return []byte("{}"), nil
    }
    DocumentchunkreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

