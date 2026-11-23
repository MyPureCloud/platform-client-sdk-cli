package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Capabilities
type Capabilities struct { 
    // Actions - Architect flow actions, mapped by action ID
    Actions map[string]Capabilitiesaction `json:"actions"`


    // DataTypes - Architect flow data types, mapped by data type ID
    DataTypes map[string]Capabilitiesdatatype `json:"dataTypes"`


    // DependencyTypes - Architect dependency types, mapped by dependency type ID
    DependencyTypes map[string]Capabilitiesdependencytype `json:"dependencyTypes"`


    // Expansions - Expansion metadata. e.g. flowType all to a list of all of the flow types
    Expansions map[string]interface{} `json:"expansions"`


    // Features - Architect features status, mapped by feature ID
    Features map[string]Capabilitiesavailability `json:"features"`


    // FeatureSupport - Architect features metadata
    FeatureSupport Capabilitiesfeaturesupport `json:"featureSupport"`


    // FlowTypes - Architect flow types, mapped by flow type ID
    FlowTypes map[string]Capabilitiesflowtype `json:"flowTypes"`


    // Functions - Architect functions, mapped by function ID
    Functions map[string]Capabilitiesfunction `json:"functions"`


    // Languages - Architect languages, mapped by language ID
    Languages map[string]Capabilitieslanguage `json:"languages"`


    // Organization - Information about the organization. e.g. list of available products
    Organization Capabilitiesorganization `json:"organization"`


    // ETag
    ETag string `json:"eTag"`

}

// String returns a JSON representation of the model
func (o *Capabilities) String() string {
     o.Actions = map[string]Capabilitiesaction{"": {}} 
     o.DataTypes = map[string]Capabilitiesdatatype{"": {}} 
     o.DependencyTypes = map[string]Capabilitiesdependencytype{"": {}} 
     o.Expansions = map[string]interface{}{"": Interface{}} 
     o.Features = map[string]Capabilitiesavailability{"": {}} 
    
     o.FlowTypes = map[string]Capabilitiesflowtype{"": {}} 
     o.Functions = map[string]Capabilitiesfunction{"": {}} 
     o.Languages = map[string]Capabilitieslanguage{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilities) MarshalJSON() ([]byte, error) {
    type Alias Capabilities

    if CapabilitiesMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesMarshalled = true

    return json.Marshal(&struct {
        
        Actions map[string]Capabilitiesaction `json:"actions"`
        
        DataTypes map[string]Capabilitiesdatatype `json:"dataTypes"`
        
        DependencyTypes map[string]Capabilitiesdependencytype `json:"dependencyTypes"`
        
        Expansions map[string]interface{} `json:"expansions"`
        
        Features map[string]Capabilitiesavailability `json:"features"`
        
        FeatureSupport Capabilitiesfeaturesupport `json:"featureSupport"`
        
        FlowTypes map[string]Capabilitiesflowtype `json:"flowTypes"`
        
        Functions map[string]Capabilitiesfunction `json:"functions"`
        
        Languages map[string]Capabilitieslanguage `json:"languages"`
        
        Organization Capabilitiesorganization `json:"organization"`
        
        ETag string `json:"eTag"`
        *Alias
    }{

        
        Actions: map[string]Capabilitiesaction{"": {}},
        


        
        DataTypes: map[string]Capabilitiesdatatype{"": {}},
        


        
        DependencyTypes: map[string]Capabilitiesdependencytype{"": {}},
        


        
        Expansions: map[string]interface{}{"": Interface{}},
        


        
        Features: map[string]Capabilitiesavailability{"": {}},
        


        


        
        FlowTypes: map[string]Capabilitiesflowtype{"": {}},
        


        
        Functions: map[string]Capabilitiesfunction{"": {}},
        


        
        Languages: map[string]Capabilitieslanguage{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

