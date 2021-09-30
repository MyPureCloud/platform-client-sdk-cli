package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitiatescreenrecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitiatescreenrecordingDud struct { 
    


    


    

}

// Initiatescreenrecording
type Initiatescreenrecording struct { 
    // RecordACW
    RecordACW bool `json:"recordACW"`


    // ArchiveRetention
    ArchiveRetention Archiveretention `json:"archiveRetention"`


    // DeleteRetention
    DeleteRetention Deleteretention `json:"deleteRetention"`

}

// String returns a JSON representation of the model
func (o *Initiatescreenrecording) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initiatescreenrecording) MarshalJSON() ([]byte, error) {
    type Alias Initiatescreenrecording

    if InitiatescreenrecordingMarshalled {
        return []byte("{}"), nil
    }
    InitiatescreenrecordingMarshalled = true

    return json.Marshal(&struct { 
        RecordACW bool `json:"recordACW"`
        
        ArchiveRetention Archiveretention `json:"archiveRetention"`
        
        DeleteRetention Deleteretention `json:"deleteRetention"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

