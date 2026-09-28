package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentsreasoncodesbulkitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentsreasoncodesbulkitemDud struct { 
    


    


    


    

}

// Updateadherenceadjustmentsreasoncodesbulkitem
type Updateadherenceadjustmentsreasoncodesbulkitem struct { 
    // Id - The ID of the reason code to update
    Id string `json:"id"`


    // Name - The display name of the reason code
    Name string `json:"name"`


    // State - The state of the reason code
    State string `json:"state"`


    // Metadata - Version metadata for the reason code
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentsreasoncodesbulkitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentsreasoncodesbulkitem) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentsreasoncodesbulkitem

    if UpdateadherenceadjustmentsreasoncodesbulkitemMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentsreasoncodesbulkitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

