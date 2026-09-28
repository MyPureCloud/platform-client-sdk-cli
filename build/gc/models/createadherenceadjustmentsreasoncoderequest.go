package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateadherenceadjustmentsreasoncoderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateadherenceadjustmentsreasoncoderequestDud struct { 
    


    

}

// Createadherenceadjustmentsreasoncoderequest
type Createadherenceadjustmentsreasoncoderequest struct { 
    // Name - The display name of the reason code
    Name string `json:"name"`


    // State - The state of the reason code
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Createadherenceadjustmentsreasoncoderequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createadherenceadjustmentsreasoncoderequest) MarshalJSON() ([]byte, error) {
    type Alias Createadherenceadjustmentsreasoncoderequest

    if CreateadherenceadjustmentsreasoncoderequestMarshalled {
        return []byte("{}"), nil
    }
    CreateadherenceadjustmentsreasoncoderequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        State string `json:"state"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

