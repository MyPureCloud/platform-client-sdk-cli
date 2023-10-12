package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExtensionpooldivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExtensionpooldivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Extensionpooldivisionview
type Extensionpooldivisionview struct { 
    // Id - The extension pool identifier
    Id string `json:"id"`


    // Name - The start number of the extension pool.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Extensionpooldivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Extensionpooldivisionview) MarshalJSON() ([]byte, error) {
    type Alias Extensionpooldivisionview

    if ExtensionpooldivisionviewMarshalled {
        return []byte("{}"), nil
    }
    ExtensionpooldivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

