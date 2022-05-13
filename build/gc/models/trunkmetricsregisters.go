package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricsregistersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricsregistersDud struct { 
    


    


    


    

}

// Trunkmetricsregisters
type Trunkmetricsregisters struct { 
    // ProxyAddress - Server proxy address that this registers array element represents.
    ProxyAddress string `json:"proxyAddress"`


    // RegisterState - True if last REGISTER message had positive response; false if error response or no response.
    RegisterState bool `json:"registerState"`


    // RegisterStateTime - ISO 8601 format UTC absolute date & time of the last change of the register state.
    RegisterStateTime time.Time `json:"registerStateTime"`


    // ErrorInfo
    ErrorInfo Trunkerrorinfo `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsregisters) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetricsregisters) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetricsregisters

    if TrunkmetricsregistersMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricsregistersMarshalled = true

    return json.Marshal(&struct {
        
        ProxyAddress string `json:"proxyAddress"`
        
        RegisterState bool `json:"registerState"`
        
        RegisterStateTime time.Time `json:"registerStateTime"`
        
        ErrorInfo Trunkerrorinfo `json:"errorInfo"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

