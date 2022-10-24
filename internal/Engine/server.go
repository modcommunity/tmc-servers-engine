package Engine

import "time"

type UserStats struct {
	Points int64 `json:"points"`
}

type User struct {
	SteamID64 string `json:"steamid64"`
	Username  string `json:"username"`
}

type Users struct {
	List []User `json:"list"`
}

type Server struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Engine         int         `json:"engine"`
	Realname       interface{} `json:"realname"`
	Laststatupdate int         `json:"laststatupdate"`
	Laststatentry  int         `json:"laststatentry"`
	Players        interface{} `json:"players"`
	Playersmax     interface{} `json:"playersmax"`
	Mapname        interface{} `json:"mapname"`
	Club           int         `json:"club"`
	ClaimKey       *string     `json:"claimkey"`
	Verified       int         `json:"verified"`
	Realclub       struct {
		Classname       string      `json:"classname"`
		ParentClassname interface{} `json:"parentClassname"`
	} `json:"realclub"`
	Gameclub     int `json:"gameclub"`
	Gameengine   int `json:"gameengine"`
	Realgameclub struct {
		Classname       string      `json:"classname"`
		ParentClassname interface{} `json:"parentClassname"`
	} `json:"realgameclub"`
	IP          string    `json:"ip"`
	Port        int       `json:"port"`
	Hostname    string    `json:"hostname"`
	Furl        string    `json:"furl"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Created     time.Time `json:"created"`
	MemberCount int       `json:"memberCount"`
	UsersOnline Users     `json:"usersonline"`
	Owner       struct {
		ID            int         `json:"id"`
		Name          string      `json:"name"`
		Title         interface{} `json:"title"`
		TimeZone      string      `json:"timeZone"`
		FormattedName string      `json:"formattedName"`
		PrimaryGroup  struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			FormattedName string `json:"formattedName"`
		} `json:"primaryGroup"`
		SecondaryGroups       []interface{} `json:"secondaryGroups"`
		Email                 string        `json:"email"`
		Joined                time.Time     `json:"joined"`
		RegistrationIPAddress string        `json:"registrationIpAddress"`
		WarningPoints         int           `json:"warningPoints"`
		ReputationPoints      int           `json:"reputationPoints"`
		PhotoURL              string        `json:"photoUrl"`
		PhotoURLIsDefault     bool          `json:"photoUrlIsDefault"`
		CoverPhotoURL         string        `json:"coverPhotoUrl"`
		ProfileURL            string        `json:"profileUrl"`
		Validating            bool          `json:"validating"`
		Posts                 int           `json:"posts"`
		LastActivity          time.Time     `json:"lastActivity"`
		LastVisit             time.Time     `json:"lastVisit"`
		LastPost              time.Time     `json:"lastPost"`
		Birthday              interface{}   `json:"birthday"`
		ProfileViews          int           `json:"profileViews"`
		CustomFields          struct {
			Num1 struct {
				Name   string `json:"name"`
				Fields struct {
					Num1 struct {
						Name  string      `json:"name"`
						Value interface{} `json:"value"`
					} `json:"1"`
				} `json:"fields"`
			} `json:"1"`
		} `json:"customFields"`
		Rank struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Icon   string `json:"icon"`
			Points int    `json:"points"`
		} `json:"rank"`
		AchievementsPoints int  `json:"achievements_points"`
		AllowAdminEmails   bool `json:"allowAdminEmails"`
		Completed          bool `json:"completed"`
	} `json:"owner"`
	Photo    string `json:"photo"`
	Featured bool   `json:"featured"`
	Paid     bool   `json:"paid"`
	Location struct {
		Lat          interface{} `json:"lat"`
		Long         interface{} `json:"long"`
		AddressLines []string    `json:"addressLines"`
		City         string      `json:"city"`
		Region       string      `json:"region"`
		Country      string      `json:"country"`
		PostalCode   string      `json:"postalCode"`
	} `json:"location"`
	About           string        `json:"about"`
	LastActivity    time.Time     `json:"lastActivity"`
	ContentCount    int           `json:"contentCount"`
	CoverPhotoURL   string        `json:"coverPhotoUrl"`
	CoverOffset     int           `json:"coverOffset"`
	CoverPhotoColor string        `json:"coverPhotoColor"`
	Members         []interface{} `json:"members"`
	Leaders         []interface{} `json:"leaders"`
	Moderators      []interface{} `json:"moderators"`
	FieldValues     []interface{} `json:"fieldValues"`
	Nodes           []interface{} `json:"nodes"`
	Approved        bool          `json:"approved"`
	JoiningFee      interface{}   `json:"joiningFee"`
	RenewalTerm     interface{}   `json:"renewalTerm"`
}

type Servers struct {
	Page         int      `json:"page"`
	PerPage      int      `json:"perPage"`
	TotalResults int      `json:"totalResults"`
	TotalPages   int      `json:"totalPages"`
	Results      []Server `json:"results"`
}
