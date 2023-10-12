package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExecutiondataentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExecutiondataentityDud struct { 
    


    


    

}

// Executiondataentity - Represents an individual result of an execution data lookup
type Executiondataentity struct { 
    // Id - The id of the execution requested
    Id string `json:"id"`


    // DownloadUri - A downloadable link to the execution data file.
    DownloadUri string `json:"downloadUri"`


    // Failed - If the retrieval failed (not found, no permission, etc;), this will be set true.
    Failed bool `json:"failed"`

}

// String returns a JSON representation of the model
func (o *Executiondataentity) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Executiondataentity) MarshalJSON() ([]byte, error) {
    type Alias Executiondataentity

    if ExecutiondataentityMarshalled {
        return []byte("{}"), nil
    }
    ExecutiondataentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DownloadUri string `json:"downloadUri"`
        
        Failed bool `json:"failed"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

