package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateservicegoaltemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateservicegoaltemplateDud struct { 
    


    


    


    


    

}

// Updateservicegoaltemplate
type Updateservicegoaltemplate struct { 
    // Name - The name of the service goal template.
    Name string `json:"name"`


    // ServiceLevel - Service level targets for this service goal template
    ServiceLevel Buservicelevel `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
    AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`


    // AbandonRate - Abandon rate targets for this service goal template
    AbandonRate Buabandonrate `json:"abandonRate"`


    // Metadata - Version metadata for the service goal template
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateservicegoaltemplate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateservicegoaltemplate) MarshalJSON() ([]byte, error) {
    type Alias Updateservicegoaltemplate

    if UpdateservicegoaltemplateMarshalled {
        return []byte("{}"), nil
    }
    UpdateservicegoaltemplateMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        ServiceLevel Buservicelevel `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`
        
        AbandonRate Buabandonrate `json:"abandonRate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

