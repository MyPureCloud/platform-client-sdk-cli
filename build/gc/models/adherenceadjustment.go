package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceadjustmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceadjustmentDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceadjustment
type Adherenceadjustment struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Agent - The agent to whom this adherence adjustment applies
    Agent Userreference `json:"agent"`


    // ManagementUnit - The management unit to which the agent belonged when the adherence adjustment was submitted
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // BusinessUnit - The business unit to which the agent belonged when the adherence adjustment was submitted
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence adjustment in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // ReasonCode - The reason code for this adherence adjustment
    ReasonCode Adherenceadjustmentsreasoncodereference `json:"reasonCode"`


    // Status - The status of the adherence adjustment
    Status string `json:"status"`


    // Expired - Indicates if the adherence adjustment is expired
    Expired bool `json:"expired"`


    // SubmitterNotes - Notes provided by the submitter for this adherence adjustment
    SubmitterNotes string `json:"submitterNotes"`


    // ReviewerNotes - Notes provided by the reviewer for this adherence adjustment
    ReviewerNotes string `json:"reviewerNotes"`


    // ReviewedBy - The user who reviewed the adherence adjustment, if applicable. The id may be 'System' if it was an automated process
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The date the adherence adjustment was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    // Metadata - Version metadata for the adherence adjustment
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceadjustment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceadjustment) MarshalJSON() ([]byte, error) {
    type Alias Adherenceadjustment

    if AdherenceadjustmentMarshalled {
        return []byte("{}"), nil
    }
    AdherenceadjustmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Agent Userreference `json:"agent"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        ReasonCode Adherenceadjustmentsreasoncodereference `json:"reasonCode"`
        
        Status string `json:"status"`
        
        Expired bool `json:"expired"`
        
        SubmitterNotes string `json:"submitterNotes"`
        
        ReviewerNotes string `json:"reviewerNotes"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

