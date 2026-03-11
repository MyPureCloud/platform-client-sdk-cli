package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermrequirementsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermrequirementsDud struct { 
    


    


    


    

}

// Longtermrequirements
type Longtermrequirements struct { 
    // ForecastMetadata - Forecast metadata
    ForecastMetadata Forecastmetadata `json:"forecastMetadata"`


    // DateGenerationStarted - Date the generation of the requirements started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateGenerationStarted time.Time `json:"dateGenerationStarted"`


    // Months - The list of months covered by the long-term staffing requirements, formatted as yyyy-MM
    Months []string `json:"months"`


    // RequirementResults - List of planning group outputs
    RequirementResults []Planninggrouprequirementoutput `json:"requirementResults"`

}

// String returns a JSON representation of the model
func (o *Longtermrequirements) String() string {
    
    
     o.Months = []string{""} 
     o.RequirementResults = []Planninggrouprequirementoutput{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermrequirements) MarshalJSON() ([]byte, error) {
    type Alias Longtermrequirements

    if LongtermrequirementsMarshalled {
        return []byte("{}"), nil
    }
    LongtermrequirementsMarshalled = true

    return json.Marshal(&struct {
        
        ForecastMetadata Forecastmetadata `json:"forecastMetadata"`
        
        DateGenerationStarted time.Time `json:"dateGenerationStarted"`
        
        Months []string `json:"months"`
        
        RequirementResults []Planninggrouprequirementoutput `json:"requirementResults"`
        *Alias
    }{

        


        


        
        Months: []string{""},
        


        
        RequirementResults: []Planninggrouprequirementoutput{{}},
        

        Alias: (*Alias)(u),
    })
}

