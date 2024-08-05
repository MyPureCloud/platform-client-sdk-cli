package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentupdatealternativeshifttraderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentupdatealternativeshifttraderequestDud struct { 
    


    

}

// Agentupdatealternativeshifttraderequest
type Agentupdatealternativeshifttraderequest struct { 
    // State - The new state of this alternative shift trade
    State string `json:"state"`


    // Metadata - Version metadata for this alternative shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Agentupdatealternativeshifttraderequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentupdatealternativeshifttraderequest) MarshalJSON() ([]byte, error) {
    type Alias Agentupdatealternativeshifttraderequest

    if AgentupdatealternativeshifttraderequestMarshalled {
        return []byte("{}"), nil
    }
    AgentupdatealternativeshifttraderequestMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

