package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomeventattributelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomeventattributelistDud struct { 
    


    

}

// Customeventattributelist
type Customeventattributelist struct { 
    // DataType - The data type of the custom attributes.
    DataType string `json:"dataType"`


    // Values - The list of custom event attribute values.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Customeventattributelist) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customeventattributelist) MarshalJSON() ([]byte, error) {
    type Alias Customeventattributelist

    if CustomeventattributelistMarshalled {
        return []byte("{}"), nil
    }
    CustomeventattributelistMarshalled = true

    return json.Marshal(&struct {
        
        DataType string `json:"dataType"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

