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


    // DataAction - The Data Action to use for this condition. Required for a dataActionCondition.
    DataAction Domainentityref `json:"dataAction"`


    // DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data. Required for a DataActionCondition.
    DataNotFoundResolution bool `json:"dataNotFoundResolution"`


    // ContactIdField - The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionCondition.
    ContactIdField string `json:"contactIdField"`


    // CallAnalysisResultField - The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionCondition.
    CallAnalysisResultField string `json:"callAnalysisResultField"`


    // AgentWrapupField - The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionCondition.
    AgentWrapupField string `json:"agentWrapupField"`


    // ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionCondition.
    ContactColumnToDataActionFieldMappings []Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings"`


    // Predicates - A list of predicates defining the comparisons to use for this condition. Required for a dataActionCondition.
    Predicates []Dataactionconditionpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Condition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Codes = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ContactColumnToDataActionFieldMappings = []Contactcolumntodataactionfieldmapping{{}} 
    
    
    
     o.Predicates = []Dataactionconditionpredicate{{}} 
    
    

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
        
        DataAction Domainentityref `json:"dataAction"`
        
        DataNotFoundResolution bool `json:"dataNotFoundResolution"`
        
        ContactIdField string `json:"contactIdField"`
        
        CallAnalysisResultField string `json:"callAnalysisResultField"`
        
        AgentWrapupField string `json:"agentWrapupField"`
        
        ContactColumnToDataActionFieldMappings []Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings"`
        
        Predicates []Dataactionconditionpredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Codes: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ContactColumnToDataActionFieldMappings: []Contactcolumntodataactionfieldmapping{{}},
        

        

        
        Predicates: []Dataactionconditionpredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

