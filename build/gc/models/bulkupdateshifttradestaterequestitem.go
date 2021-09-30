package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateshifttradestaterequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateshifttradestaterequestitemDud struct { 
    


    


    

}

// Bulkupdateshifttradestaterequestitem
type Bulkupdateshifttradestaterequestitem struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // State - The new state to set on the shift trade
    State string `json:"state"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestaterequestitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateshifttradestaterequestitem) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateshifttradestaterequestitem

    if BulkupdateshifttradestaterequestitemMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateshifttradestaterequestitemMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

