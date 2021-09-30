package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionDud struct { 
    


    


    


    


    


    


    


    


    

}

// Condition
type Condition struct { 
    // VarType - The type of the condition.
    VarType string `json:"type"`


    // Inverted - If true, inverts the result of evaluating this Condition. Default is false.
    Inverted bool `json:"inverted"`


    // AttributeName - An attribute name associated with this Condition. Required for a contactAttributeCondition.
    AttributeName string `json:"attributeName"`


    // Value - A value associated with this Condition. This could be text, a number, or a relative time. Not used for a DataActionCondition.
    Value string `json:"value"`


    // ValueType - The type of the value associated with this Condition. Not used for a DataActionCondition.
    ValueType string `json:"valueType"`


    // Operator - An operation with which to evaluate the Condition. Not used for a DataActionCondition.
    Operator string `json:"operator"`


    // Codes - List of wrap-up code identifiers. Required for a wrapupCondition.
    Codes []string `json:"codes"`


    // Property - A value associated with the property type of this Condition. Required for a contactPropertyCondition.
    Property string `json:"property"`


    // PropertyType - The type of the property associated with this Condition. Required for a contactPropertyCondition.
    PropertyType string `json:"propertyType"`

}

// String returns a JSON representation of the model
func (o *Condition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Codes = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Condition) MarshalJSON() ([]byte, error) {
    type Alias Condition

    if ConditionMarshalled {
        return []byte("{}"), nil
    }
    ConditionMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Inverted bool `json:"inverted"`
        
        AttributeName string `json:"attributeName"`
        
        Value string `json:"value"`
        
        ValueType string `json:"valueType"`
        
        Operator string `json:"operator"`
        
        Codes []string `json:"codes"`
        
        Property string `json:"property"`
        
        PropertyType string `json:"propertyType"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Codes: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

