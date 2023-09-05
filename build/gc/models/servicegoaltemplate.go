package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicegoaltemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicegoaltemplateDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Servicegoaltemplate
type Servicegoaltemplate struct { 
    


    // Name
    Name string `json:"name"`


    // ServiceLevel - Service level targets for this service goal template
    ServiceLevel Buservicelevel `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
    AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`


    // AbandonRate - Abandon rate targets for this service goal template
    AbandonRate Buabandonrate `json:"abandonRate"`


    // Metadata - Version metadata for the service goal template
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // ImpactOverride - Settings controlling max percent increase and decrease of service goals for this service goal template
    ImpactOverride Servicegoaltemplateimpactoverride `json:"impactOverride"`


    

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplate) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicegoaltemplate) MarshalJSON() ([]byte, error) {
    type Alias Servicegoaltemplate

    if ServicegoaltemplateMarshalled {
        return []byte("{}"), nil
    }
    ServicegoaltemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ServiceLevel Buservicelevel `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`
        
        AbandonRate Buabandonrate `json:"abandonRate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        ImpactOverride Servicegoaltemplateimpactoverride `json:"impactOverride"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

