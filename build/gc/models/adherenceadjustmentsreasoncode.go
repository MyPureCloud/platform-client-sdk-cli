package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceadjustmentsreasoncodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceadjustmentsreasoncodeDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceadjustmentsreasoncode
type Adherenceadjustmentsreasoncode struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // State - The state of the reason code
    State string `json:"state"`


    // Metadata - Version metadata for the reason code
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceadjustmentsreasoncode) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceadjustmentsreasoncode) MarshalJSON() ([]byte, error) {
    type Alias Adherenceadjustmentsreasoncode

    if AdherenceadjustmentsreasoncodeMarshalled {
        return []byte("{}"), nil
    }
    AdherenceadjustmentsreasoncodeMarshalled = true

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

