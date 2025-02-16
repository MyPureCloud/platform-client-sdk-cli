package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicytestpayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicytestpayloadDud struct { 
    

}

// Policytestpayload
type Policytestpayload struct { 
    // AttributeData - A map of attribute names to attribute type and string representation of value. All attributes returned by api/v2/authorization/policies/{policyId}/attributes are required
    AttributeData map[string]Typedattribute `json:"attributeData"`

}

// String returns a JSON representation of the model
func (o *Policytestpayload) String() string {
     o.AttributeData = map[string]Typedattribute{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policytestpayload) MarshalJSON() ([]byte, error) {
    type Alias Policytestpayload

    if PolicytestpayloadMarshalled {
        return []byte("{}"), nil
    }
    PolicytestpayloadMarshalled = true

    return json.Marshal(&struct {
        
        AttributeData map[string]Typedattribute `json:"attributeData"`
        *Alias
    }{

        
        AttributeData: map[string]Typedattribute{"": {}},
        

        Alias: (*Alias)(u),
    })
}

