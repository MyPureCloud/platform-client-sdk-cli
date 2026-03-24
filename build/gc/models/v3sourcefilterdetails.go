package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefilterdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefilterdetailsDud struct { 
    


    

}

// V3sourcefilterdetails
type V3sourcefilterdetails struct { 
    // Site - Details about the site associated with the Fabric Source.
    Site V3sourcesitedetails `json:"site"`


    // Folders - Details about the folders associated with the Fabric Source.
    Folders []V3sourcefolderdetails `json:"folders"`

}

// String returns a JSON representation of the model
func (o *V3sourcefilterdetails) String() string {
    
     o.Folders = []V3sourcefolderdetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefilterdetails) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefilterdetails

    if V3sourcefilterdetailsMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefilterdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Site V3sourcesitedetails `json:"site"`
        
        Folders []V3sourcefolderdetails `json:"folders"`
        *Alias
    }{

        


        
        Folders: []V3sourcefolderdetails{{}},
        

        Alias: (*Alias)(u),
    })
}

