package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceexplanationjob
type Adherenceexplanationjob struct { 
    


    // VarType - The type of the adherence explanation job
    VarType string `json:"type"`


    // Status - The status of the adherence explanation job
    Status string `json:"status"`


    // AdherenceExplanation - The adherence explanation added or modified by the job once complete; may be null if status == 'Error'. Used if type is in [ 'AddExplanation', 'UpdateExplanation' ]
    AdherenceExplanation Adherenceexplanationresponse `json:"adherenceExplanation"`


    // DownloadUrl - A URL to fetch results of the job. Only set if status == 'Complete' and type is in [ 'QueryAgentExplanations', 'QueryBuExplanations' ]
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Error details if status == 'Error'
    VarError Errorbody `json:"error"`


    // AgentQueryResponseTemplate - Schema template for deserializing data returned from the downloadUrl. Use if type == 'QueryAgentExplanations'
    AgentQueryResponseTemplate Adherenceexplanationlistingagentqueryresponse `json:"agentQueryResponseTemplate"`


    // BuQueryResponseTemplate - Schema template for deserializing data returned from the downloadUrl. Use if type == 'QueryBuExplanations'
    BuQueryResponseTemplate Adherenceexplanationlistingbuqueryresponse `json:"buQueryResponseTemplate"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationjob) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationjob) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationjob

    if AdherenceexplanationjobMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationjobMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Status string `json:"status"`
        
        AdherenceExplanation Adherenceexplanationresponse `json:"adherenceExplanation"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        AgentQueryResponseTemplate Adherenceexplanationlistingagentqueryresponse `json:"agentQueryResponseTemplate"`
        
        BuQueryResponseTemplate Adherenceexplanationlistingbuqueryresponse `json:"buQueryResponseTemplate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

