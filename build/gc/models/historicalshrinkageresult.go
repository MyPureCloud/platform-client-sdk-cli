package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalshrinkageresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalshrinkageresultDud struct { 
    


    


    


    


    


    


    

}

// Historicalshrinkageresult
type Historicalshrinkageresult struct { 
    // StartDate - Beginning of the date range that was queried, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
    EndDate time.Time `json:"endDate"`


    // TotalScheduledDurationSeconds - Total duration in seconds for which agents in the management unit are scheduled
    TotalScheduledDurationSeconds int `json:"totalScheduledDurationSeconds"`


    // TotalLoggedInDurationSeconds - Total duration in seconds for which agents in the management unit are actually logged-in
    TotalLoggedInDurationSeconds int `json:"totalLoggedInDurationSeconds"`


    // AggregatedShrinkage - Aggregated shrinkage data for all the activity categories
    AggregatedShrinkage Historicalshrinkageaggregateresponse `json:"aggregatedShrinkage"`


    // ShrinkageForActivityCategories - Shrinkage for activity categories
    ShrinkageForActivityCategories []Historicalshrinkageactivitycategoryresponse `json:"shrinkageForActivityCategories"`


    // BusinessUnitIds - List of all business units of all the agents in response
    BusinessUnitIds []string `json:"businessUnitIds"`

}

// String returns a JSON representation of the model
func (o *Historicalshrinkageresult) String() string {
    
    
    
    
    
     o.ShrinkageForActivityCategories = []Historicalshrinkageactivitycategoryresponse{{}} 
     o.BusinessUnitIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalshrinkageresult) MarshalJSON() ([]byte, error) {
    type Alias Historicalshrinkageresult

    if HistoricalshrinkageresultMarshalled {
        return []byte("{}"), nil
    }
    HistoricalshrinkageresultMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        TotalScheduledDurationSeconds int `json:"totalScheduledDurationSeconds"`
        
        TotalLoggedInDurationSeconds int `json:"totalLoggedInDurationSeconds"`
        
        AggregatedShrinkage Historicalshrinkageaggregateresponse `json:"aggregatedShrinkage"`
        
        ShrinkageForActivityCategories []Historicalshrinkageactivitycategoryresponse `json:"shrinkageForActivityCategories"`
        
        BusinessUnitIds []string `json:"businessUnitIds"`
        *Alias
    }{

        


        


        


        


        


        
        ShrinkageForActivityCategories: []Historicalshrinkageactivitycategoryresponse{{}},
        


        
        BusinessUnitIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

