package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricsoptionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricsoptionsDud struct { 
    


    


    


    

}

// Trunkmetricsoptions
type Trunkmetricsoptions struct { 
    // ProxyAddress - Server proxy address that this options array element represents.
    ProxyAddress string `json:"proxyAddress"`


    // OptionState
    OptionState bool `json:"optionState"`


    // OptionStateTime - ISO 8601 format UTC absolute date & time of the last change of the option state.
    OptionStateTime time.Time `json:"optionStateTime"`


    // ErrorInfo
    ErrorInfo Trunkerrorinfo `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsoptions) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetricsoptions) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetricsoptions

    if TrunkmetricsoptionsMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricsoptionsMarshalled = true

    return json.Marshal(&struct { 
        ProxyAddress string `json:"proxyAddress"`
        
        OptionState bool `json:"optionState"`
        
        OptionStateTime time.Time `json:"optionStateTime"`
        
        ErrorInfo Trunkerrorinfo `json:"errorInfo"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

