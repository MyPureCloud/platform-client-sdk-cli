package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicenowsourcerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicenowsourcerequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Servicenowsourcerequest
type Servicenowsourcerequest struct { 
    


    // Name - The name of the integration source.
    Name string `json:"name"`


    // IntegrationId - The integration associated with the source.
    IntegrationId string `json:"integrationId"`


    // SchedulePeriod - The schedule period of the source in hours. Must be at least 6 and at most 48 hours.
    SchedulePeriod int `json:"schedulePeriod"`


    // Settings - The settings of the source.
    Settings Servicenowsettings `json:"settings"`


    

}

// String returns a JSON representation of the model
func (o *Servicenowsourcerequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicenowsourcerequest) MarshalJSON() ([]byte, error) {
    type Alias Servicenowsourcerequest

    if ServicenowsourcerequestMarshalled {
        return []byte("{}"), nil
    }
    ServicenowsourcerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        IntegrationId string `json:"integrationId"`
        
        SchedulePeriod int `json:"schedulePeriod"`
        
        Settings Servicenowsettings `json:"settings"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

