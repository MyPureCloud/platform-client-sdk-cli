package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateservicegoaltemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateservicegoaltemplateDud struct { 
    


    


    


    

}

// Createservicegoaltemplate
type Createservicegoaltemplate struct { 
    // Name - The name of the service goal template.
    Name string `json:"name"`


    // ServiceLevel - Service level targets for this service goal template
    ServiceLevel Buservicelevel `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
    AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`


    // AbandonRate - Abandon rate targets for this service goal template
    AbandonRate Buabandonrate `json:"abandonRate"`

}

// String returns a JSON representation of the model
func (o *Createservicegoaltemplate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createservicegoaltemplate) MarshalJSON() ([]byte, error) {
    type Alias Createservicegoaltemplate

    if CreateservicegoaltemplateMarshalled {
        return []byte("{}"), nil
    }
    CreateservicegoaltemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ServiceLevel Buservicelevel `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Buaveragespeedofanswer `json:"averageSpeedOfAnswer"`
        
        AbandonRate Buabandonrate `json:"abandonRate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

