package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshifttraderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshifttraderesponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Alternativeshifttraderesponse
type Alternativeshifttraderesponse struct { 
    


    // ShiftOfferJobId - The job ID of the alternative shift offer listing, from which the trade was chosen
    ShiftOfferJobId string `json:"shiftOfferJobId"`


    // ExistingShifts - The existing shifts from the offer, may be empty
    ExistingShifts []Alternativeshiftagentscheduledshift `json:"existingShifts"`


    // OfferedShifts - The offered shifts from the offer, may be empty
    OfferedShifts []Alternativeshiftagentscheduledshift `json:"offeredShifts"`


    // Schedule - The existing schedule information associated with the trade
    Schedule Alternativeshiftschedulelookup `json:"schedule"`


    // ManagementUnit - The management unit of this alternative shift trade request
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // User - The user who submitted the trade request
    User Userreference `json:"user"`


    // WeekDate - The start week date of the associated schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // ExpirationDate - The date when the trade will expire in ISO-8601 format. The trade cannot be approved after expiration
    ExpirationDate time.Time `json:"expirationDate"`


    // State - The state of this alternative shift trade
    State string `json:"state"`


    // ProcessingStatus - The processing status of this alternative shift trade
    ProcessingStatus string `json:"processingStatus"`


    // SystemDateReviewed - The timestamp of when the trade request was reviewed by the system in ISO-8601 format
    SystemDateReviewed time.Time `json:"systemDateReviewed"`


    // AdminDateReviewed - The timestamp of when the trade request was reviewed by an admin in ISO-8601 format
    AdminDateReviewed time.Time `json:"adminDateReviewed"`


    // AdminReviewedBy - The admin who reviewed this alternative shift trade after system denial
    AdminReviewedBy Userreference `json:"adminReviewedBy"`


    // Violations - A list of trade match violations
    Violations []string `json:"violations"`


    // Metadata - Version metadata for this alternative shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Alternativeshifttraderesponse) String() string {
    
     o.ExistingShifts = []Alternativeshiftagentscheduledshift{{}} 
     o.OfferedShifts = []Alternativeshiftagentscheduledshift{{}} 
    
    
    
    
    
    
    
    
    
    
     o.Violations = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshifttraderesponse) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshifttraderesponse

    if AlternativeshifttraderesponseMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshifttraderesponseMarshalled = true

    return json.Marshal(&struct {
        
        ShiftOfferJobId string `json:"shiftOfferJobId"`
        
        ExistingShifts []Alternativeshiftagentscheduledshift `json:"existingShifts"`
        
        OfferedShifts []Alternativeshiftagentscheduledshift `json:"offeredShifts"`
        
        Schedule Alternativeshiftschedulelookup `json:"schedule"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        User Userreference `json:"user"`
        
        WeekDate time.Time `json:"weekDate"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        
        State string `json:"state"`
        
        ProcessingStatus string `json:"processingStatus"`
        
        SystemDateReviewed time.Time `json:"systemDateReviewed"`
        
        AdminDateReviewed time.Time `json:"adminDateReviewed"`
        
        AdminReviewedBy Userreference `json:"adminReviewedBy"`
        
        Violations []string `json:"violations"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        
        ExistingShifts: []Alternativeshiftagentscheduledshift{{}},
        


        
        OfferedShifts: []Alternativeshiftagentscheduledshift{{}},
        


        


        


        


        


        


        


        


        


        


        


        
        Violations: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

