package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastinputsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastinputsresponseDud struct { 
    


    


    


    

}

// Capacityplanforecastinputsresponse
type Capacityplanforecastinputsresponse struct { 
    // BusinessUnit - The business unit to which the capacity plan forecast inputs belongs
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // CapacityPlan - The capacity plan associated with these forecast inputs
    CapacityPlan Capacityplanreference `json:"capacityPlan"`


    // DownloadUrl - The URL to get the forecast inputs for the capacity plan
    DownloadUrl string `json:"downloadUrl"`


    // DownloadTemplate - Forecast inputs always come through downloadUrl, the schema included here is just for documentation
    DownloadTemplate Capacityplanforecastinputstemplate `json:"downloadTemplate"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastinputsresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastinputsresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastinputsresponse

    if CapacityplanforecastinputsresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastinputsresponseMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        CapacityPlan Capacityplanreference `json:"capacityPlan"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadTemplate Capacityplanforecastinputstemplate `json:"downloadTemplate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

