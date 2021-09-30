package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetricdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetricdefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Metricdefinition
type Metricdefinition struct { 
    


    // Name
    Name string `json:"name"`


    // UnitType - The type of associated metric unit
    UnitType string `json:"unitType"`


    // ShortName - An alternate name for this metric definition, often abbreviation
    ShortName string `json:"shortName"`


    // DividendMetrics - Metric names used as dividend
    DividendMetrics []string `json:"dividendMetrics"`


    // DivisorMetrics - Metric names used as divisor
    DivisorMetrics []string `json:"divisorMetrics"`


    // DefaultObjective - A predefined default objective for this metric
    DefaultObjective Defaultobjective `json:"defaultObjective"`


    // LockTemplateId - An optional field to specify if this metric definition is locked to certain template. e.g. punctuality
    LockTemplateId string `json:"lockTemplateId"`


    

}

// String returns a JSON representation of the model
func (o *Metricdefinition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.DividendMetrics = []string{""} 
    
    
    
     o.DivisorMetrics = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metricdefinition) MarshalJSON() ([]byte, error) {
    type Alias Metricdefinition

    if MetricdefinitionMarshalled {
        return []byte("{}"), nil
    }
    MetricdefinitionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        UnitType string `json:"unitType"`
        
        ShortName string `json:"shortName"`
        
        DividendMetrics []string `json:"dividendMetrics"`
        
        DivisorMetrics []string `json:"divisorMetrics"`
        
        DefaultObjective Defaultobjective `json:"defaultObjective"`
        
        LockTemplateId string `json:"lockTemplateId"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        DividendMetrics: []string{""},
        

        

        
        DivisorMetrics: []string{""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

