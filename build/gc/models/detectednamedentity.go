package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DetectednamedentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DetectednamedentityDud struct { 
    Name string `json:"name"`


    EntityType string `json:"entityType"`


    Probability float64 `json:"probability"`


    Value Detectednamedentityvalue `json:"value"`

}

// Detectednamedentity
type Detectednamedentity struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Detectednamedentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Detectednamedentity) MarshalJSON() ([]byte, error) {
    type Alias Detectednamedentity

    if DetectednamedentityMarshalled {
        return []byte("{}"), nil
    }
    DetectednamedentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

