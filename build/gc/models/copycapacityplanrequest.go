package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopycapacityplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopycapacityplanrequestDud struct { 
    


    


    


    


    


    

}

// Copycapacityplanrequest
type Copycapacityplanrequest struct { 
    // Name - The name of the new capacity plan
    Name string `json:"name"`


    // Description - Description of the new capacity plan
    Description string `json:"description"`


    // StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`


    // Forecast - The generated forecast for the capacity plan. Uses forecast from original capacity plan if not specified
    Forecast Valuewrapperbushorttermforecastreference `json:"forecast"`


    // ImportedForecast - The externally imported forecast for the capacity plan. Uses forecast from original capacity plan if not specified
    ImportedForecast Valuewrappercapacityplanimportedforecastrequest `json:"importedForecast"`

}

// String returns a JSON representation of the model
func (o *Copycapacityplanrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copycapacityplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Copycapacityplanrequest

    if CopycapacityplanrequestMarshalled {
        return []byte("{}"), nil
    }
    CopycapacityplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        
        Forecast Valuewrapperbushorttermforecastreference `json:"forecast"`
        
        ImportedForecast Valuewrappercapacityplanimportedforecastrequest `json:"importedForecast"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

