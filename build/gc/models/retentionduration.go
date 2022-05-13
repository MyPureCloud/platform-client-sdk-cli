package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RetentiondurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RetentiondurationDud struct { 
    


    

}

// Retentionduration
type Retentionduration struct { 
    // ArchiveRetention
    ArchiveRetention Archiveretention `json:"archiveRetention"`


    // DeleteRetention
    DeleteRetention Deleteretention `json:"deleteRetention"`

}

// String returns a JSON representation of the model
func (o *Retentionduration) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Retentionduration) MarshalJSON() ([]byte, error) {
    type Alias Retentionduration

    if RetentiondurationMarshalled {
        return []byte("{}"), nil
    }
    RetentiondurationMarshalled = true

    return json.Marshal(&struct {
        
        ArchiveRetention Archiveretention `json:"archiveRetention"`
        
        DeleteRetention Deleteretention `json:"deleteRetention"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

