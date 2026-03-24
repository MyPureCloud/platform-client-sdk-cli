package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplancreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplancreateDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Caseplancreate
type Caseplancreate struct { 
    // Name - The name of the Caseplan. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // DefaultDueDurationInSeconds - The default due duration in seconds for Cases created from the Caseplan. Valid range is between 1 and 31536000 seconds.
    DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`


    // DefaultTtlSeconds - The default TTL in seconds for Cases created from the Caseplan. Valid range is between 86400 and 31536000 seconds.
    DefaultTtlSeconds int `json:"defaultTtlSeconds"`


    // ReferencePrefix - The prefix of the Caseplan reference. Valid length between 2 and 8 alphanumeric characters.
    ReferencePrefix string `json:"referencePrefix"`


    // CustomerIntentId - The ID of the customer intent associated with this Caseplan.
    CustomerIntentId string `json:"customerIntentId"`


    // Description - The description of the Caseplan. Maximum length of 512 characters.
    Description string `json:"description"`


    // DefaultCaseOwnerId - The ID of the default owner of a Case created from the Caseplan.
    DefaultCaseOwnerId string `json:"defaultCaseOwnerId"`


    // DivisionId - The ID of the division the Caseplan belongs to. Use '*' for divisionless caseplans.
    DivisionId string `json:"divisionId"`


    // DataSchemas - The schemas that define all data for cases from this Caseplan. The schema must be defined in the TaskManagement namespace.
    DataSchemas []Caseplandataschema `json:"dataSchemas"`


    // IntakeSettings - The intake format when collecting data for a case from this caseplan. There can be a maximum of 10 IntakeSettings defined for a Caseplan.
    IntakeSettings []Intakesetting `json:"intakeSettings"`

}

// String returns a JSON representation of the model
func (o *Caseplancreate) String() string {
    
    
    
    
    
    
    
    
     o.DataSchemas = []Caseplandataschema{{}} 
     o.IntakeSettings = []Intakesetting{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplancreate) MarshalJSON() ([]byte, error) {
    type Alias Caseplancreate

    if CaseplancreateMarshalled {
        return []byte("{}"), nil
    }
    CaseplancreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`
        
        DefaultTtlSeconds int `json:"defaultTtlSeconds"`
        
        ReferencePrefix string `json:"referencePrefix"`
        
        CustomerIntentId string `json:"customerIntentId"`
        
        Description string `json:"description"`
        
        DefaultCaseOwnerId string `json:"defaultCaseOwnerId"`
        
        DivisionId string `json:"divisionId"`
        
        DataSchemas []Caseplandataschema `json:"dataSchemas"`
        
        IntakeSettings []Intakesetting `json:"intakeSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        DataSchemas: []Caseplandataschema{{}},
        


        
        IntakeSettings: []Intakesetting{{}},
        

        Alias: (*Alias)(u),
    })
}

