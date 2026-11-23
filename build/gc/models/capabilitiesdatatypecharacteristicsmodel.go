package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypecharacteristicsmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypecharacteristicsmodelDud struct { 
    


    


    

}

// Capabilitiesdatatypecharacteristicsmodel
type Capabilitiesdatatypecharacteristicsmodel struct { 
    // Equality
    Equality Capabilitiesdatatypecharacteristicmodel `json:"equality"`


    // InvariantString
    InvariantString Capabilitiesdatatypecharacteristicmodel `json:"invariantString"`


    // LocalizedString
    LocalizedString Capabilitiesdatatypecharacteristicmodel `json:"localizedString"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypecharacteristicsmodel) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypecharacteristicsmodel) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypecharacteristicsmodel

    if CapabilitiesdatatypecharacteristicsmodelMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypecharacteristicsmodelMarshalled = true

    return json.Marshal(&struct {
        
        Equality Capabilitiesdatatypecharacteristicmodel `json:"equality"`
        
        InvariantString Capabilitiesdatatypecharacteristicmodel `json:"invariantString"`
        
        LocalizedString Capabilitiesdatatypecharacteristicmodel `json:"localizedString"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

