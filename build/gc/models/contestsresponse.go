package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestsresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contestsresponse
type Contestsresponse struct { 
    


    // Division - The division for this performance profile associate to
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


    // Version - The Contest Version
    Version int `json:"version"`


    // CreatedBy - The creator of the contest
    CreatedBy Userreference `json:"createdBy"`


    // Profile - The performance profile
    Profile Contestprofile `json:"profile"`


    // Participants - The Contest's participants
    Participants []Userreference `json:"participants"`


    // Status - The Contest status
    Status string `json:"status"`


    // ParticipantCount - The Number of participants in the contest
    ParticipantCount int `json:"participantCount"`


    // DateFinalized - The Contest's finalize datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateFinalized time.Time `json:"dateFinalized"`


    // DateCancelled - The Contest's cancelled datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCancelled time.Time `json:"dateCancelled"`


    // DateModified - The Contest's last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateScoresModified - The datetime the contest scores were last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateScoresModified time.Time `json:"dateScoresModified"`


    // Winners - The Contest Winners
    Winners []Contestwinners `json:"winners"`


    // DisqualifiedAgents - The Contest's disqualified agents, returned when a contest is complete
    DisqualifiedAgents []Contestdisqualifiedagents `json:"disqualifiedAgents"`


    

}

// String returns a JSON representation of the model
func (o *Contestsresponse) String() string {
    
    
    
    
    
    
    
    
    
     o.Metrics = []Contestmetrics{{}} 
     o.Prizes = []Contestprizes{{}} 
    
    
    
     o.Participants = []Userreference{{}} 
    
    
    
    
    
    
     o.Winners = []Contestwinners{{}} 
     o.DisqualifiedAgents = []Contestdisqualifiedagents{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestsresponse) MarshalJSON() ([]byte, error) {
    type Alias Contestsresponse

    if ContestsresponseMarshalled {
        return []byte("{}"), nil
    }
    ContestsresponseMarshalled = true

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
        
        Version int `json:"version"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        Profile Contestprofile `json:"profile"`
        
        Participants []Userreference `json:"participants"`
        
        Status string `json:"status"`
        
        ParticipantCount int `json:"participantCount"`
        
        DateFinalized time.Time `json:"dateFinalized"`
        
        DateCancelled time.Time `json:"dateCancelled"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateScoresModified time.Time `json:"dateScoresModified"`
        
        Winners []Contestwinners `json:"winners"`
        
        DisqualifiedAgents []Contestdisqualifiedagents `json:"disqualifiedAgents"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        Metrics: []Contestmetrics{{}},
        


        
        Prizes: []Contestprizes{{}},
        


        


        


        


        
        Participants: []Userreference{{}},
        


        


        


        


        


        


        


        
        Winners: []Contestwinners{{}},
        


        
        DisqualifiedAgents: []Contestdisqualifiedagents{{}},
        


        

        Alias: (*Alias)(u),
    })
}

