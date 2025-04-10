package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestsessentialsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestsessentialsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contestsessentials
type Contestsessentials struct { 
    


    // Title - The Contest title
    Title string `json:"title"`


    // Status - The Contest status
    Status string `json:"status"`


    // DateStart - Start date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // Profile - The performance profile
    Profile Contestprofile `json:"profile"`


    // ParticipantCount - The Number of participants in the contest
    ParticipantCount int `json:"participantCount"`


    // DateAnnounced - The Contest's Announcement datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAnnounced time.Time `json:"dateAnnounced"`


    // DateFinalized - The Contest's finalize datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateFinalized time.Time `json:"dateFinalized"`


    // DateCancelled - The Contest's cancelled datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCancelled time.Time `json:"dateCancelled"`


    // DateModified - The Contest's last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateScoresModified - The datetime the contest scores were last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateScoresModified time.Time `json:"dateScoresModified"`


    // Metrics - The Contest's Metrics
    Metrics []Contestmetrics `json:"metrics"`


    // RequestingParticipantContestInfo - The Most Recent Contest Info for the requesting participant
    RequestingParticipantContestInfo Contestrequesingparticipantdailyinfo `json:"requestingParticipantContestInfo"`


    

}

// String returns a JSON representation of the model
func (o *Contestsessentials) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.Metrics = []Contestmetrics{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestsessentials) MarshalJSON() ([]byte, error) {
    type Alias Contestsessentials

    if ContestsessentialsMarshalled {
        return []byte("{}"), nil
    }
    ContestsessentialsMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Status string `json:"status"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Profile Contestprofile `json:"profile"`
        
        ParticipantCount int `json:"participantCount"`
        
        DateAnnounced time.Time `json:"dateAnnounced"`
        
        DateFinalized time.Time `json:"dateFinalized"`
        
        DateCancelled time.Time `json:"dateCancelled"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateScoresModified time.Time `json:"dateScoresModified"`
        
        Metrics []Contestmetrics `json:"metrics"`
        
        RequestingParticipantContestInfo Contestrequesingparticipantdailyinfo `json:"requestingParticipantContestInfo"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Metrics: []Contestmetrics{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

