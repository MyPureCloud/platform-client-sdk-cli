package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentsreasoncoderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentsreasoncoderequestDud struct { 
    


    


    

}

// Updateadherenceadjustmentsreasoncoderequest
type Updateadherenceadjustmentsreasoncoderequest struct { 
    // Name - The display name of the reason code
    Name string `json:"name"`


    // State - The state of the reason code
    State string `json:"state"`


    // Metadata - Version metadata for the reason code
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentsreasoncoderequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentsreasoncoderequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentsreasoncoderequest

    if UpdateadherenceadjustmentsreasoncoderequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentsreasoncoderequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

