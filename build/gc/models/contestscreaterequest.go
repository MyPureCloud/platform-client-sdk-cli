package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscreaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contestscreaterequest
type Contestscreaterequest struct { 
    


    // Division - The division for this performance profile associate to. Only set for DEFAULT profile.
    Division Writabledivision `json:"division"`


    // Title - The Contest title
    Title string `json:"title"`


    // Description - The Contest description
    Description string `json:"description"`


    // DateStart - Start date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // WinningCriteria - The Contest winning criteria
    WinningCriteria string `json:"winningCriteria"`


    // DateAnnounced - The Contest's Announcement Datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAnnounced time.Time `json:"dateAnnounced"`


    // AnnouncementTimezone - The Contest's Announcement Timezone. Valid values are strings of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    AnnouncementTimezone string `json:"announcementTimezone"`


    // Anonymization - The Contest anonymization
    Anonymization string `json:"anonymization"`


    // Metrics - The Contest's Metrics
    Metrics []Contestmetrics `json:"metrics"`


    // Prizes - The Contest Prizes
    Prizes []Contestprizes `json:"prizes"`


    // ProfileId - The Contest profile
    ProfileId string `json:"profileId"`


    // ParticipantIds - The Contest's participants
    ParticipantIds []string `json:"participantIds"`


    

}

// String returns a JSON representation of the model
func (o *Contestscreaterequest) String() string {
    
    
    
    
    
    
    
    
    
     o.Metrics = []Contestmetrics{{}} 
     o.Prizes = []Contestprizes{{}} 
    
     o.ParticipantIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Contestscreaterequest

    if ContestscreaterequestMarshalled {
        return []byte("{}"), nil
    }
    ContestscreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Division Writabledivision `json:"division"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        WinningCriteria string `json:"winningCriteria"`
        
        DateAnnounced time.Time `json:"dateAnnounced"`
        
        AnnouncementTimezone string `json:"announcementTimezone"`
        
        Anonymization string `json:"anonymization"`
        
        Metrics []Contestmetrics `json:"metrics"`
        
        Prizes []Contestprizes `json:"prizes"`
        
        ProfileId string `json:"profileId"`
        
        ParticipantIds []string `json:"participantIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        Metrics: []Contestmetrics{{}},
        


        
        Prizes: []Contestprizes{{}},
        


        


        
        ParticipantIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

