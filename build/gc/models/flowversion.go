package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowversionDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    NluInfo Nluinfo `json:"nluInfo"`


    SupportedLanguages []Supportedlanguage `json:"supportedLanguages"`


    CompatibleFlowTypes []string `json:"compatibleFlowTypes"`


    SelfUri string `json:"selfUri"`

}

// Flowversion
type Flowversion struct { 
    // Id - The flow version identifier
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // CommitVersion
    CommitVersion string `json:"commitVersion"`


    // ConfigurationVersion
    ConfigurationVersion string `json:"configurationVersion"`


    // VarType
    VarType string `json:"type"`


    // Secure
    Secure bool `json:"secure"`


    // Debug
    Debug bool `json:"debug"`


    // CreatedBy
    CreatedBy User `json:"createdBy"`


    // CreatedByClient
    CreatedByClient Domainentityref `json:"createdByClient"`


    // ConfigurationUri
    ConfigurationUri string `json:"configurationUri"`


    // DateCreated
    DateCreated int `json:"dateCreated"`


    // DateCheckedIn
    DateCheckedIn int `json:"dateCheckedIn"`


    // DateSaved
    DateSaved int `json:"dateSaved"`


    // GenerationId
    GenerationId string `json:"generationId"`


    // PublishResultUri
    PublishResultUri string `json:"publishResultUri"`


    // InputSchema
    InputSchema Jsonschemadocument `json:"inputSchema"`


    // OutputSchema
    OutputSchema Jsonschemadocument `json:"outputSchema"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Flowversion) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowversion) MarshalJSON() ([]byte, error) {
    type Alias Flowversion

    if FlowversionMarshalled {
        return []byte("{}"), nil
    }
    FlowversionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        CommitVersion string `json:"commitVersion"`
        
        ConfigurationVersion string `json:"configurationVersion"`
        
        VarType string `json:"type"`
        
        Secure bool `json:"secure"`
        
        Debug bool `json:"debug"`
        
        CreatedBy User `json:"createdBy"`
        
        CreatedByClient Domainentityref `json:"createdByClient"`
        
        ConfigurationUri string `json:"configurationUri"`
        
        DateCreated int `json:"dateCreated"`
        
        DateCheckedIn int `json:"dateCheckedIn"`
        
        DateSaved int `json:"dateSaved"`
        
        GenerationId string `json:"generationId"`
        
        PublishResultUri string `json:"publishResultUri"`
        
        InputSchema Jsonschemadocument `json:"inputSchema"`
        
        OutputSchema Jsonschemadocument `json:"outputSchema"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

