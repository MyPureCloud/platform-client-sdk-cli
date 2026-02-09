package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Summarysetting
type Summarysetting struct { 
    


    // Name - Name of the summary setting.
    Name string `json:"name"`


    // Language - Language of the generated summary, e.g. en-US, it-IT.
    Language string `json:"language"`


    // SummaryType - Level of detail of the generated summary.
    SummaryType string `json:"summaryType"`


    // Format - Format of the generated summary.
    Format string `json:"format"`


    // MaskPII - Displaying PII in the generated summary.
    MaskPII Summarysettingpii `json:"maskPII"`


    // ParticipantLabels - How to refer to interaction participants in the generated summary.
    ParticipantLabels Summarysettingparticipantlabels `json:"participantLabels"`


    // PredefinedInsights - Set which insights to include in the generated summary by default.
    PredefinedInsights []string `json:"predefinedInsights"`


    // CustomEntities - Custom entity definition.
    CustomEntities []Summarysettingcustomentity `json:"customEntities"`


    // SettingType - Type of the summary setting.
    SettingType string `json:"settingType"`


    // Prompt - Custom prompt of summary setting.
    Prompt string `json:"prompt"`


    // ServiceType - Service type for summarization. Can be 'Native' for Genesys native summarization engine or 'External' for external service. If specified as 'External', integrationId must be provided.
    ServiceType string `json:"serviceType"`


    // IntegrationId - Integration ID for the external summarization service. Required when serviceType is External.
    IntegrationId string `json:"integrationId"`


    // TimeoutDuration - Timeout duration in seconds for the external summarization service request.
    TimeoutDuration int `json:"timeoutDuration"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Summarysetting) String() string {
    
    
    
    
    
    
     o.PredefinedInsights = []string{""} 
     o.CustomEntities = []Summarysettingcustomentity{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysetting) MarshalJSON() ([]byte, error) {
    type Alias Summarysetting

    if SummarysettingMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Language string `json:"language"`
        
        SummaryType string `json:"summaryType"`
        
        Format string `json:"format"`
        
        MaskPII Summarysettingpii `json:"maskPII"`
        
        ParticipantLabels Summarysettingparticipantlabels `json:"participantLabels"`
        
        PredefinedInsights []string `json:"predefinedInsights"`
        
        CustomEntities []Summarysettingcustomentity `json:"customEntities"`
        
        SettingType string `json:"settingType"`
        
        Prompt string `json:"prompt"`
        
        ServiceType string `json:"serviceType"`
        
        IntegrationId string `json:"integrationId"`
        
        TimeoutDuration int `json:"timeoutDuration"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        PredefinedInsights: []string{""},
        


        
        CustomEntities: []Summarysettingcustomentity{{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

