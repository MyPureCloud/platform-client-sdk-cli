package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Documentreference
type Documentreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Documentreference) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentreference) MarshalJSON() ([]byte, error) {
    type Alias Documentreference

    if DocumentreferenceMarshalled {
        return []byte("{}"), nil
    }
    DocumentreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

