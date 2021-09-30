package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingexportmetadatajobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingexportmetadatajobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reportingexportmetadatajobresponse
type Reportingexportmetadatajobresponse struct { 
    


    // Name
    Name string `json:"name"`


    // ViewType - The view type of the export metadata
    ViewType string `json:"viewType"`


    // DateLimitations - The date limitations of the export metadata
    DateLimitations string `json:"dateLimitations"`


    // RequiredFilters - The list of required filters for the export metadata
    RequiredFilters []string `json:"requiredFilters"`


    // SupportedFilters - The list of supported filters for the export metadata
    SupportedFilters []string `json:"supportedFilters"`


    // RequiredColumnIds - The list of required column ids for the export metadata
    RequiredColumnIds []string `json:"requiredColumnIds"`


    // DependentColumnIds - The list of dependent column ids for the export metadata
    DependentColumnIds map[string][]string `json:"dependentColumnIds"`


    // AvailableColumnIds - The list of available column ids for the export metadata
    AvailableColumnIds []string `json:"availableColumnIds"`


    

}

// String returns a JSON representation of the model
func (o *Reportingexportmetadatajobresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RequiredFilters = []string{""} 
    
    
    
     o.SupportedFilters = []string{""} 
    
    
    
     o.RequiredColumnIds = []string{""} 
    
    
    
     o.DependentColumnIds = map[string][]string{"": {}} 
    
    
    
     o.AvailableColumnIds = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingexportmetadatajobresponse) MarshalJSON() ([]byte, error) {
    type Alias Reportingexportmetadatajobresponse

    if ReportingexportmetadatajobresponseMarshalled {
        return []byte("{}"), nil
    }
    ReportingexportmetadatajobresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ViewType string `json:"viewType"`
        
        DateLimitations string `json:"dateLimitations"`
        
        RequiredFilters []string `json:"requiredFilters"`
        
        SupportedFilters []string `json:"supportedFilters"`
        
        RequiredColumnIds []string `json:"requiredColumnIds"`
        
        DependentColumnIds map[string][]string `json:"dependentColumnIds"`
        
        AvailableColumnIds []string `json:"availableColumnIds"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        RequiredFilters: []string{""},
        

        

        
        SupportedFilters: []string{""},
        

        

        
        RequiredColumnIds: []string{""},
        

        

        
        DependentColumnIds: map[string][]string{"": {}},
        

        

        
        AvailableColumnIds: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

