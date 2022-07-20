package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    InUse bool `json:"inUse"`


    


    SelfUri string `json:"selfUri"`

}

// Externalmetricdefinition
type Externalmetricdefinition struct { 
    


    // Name - The name of the External Metric Definition
    Name string `json:"name"`


    // Unit - The unit of the External Metric Definition
    Unit string `json:"unit"`


    // UnitDefinition - The unit definition of the External Metric Definition
    UnitDefinition string `json:"unitDefinition"`


    // Precision - The decimal precision of the External Metric Definition
    Precision int `json:"precision"`


    // DefaultObjectiveType - The default objective type of the External Metric Definition
    DefaultObjectiveType string `json:"defaultObjectiveType"`


    // RetentionMonths - The retention in months of the External Metric Definition
    RetentionMonths int `json:"retentionMonths"`


    // Enabled - True if the External Metric Definition is enabled
    Enabled bool `json:"enabled"`


    


    // DateLastRefreshed - The last date and time that the metric data was refreshed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastRefreshed time.Time `json:"dateLastRefreshed"`


    

}

// String returns a JSON representation of the model
func (o *Externalmetricdefinition) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdefinition) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdefinition

    if ExternalmetricdefinitionMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Unit string `json:"unit"`
        
        UnitDefinition string `json:"unitDefinition"`
        
        Precision int `json:"precision"`
        
        DefaultObjectiveType string `json:"defaultObjectiveType"`
        
        RetentionMonths int `json:"retentionMonths"`
        
        Enabled bool `json:"enabled"`
        
        DateLastRefreshed time.Time `json:"dateLastRefreshed"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

