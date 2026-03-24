package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefolderdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefolderdetailsDud struct { 
    


    


    

}

// V3sourcefolderdetails
type V3sourcefolderdetails struct { 
    // Id - The folder's id.
    Id string `json:"id"`


    // Name - The folder's display name.
    Name string `json:"name"`


    // FullPath - The folder's full path from the root.
    FullPath string `json:"fullPath"`

}

// String returns a JSON representation of the model
func (o *V3sourcefolderdetails) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefolderdetails) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefolderdetails

    if V3sourcefolderdetailsMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefolderdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        FullPath string `json:"fullPath"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

