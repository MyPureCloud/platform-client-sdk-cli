package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastplanninggroupresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastplanninggroupresponseDud struct { 
    


    


    


    

}

// Forecastplanninggroupresponse
type Forecastplanninggroupresponse struct { 
    // Id - The ID of the planning group
    Id string `json:"id"`


    // Name - The name of the planning group
    Name string `json:"name"`


    // RoutePaths - Route path configuration for this planning group
    RoutePaths []Routepathresponse `json:"routePaths"`


    // ServiceGoalTemplate - Service goals for this planning group
    ServiceGoalTemplate Forecastservicegoaltemplateresponse `json:"serviceGoalTemplate"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupresponse) String() string {
    
    
     o.RoutePaths = []Routepathresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastplanninggroupresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastplanninggroupresponse

    if ForecastplanninggroupresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastplanninggroupresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        RoutePaths []Routepathresponse `json:"routePaths"`
        
        ServiceGoalTemplate Forecastservicegoaltemplateresponse `json:"serviceGoalTemplate"`
        *Alias
    }{

        


        


        
        RoutePaths: []Routepathresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

