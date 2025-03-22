package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelectedcustomcalculationcolumnsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelectedcustomcalculationcolumnsDud struct { 
    


    


    

}

// Selectedcustomcalculationcolumns
type Selectedcustomcalculationcolumns struct { 
    // CustomCalculation - Custom calculation added as a column
    CustomCalculation Addressableentityref `json:"customCalculation"`


    // Restricted - Indicates if selected custom calculation column is deleted or access revoked for the user
    Restricted bool `json:"restricted"`


    // SoftDeleted - Is selected custom calculation column soft deleted
    SoftDeleted bool `json:"softDeleted"`

}

// String returns a JSON representation of the model
func (o *Selectedcustomcalculationcolumns) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selectedcustomcalculationcolumns) MarshalJSON() ([]byte, error) {
    type Alias Selectedcustomcalculationcolumns

    if SelectedcustomcalculationcolumnsMarshalled {
        return []byte("{}"), nil
    }
    SelectedcustomcalculationcolumnsMarshalled = true

    return json.Marshal(&struct {
        
        CustomCalculation Addressableentityref `json:"customCalculation"`
        
        Restricted bool `json:"restricted"`
        
        SoftDeleted bool `json:"softDeleted"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

