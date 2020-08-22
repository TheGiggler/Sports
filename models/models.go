package models

import (
		"sync"
		"time"
)

type PlayByPlay struct {
	Text   string
	Index  int
	GameID string
}

type Game struct {
	Description    string
	GameID         int
	LeagueID       int
	HomeTeamID     int
	VisitingTeamID int
}
type GamePlayByPlays []PlayByPlay

func (p PlayByPlay) GetPlayByPlay() string {
	return p.Text
}

func (p *PlayByPlay) IncrementIndex(m *sync.Mutex) {
	m.Lock()
	p.Index++
	m.Unlock()

}


type GameFeed struct {
	LastUpdatedOn time.Time `json:"lastUpdatedOn"`
	GameDate time.Time
	Games         []struct {
		Schedule struct {
			ID        int         `json:"id"`
			StartTime time.Time   `json:"startTime"`
			EndedTime interface{} `json:"endedTime"`
			AwayTeam  struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"awayTeam"`
			HomeTeam struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"homeTeam"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"venue"`
			VenueAllegiance          string      `json:"venueAllegiance"`
			ScheduleStatus           string      `json:"scheduleStatus"`
			OriginalStartTime        interface{} `json:"originalStartTime"`
			DelayedOrPostponedReason interface{} `json:"delayedOrPostponedReason"`
			PlayedStatus             string      `json:"playedStatus"`
			Attendance               interface{} `json:"attendance"`
			Officials                []struct {
				ID        int    `json:"id"`
				Title     string `json:"title"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
			} `json:"officials"`
			Broadcasters []string `json:"broadcasters"`
			Weather      struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Wind        struct {
					Speed struct {
						MilesPerHour      int `json:"milesPerHour"`
						KilometersPerHour int `json:"kilometersPerHour"`
					} `json:"speed"`
					Direction struct {
						Degrees int    `json:"degrees"`
						Label   string `json:"label"`
					} `json:"direction"`
				} `json:"wind"`
				Temperature struct {
					Fahrenheit int `json:"fahrenheit"`
					Celsius    int `json:"celsius"`
				} `json:"temperature"`
				Precipitation struct {
					Type    interface{} `json:"type"`
					Percent interface{} `json:"percent"`
					Amount  struct {
						Millimeters interface{} `json:"millimeters"`
						Centimeters interface{} `json:"centimeters"`
						Inches      interface{} `json:"inches"`
						Feet        interface{} `json:"feet"`
					} `json:"amount"`
				} `json:"precipitation"`
				HumidityPercent int `json:"humidityPercent"`
			} `json:"weather"`
		} `json:"schedule"`
		Score struct {
			CurrentInning       interface{} `json:"currentInning"`
			CurrentInningHalf   interface{} `json:"currentInningHalf"`
			CurrentIntermission int         `json:"currentIntermission"`
			PlayStatus          struct {
				Batter struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"batter"`
				FirstBaseRunner  interface{} `json:"firstBaseRunner"`
				SecondBaseRunner interface{} `json:"secondBaseRunner"`
				ThirdBaseRunner  interface{} `json:"thirdBaseRunner"`
				Pitcher          struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"pitcher"`
				Catcher struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"catcher"`
				FirstBaseman struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"firstBaseman"`
				SecondBaseman struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"secondBaseman"`
				ThirdBaseman struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"thirdBaseman"`
				ShortStop struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"shortStop"`
				LeftFielder struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"leftFielder"`
				CenterFielder struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"centerFielder"`
				RightFielder struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"rightFielder"`
				OutFielder  interface{} `json:"outFielder"`
				BallCount   int         `json:"ballCount"`
				StrikeCount int         `json:"strikeCount"`
				OutCount    int         `json:"outCount"`
			} `json:"playStatus"`
			AwayScoreTotal  int `json:"awayScoreTotal"`
			AwayHitsTotal   int `json:"awayHitsTotal"`
			AwayErrorsTotal int `json:"awayErrorsTotal"`
			HomeScoreTotal  int `json:"homeScoreTotal"`
			HomeHitsTotal   int `json:"homeHitsTotal"`
			HomeErrorsTotal int `json:"homeErrorsTotal"`
			Innings         []struct {
				InningNumber int `json:"inningNumber"`
				AwayScore    int `json:"awayScore"`
				HomeScore    int `json:"homeScore"`
			} `json:"innings"`
		} `json:"score"`
	} `json:"games"`
	References struct {
		TeamReferences []struct {
			ID           int    `json:"id"`
			City         string `json:"city"`
			Name         string `json:"name"`
			Abbreviation string `json:"abbreviation"`
			HomeVenue    struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"homeVenue"`
			TeamColoursHex      []interface{} `json:"teamColoursHex"`
			SocialMediaAccounts []struct {
				MediaType string `json:"mediaType"`
				Value     string `json:"value"`
			} `json:"socialMediaAccounts"`
			OfficialLogoImageSrc interface{} `json:"officialLogoImageSrc"`
		} `json:"teamReferences"`
		VenueReferences []struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			City           string `json:"city"`
			Country        string `json:"country"`
			GeoCoordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoCoordinates"`
			CapacitiesByEventType []struct {
				EventType string `json:"eventType"`
				Capacity  int    `json:"capacity"`
			} `json:"capacitiesByEventType"`
			PlayingSurface     string `json:"playingSurface"`
			BaseballDimensions []struct {
				DimensionType string `json:"dimensionType"`
				DistanceFeet  int    `json:"distanceFeet"`
			} `json:"baseballDimensions"`
			HasRoof            bool `json:"hasRoof"`
			HasRetractableRoof bool `json:"hasRetractableRoof"`
		} `json:"venueReferences"`
	} `json:"references"`
}