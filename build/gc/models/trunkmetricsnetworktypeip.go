package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricsnetworktypeipMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricsnetworktypeipDud struct { 
    


    

}

// Trunkmetricsnetworktypeip
type Trunkmetricsnetworktypeip struct { 
    // Address - Assigned IP Address for the interface
    Address string `json:"address"`


    // ErrorInfo - Information about the error.
    ErrorInfo Trunkerrorinfo `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsnetworktypeip) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetricsnetworktypeip) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetricsnetworktypeip

    if TrunkmetricsnetworktypeipMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricsnetworktypeipMarshalled = true

    return json.Marshal(&struct {
        
        Address string `json:"address"`
        
        ErrorInfo Trunkerrorinfo `json:"errorInfo"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

