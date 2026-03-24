package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefilterDud struct { 
    


    

}

// V3sourcefilter
type V3sourcefilter struct { 
    // Site - The site from which to sync data.
    Site string `json:"site"`


    // Folders - The folders from which to sync data.
    Folders []string `json:"folders"`

}

// String returns a JSON representation of the model
func (o *V3sourcefilter) String() string {
    
     o.Folders = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefilter) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefilter

    if V3sourcefilterMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefilterMarshalled = true

    return json.Marshal(&struct {
        
        Site string `json:"site"`
        
        Folders []string `json:"folders"`
        *Alias
    }{

        


        
        Folders: []string{""},
        

        Alias: (*Alias)(u),
    })
}

