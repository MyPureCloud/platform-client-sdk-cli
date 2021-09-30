package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomeventattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomeventattributeDud struct { 
    


    

}

// Customeventattribute
type Customeventattribute struct { 
    // DataType - The data type of the custom attribute.
    DataType string `json:"dataType"`


    // Value - The value of the custom attribute.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Customeventattribute) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customeventattribute) MarshalJSON() ([]byte, error) {
    type Alias Customeventattribute

    if CustomeventattributeMarshalled {
        return []byte("{}"), nil
    }
    CustomeventattributeMarshalled = true

    return json.Marshal(&struct { 
        DataType string `json:"dataType"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

