package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformpolicyactionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformpolicyactionsDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Crossplatformpolicyactions
type Crossplatformpolicyactions struct { 
    // RetainRecording - true to retain the recording associated with the conversation. Default = true
    RetainRecording bool `json:"retainRecording"`


    // DeleteRecording - true to delete the recording associated with the conversation. If retainRecording = true, this will be ignored. Default = false
    DeleteRecording bool `json:"deleteRecording"`


    // AlwaysDelete - true to delete the recording associated with the conversation regardless of the values of retainRecording or deleteRecording. Default = false
    AlwaysDelete bool `json:"alwaysDelete"`


    // AssignEvaluations
    AssignEvaluations []Evaluationassignment `json:"assignEvaluations"`


    // AssignMeteredEvaluations
    AssignMeteredEvaluations []Meteredevaluationassignment `json:"assignMeteredEvaluations"`


    // AssignMeteredAssignmentByAgent
    AssignMeteredAssignmentByAgent []Meteredassignmentbyagent `json:"assignMeteredAssignmentByAgent"`


    // AssignCalibrations
    AssignCalibrations []Calibrationassignment `json:"assignCalibrations"`


    // RetentionDuration
    RetentionDuration Retentionduration `json:"retentionDuration"`


    // MediaTranscriptions
    MediaTranscriptions []Mediatranscription `json:"mediaTranscriptions"`


    // IntegrationExport - Policy action for exporting recordings using an integration to 3rd party s3.
    IntegrationExport Integrationexport `json:"integrationExport"`

}

// String returns a JSON representation of the model
func (o *Crossplatformpolicyactions) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AssignEvaluations = []Evaluationassignment{{}} 
    
    
    
     o.AssignMeteredEvaluations = []Meteredevaluationassignment{{}} 
    
    
    
     o.AssignMeteredAssignmentByAgent = []Meteredassignmentbyagent{{}} 
    
    
    
     o.AssignCalibrations = []Calibrationassignment{{}} 
    
    
    
    
    
    
    
     o.MediaTranscriptions = []Mediatranscription{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformpolicyactions) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformpolicyactions

    if CrossplatformpolicyactionsMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformpolicyactionsMarshalled = true

    return json.Marshal(&struct { 
        RetainRecording bool `json:"retainRecording"`
        
        DeleteRecording bool `json:"deleteRecording"`
        
        AlwaysDelete bool `json:"alwaysDelete"`
        
        AssignEvaluations []Evaluationassignment `json:"assignEvaluations"`
        
        AssignMeteredEvaluations []Meteredevaluationassignment `json:"assignMeteredEvaluations"`
        
        AssignMeteredAssignmentByAgent []Meteredassignmentbyagent `json:"assignMeteredAssignmentByAgent"`
        
        AssignCalibrations []Calibrationassignment `json:"assignCalibrations"`
        
        RetentionDuration Retentionduration `json:"retentionDuration"`
        
        MediaTranscriptions []Mediatranscription `json:"mediaTranscriptions"`
        
        IntegrationExport Integrationexport `json:"integrationExport"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        AssignEvaluations: []Evaluationassignment{{}},
        

        

        
        AssignMeteredEvaluations: []Meteredevaluationassignment{{}},
        

        

        
        AssignMeteredAssignmentByAgent: []Meteredassignmentbyagent{{}},
        

        

        
        AssignCalibrations: []Calibrationassignment{{}},
        

        

        

        

        
        MediaTranscriptions: []Mediatranscription{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

