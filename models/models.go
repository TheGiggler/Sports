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
	HomeTeamName   string
	HomeTeamScore int
	AwayTeamID int
	AwayTeamName string
	AwayTeamScore int
	TimePeriodType TimePeriodType
	TimePeriod int
	TimeRemainingSeconds int

}

type TimePeriodType string

const(
	Period TimePeriodType = "Period"
	Inning = "Inning"
	Quarter = "Quarter"
	Half = "Half"
)


type GamePlayByPlays []PlayByPlay

func (p PlayByPlay) GetPlayByPlay() string {
	return p.Text
}

func (p *PlayByPlay) IncrementIndex(m *sync.Mutex) {
	m.Lock()
	p.Index++
	m.Unlock()

}

type Boxscore struct {
	LastUpdatedOn time.Time `json:"lastUpdatedOn"`
	Game          struct {
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
	} `json:"game"`
	Scoring struct {
		CurrentInning       int         `json:"currentInning"`
		CurrentInningHalf   string      `json:"currentInningHalf"`
		CurrentIntermission interface{} `json:"currentIntermission"`
		AwayScoreTotal      int         `json:"awayScoreTotal"`
		AwayHitsTotal       int         `json:"awayHitsTotal"`
		AwayErrorsTotal     int         `json:"awayErrorsTotal"`
		HomeScoreTotal      int         `json:"homeScoreTotal"`
		HomeHitsTotal       int         `json:"homeHitsTotal"`
		HomeErrorsTotal     int         `json:"homeErrorsTotal"`
		Innings             []struct {
			InningNumber int `json:"inningNumber"`
			AwayScore    int `json:"awayScore"`
			HomeScore    int `json:"homeScore"`
			ScoringPlays []struct {
				InningHalf string `json:"inningHalf"`
				Team       struct {
					ID           int    `json:"id"`
					Abbreviation string `json:"abbreviation"`
				} `json:"team"`
				ScoreChange     int    `json:"scoreChange"`
				AwayScore       int    `json:"awayScore"`
				HomeScore       int    `json:"homeScore"`
				PlayDescription string `json:"playDescription"`
			} `json:"scoringPlays"`
		} `json:"innings"`
	} `json:"scoring"`
	Stats struct {
		Away struct {
			TeamStats []struct {
				Batting struct {
					AtBats                       int     `json:"atBats"`
					Runs                         int     `json:"runs"`
					Hits                         int     `json:"hits"`
					SecondBaseHits               int     `json:"secondBaseHits"`
					ThirdBaseHits                int     `json:"thirdBaseHits"`
					Homeruns                     int     `json:"homeruns"`
					RunsBattedIn                 int     `json:"runsBattedIn"`
					EarnedRuns                   int     `json:"earnedRuns"`
					UnearnedRuns                 int     `json:"unearnedRuns"`
					BatterWalks                  int     `json:"batterWalks"`
					BatterSwings                 int     `json:"batterSwings"`
					BatterStrikes                int     `json:"batterStrikes"`
					BatterStrikesFoul            int     `json:"batterStrikesFoul"`
					BatterStrikesMiss            int     `json:"batterStrikesMiss"`
					BatterStrikesLooking         int     `json:"batterStrikesLooking"`
					BatterGroundBalls            int     `json:"batterGroundBalls"`
					BatterFlyBalls               int     `json:"batterFlyBalls"`
					BatterLineDrives             int     `json:"batterLineDrives"`
					BatterStrikeouts             int     `json:"batterStrikeouts"`
					Batter2SeamFastballs         int     `json:"batter2SeamFastballs"`
					Batter4SeamFastballs         int     `json:"batter4SeamFastballs"`
					BatterCurveballs             int     `json:"batterCurveballs"`
					BatterChangeups              int     `json:"batterChangeups"`
					BatterCutters                int     `json:"batterCutters"`
					BatterSliders                int     `json:"batterSliders"`
					BatterSinkers                int     `json:"batterSinkers"`
					BatterSplitters              int     `json:"batterSplitters"`
					LeftOnBase                   int     `json:"leftOnBase"`
					OpponentsLeftOnBase          int     `json:"opponentsLeftOnBase"`
					StolenBases                  int     `json:"stolenBases"`
					CaughtBaseSteals             int     `json:"caughtBaseSteals"`
					BatterStolenBasePct          float64 `json:"batterStolenBasePct"`
					BattingAvg                   float64 `json:"battingAvg"`
					BatterOnBasePct              float64 `json:"batterOnBasePct"`
					BatterSluggingPct            float64 `json:"batterSluggingPct"`
					BatterOnBasePlusSluggingPct  float64 `json:"batterOnBasePlusSluggingPct"`
					BatterIntentionalWalks       int     `json:"batterIntentionalWalks"`
					HitByPitch                   int     `json:"hitByPitch"`
					BatterSacrificeBunts         int     `json:"batterSacrificeBunts"`
					BatterSacrificeFlies         int     `json:"batterSacrificeFlies"`
					TotalBases                   int     `json:"totalBases"`
					ExtraBaseHits                int     `json:"extraBaseHits"`
					BatterDoublePlays            int     `json:"batterDoublePlays"`
					BatterTriplePlays            int     `json:"batterTriplePlays"`
					BatterTagOuts                int     `json:"batterTagOuts"`
					BatterForceOuts              int     `json:"batterForceOuts"`
					BatterPutOuts                int     `json:"batterPutOuts"`
					BatterGroundOuts             int     `json:"batterGroundOuts"`
					BatterFlyOuts                int     `json:"batterFlyOuts"`
					BatterGroundOutToFlyOutRatio float64 `json:"batterGroundOutToFlyOutRatio"`
					PitchesFaced                 int     `json:"pitchesFaced"`
					PlateAppearances             int     `json:"plateAppearances"`
					OpponentAtBats               int     `json:"opponentAtBats"`
				} `json:"batting"`
				Pitching struct {
					EarnedRunAvg                  float64 `json:"earnedRunAvg"`
					InningsPitched                float64 `json:"inningsPitched"`
					HitsAllowed                   int     `json:"hitsAllowed"`
					SecondBaseHitsAllowed         int     `json:"secondBaseHitsAllowed"`
					ThirdBaseHitsAllowed          int     `json:"thirdBaseHitsAllowed"`
					RunsAllowed                   int     `json:"runsAllowed"`
					EarnedRunsAllowed             int     `json:"earnedRunsAllowed"`
					HomerunsAllowed               int     `json:"homerunsAllowed"`
					PitcherWalks                  int     `json:"pitcherWalks"`
					PitcherSwings                 int     `json:"pitcherSwings"`
					PitcherStrikes                int     `json:"pitcherStrikes"`
					PitcherStrikesFoul            int     `json:"pitcherStrikesFoul"`
					PitcherStrikesMiss            int     `json:"pitcherStrikesMiss"`
					PitcherStrikesLooking         int     `json:"pitcherStrikesLooking"`
					PitcherGroundBalls            int     `json:"pitcherGroundBalls"`
					PitcherFlyBalls               int     `json:"pitcherFlyBalls"`
					PitcherLineDrives             int     `json:"pitcherLineDrives"`
					PitcherSacrificeBunts         int     `json:"pitcherSacrificeBunts"`
					Pitcher2SeamFastballs         int     `json:"pitcher2SeamFastballs"`
					Pitcher4SeamFastballs         int     `json:"pitcher4SeamFastballs"`
					PitcherCurveballs             int     `json:"pitcherCurveballs"`
					PitcherChangeups              int     `json:"pitcherChangeups"`
					PitcherCutters                int     `json:"pitcherCutters"`
					PitcherSliders                int     `json:"pitcherSliders"`
					PitcherSinkers                int     `json:"pitcherSinkers"`
					PitcherSplitters              int     `json:"pitcherSplitters"`
					PitcherSacrificeFlies         int     `json:"pitcherSacrificeFlies"`
					PitcherStrikeouts             int     `json:"pitcherStrikeouts"`
					PitchingAvg                   float64 `json:"pitchingAvg"`
					WalksAndHitsPerInningPitched  float64 `json:"walksAndHitsPerInningPitched"`
					Shutouts                      int     `json:"shutouts"`
					BattersHit                    int     `json:"battersHit"`
					PitcherIntentionalWalks       int     `json:"pitcherIntentionalWalks"`
					PitcherGroundOuts             int     `json:"pitcherGroundOuts"`
					PitcherFlyOuts                int     `json:"pitcherFlyOuts"`
					PitcherWildPitches            int     `json:"pitcherWildPitches"`
					Balks                         int     `json:"balks"`
					PitcherStolenBasesAllowed     int     `json:"pitcherStolenBasesAllowed"`
					PitcherCaughtStealing         int     `json:"pitcherCaughtStealing"`
					Pickoffs                      int     `json:"pickoffs"`
					PickoffAttempts               int     `json:"pickoffAttempts"`
					TotalBattersFaced             int     `json:"totalBattersFaced"`
					PitchesThrown                 int     `json:"pitchesThrown"`
					PitcherGroundOutToFlyOutRatio float64 `json:"pitcherGroundOutToFlyOutRatio"`
					PitcherOnBasePct              float64 `json:"pitcherOnBasePct"`
					PitcherSluggingPct            float64 `json:"pitcherSluggingPct"`
					PitcherOnBasePlusSluggingPct  float64 `json:"pitcherOnBasePlusSluggingPct"`
					StrikeoutsPer9Innings         float64 `json:"strikeoutsPer9Innings"`
					WalksAllowedPer9Innings       float64 `json:"walksAllowedPer9Innings"`
					HitsAllowedPer9Innings        float64 `json:"hitsAllowedPer9Innings"`
					StrikeoutsToWalksRatio        float64 `json:"strikeoutsToWalksRatio"`
					PitchesPerInning              float64 `json:"pitchesPerInning"`
				} `json:"pitching"`
				Fielding struct {
					InningsPlayed             float64 `json:"inningsPlayed"`
					TotalChances              int     `json:"totalChances"`
					FielderTagOuts            int     `json:"fielderTagOuts"`
					FielderForceOuts          int     `json:"fielderForceOuts"`
					FielderPutOuts            int     `json:"fielderPutOuts"`
					Assists                   int     `json:"assists"`
					Errors                    int     `json:"errors"`
					FielderDoublePlays        int     `json:"fielderDoublePlays"`
					FielderTriplePlays        int     `json:"fielderTriplePlays"`
					FielderStolenBasesAllowed int     `json:"fielderStolenBasesAllowed"`
					FielderCaughtStealing     int     `json:"fielderCaughtStealing"`
					FielderStolenBasePct      float64 `json:"fielderStolenBasePct"`
					PassedBalls               int     `json:"passedBalls"`
					FielderWildPitches        int     `json:"fielderWildPitches"`
					FieldingPct               float64 `json:"fieldingPct"`
					DefenceEfficiencyRatio    float64 `json:"defenceEfficiencyRatio"`
					OutsFaced                 int     `json:"outsFaced"`
				} `json:"fielding"`
				Standings struct {
					Wins            int     `json:"wins"`
					Losses          int     `json:"losses"`
					WinPct          float64 `json:"winPct"`
					GamesBack       float64 `json:"gamesBack"`
					RunsFor         int     `json:"runsFor"`
					RunsAgainst     int     `json:"runsAgainst"`
					RunDifferential int     `json:"runDifferential"`
				} `json:"standings"`
			} `json:"teamStats"`
			Players []struct {
				Player struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"player"`
				PlayerStats []struct {
					Batting struct {
						AtBats                       int     `json:"atBats"`
						Runs                         int     `json:"runs"`
						Hits                         int     `json:"hits"`
						SecondBaseHits               int     `json:"secondBaseHits"`
						ThirdBaseHits                int     `json:"thirdBaseHits"`
						Homeruns                     int     `json:"homeruns"`
						EarnedRuns                   int     `json:"earnedRuns"`
						UnearnedRuns                 int     `json:"unearnedRuns"`
						RunsBattedIn                 int     `json:"runsBattedIn"`
						BatterWalks                  int     `json:"batterWalks"`
						BatterSwings                 int     `json:"batterSwings"`
						BatterStrikes                int     `json:"batterStrikes"`
						BatterStrikesFoul            int     `json:"batterStrikesFoul"`
						BatterStrikesMiss            int     `json:"batterStrikesMiss"`
						BatterStrikesLooking         int     `json:"batterStrikesLooking"`
						BatterTagOuts                int     `json:"batterTagOuts"`
						BatterForceOuts              int     `json:"batterForceOuts"`
						BatterPutOuts                int     `json:"batterPutOuts"`
						BatterGroundBalls            int     `json:"batterGroundBalls"`
						BatterFlyBalls               int     `json:"batterFlyBalls"`
						BatterLineDrives             int     `json:"batterLineDrives"`
						Batter2SeamFastballs         int     `json:"batter2SeamFastballs"`
						Batter4SeamFastballs         int     `json:"batter4SeamFastballs"`
						BatterCurveballs             int     `json:"batterCurveballs"`
						BatterChangeups              int     `json:"batterChangeups"`
						BatterCutters                int     `json:"batterCutters"`
						BatterSliders                int     `json:"batterSliders"`
						BatterSinkers                int     `json:"batterSinkers"`
						BatterSplitters              int     `json:"batterSplitters"`
						BatterStrikeouts             int     `json:"batterStrikeouts"`
						StolenBases                  int     `json:"stolenBases"`
						CaughtBaseSteals             int     `json:"caughtBaseSteals"`
						BatterStolenBasePct          float64 `json:"batterStolenBasePct"`
						BattingAvg                   float64 `json:"battingAvg"`
						BatterOnBasePct              float64 `json:"batterOnBasePct"`
						BatterSluggingPct            float64 `json:"batterSluggingPct"`
						BatterOnBasePlusSluggingPct  float64 `json:"batterOnBasePlusSluggingPct"`
						BatterIntentionalWalks       int     `json:"batterIntentionalWalks"`
						HitByPitch                   int     `json:"hitByPitch"`
						BatterSacrificeBunts         int     `json:"batterSacrificeBunts"`
						BatterSacrificeFlies         int     `json:"batterSacrificeFlies"`
						TotalBases                   int     `json:"totalBases"`
						ExtraBaseHits                int     `json:"extraBaseHits"`
						BatterDoublePlays            int     `json:"batterDoublePlays"`
						BatterTriplePlays            int     `json:"batterTriplePlays"`
						BatterGroundOuts             int     `json:"batterGroundOuts"`
						BatterFlyOuts                int     `json:"batterFlyOuts"`
						BatterGroundOutToFlyOutRatio float64 `json:"batterGroundOutToFlyOutRatio"`
						PitchesFaced                 int     `json:"pitchesFaced"`
						PlateAppearances             int     `json:"plateAppearances"`
						LeftOnBase                   int     `json:"leftOnBase"`
					} `json:"batting"`
					Fielding struct {
						InningsPlayed             float64 `json:"inningsPlayed"`
						TotalChances              int     `json:"totalChances"`
						FielderTagOuts            int     `json:"fielderTagOuts"`
						FielderForceOuts          int     `json:"fielderForceOuts"`
						FielderPutOuts            int     `json:"fielderPutOuts"`
						OutsFaced                 int     `json:"outsFaced"`
						Assists                   int     `json:"assists"`
						Errors                    int     `json:"errors"`
						FielderDoublePlays        int     `json:"fielderDoublePlays"`
						FielderTriplePlays        int     `json:"fielderTriplePlays"`
						FielderStolenBasesAllowed int     `json:"fielderStolenBasesAllowed"`
						FielderCaughtStealing     int     `json:"fielderCaughtStealing"`
						FielderStolenBasePct      float64 `json:"fielderStolenBasePct"`
						PassedBalls               int     `json:"passedBalls"`
						FielderWildPitches        int     `json:"fielderWildPitches"`
						FieldingPct               float64 `json:"fieldingPct"`
						RangeFactor               float64 `json:"rangeFactor"`
					} `json:"fielding"`
					Miscellaneous struct {
						GamesStarted int `json:"gamesStarted"`
					} `json:"miscellaneous"`
				} `json:"playerStats"`
			} `json:"players"`
		} `json:"away"`
		Home struct {
			TeamStats []struct {
				Batting struct {
					AtBats                       int     `json:"atBats"`
					Runs                         int     `json:"runs"`
					Hits                         int     `json:"hits"`
					SecondBaseHits               int     `json:"secondBaseHits"`
					ThirdBaseHits                int     `json:"thirdBaseHits"`
					Homeruns                     int     `json:"homeruns"`
					RunsBattedIn                 int     `json:"runsBattedIn"`
					EarnedRuns                   int     `json:"earnedRuns"`
					UnearnedRuns                 int     `json:"unearnedRuns"`
					BatterWalks                  int     `json:"batterWalks"`
					BatterSwings                 int     `json:"batterSwings"`
					BatterStrikes                int     `json:"batterStrikes"`
					BatterStrikesFoul            int     `json:"batterStrikesFoul"`
					BatterStrikesMiss            int     `json:"batterStrikesMiss"`
					BatterStrikesLooking         int     `json:"batterStrikesLooking"`
					BatterGroundBalls            int     `json:"batterGroundBalls"`
					BatterFlyBalls               int     `json:"batterFlyBalls"`
					BatterLineDrives             int     `json:"batterLineDrives"`
					BatterStrikeouts             int     `json:"batterStrikeouts"`
					Batter2SeamFastballs         int     `json:"batter2SeamFastballs"`
					Batter4SeamFastballs         int     `json:"batter4SeamFastballs"`
					BatterCurveballs             int     `json:"batterCurveballs"`
					BatterChangeups              int     `json:"batterChangeups"`
					BatterCutters                int     `json:"batterCutters"`
					BatterSliders                int     `json:"batterSliders"`
					BatterSinkers                int     `json:"batterSinkers"`
					BatterSplitters              int     `json:"batterSplitters"`
					LeftOnBase                   int     `json:"leftOnBase"`
					OpponentsLeftOnBase          int     `json:"opponentsLeftOnBase"`
					StolenBases                  int     `json:"stolenBases"`
					CaughtBaseSteals             int     `json:"caughtBaseSteals"`
					BatterStolenBasePct          float64 `json:"batterStolenBasePct"`
					BattingAvg                   float64 `json:"battingAvg"`
					BatterOnBasePct              float64 `json:"batterOnBasePct"`
					BatterSluggingPct            float64 `json:"batterSluggingPct"`
					BatterOnBasePlusSluggingPct  float64 `json:"batterOnBasePlusSluggingPct"`
					BatterIntentionalWalks       int     `json:"batterIntentionalWalks"`
					HitByPitch                   int     `json:"hitByPitch"`
					BatterSacrificeBunts         int     `json:"batterSacrificeBunts"`
					BatterSacrificeFlies         int     `json:"batterSacrificeFlies"`
					TotalBases                   int     `json:"totalBases"`
					ExtraBaseHits                int     `json:"extraBaseHits"`
					BatterDoublePlays            int     `json:"batterDoublePlays"`
					BatterTriplePlays            int     `json:"batterTriplePlays"`
					BatterTagOuts                int     `json:"batterTagOuts"`
					BatterForceOuts              int     `json:"batterForceOuts"`
					BatterPutOuts                int     `json:"batterPutOuts"`
					BatterGroundOuts             int     `json:"batterGroundOuts"`
					BatterFlyOuts                int     `json:"batterFlyOuts"`
					BatterGroundOutToFlyOutRatio float64 `json:"batterGroundOutToFlyOutRatio"`
					PitchesFaced                 int     `json:"pitchesFaced"`
					PlateAppearances             int     `json:"plateAppearances"`
					OpponentAtBats               int     `json:"opponentAtBats"`
				} `json:"batting"`
				Pitching struct {
					EarnedRunAvg                  float64 `json:"earnedRunAvg"`
					InningsPitched                float64 `json:"inningsPitched"`
					HitsAllowed                   int     `json:"hitsAllowed"`
					SecondBaseHitsAllowed         int     `json:"secondBaseHitsAllowed"`
					ThirdBaseHitsAllowed          int     `json:"thirdBaseHitsAllowed"`
					RunsAllowed                   int     `json:"runsAllowed"`
					EarnedRunsAllowed             int     `json:"earnedRunsAllowed"`
					HomerunsAllowed               int     `json:"homerunsAllowed"`
					PitcherWalks                  int     `json:"pitcherWalks"`
					PitcherSwings                 int     `json:"pitcherSwings"`
					PitcherStrikes                int     `json:"pitcherStrikes"`
					PitcherStrikesFoul            int     `json:"pitcherStrikesFoul"`
					PitcherStrikesMiss            int     `json:"pitcherStrikesMiss"`
					PitcherStrikesLooking         int     `json:"pitcherStrikesLooking"`
					PitcherGroundBalls            int     `json:"pitcherGroundBalls"`
					PitcherFlyBalls               int     `json:"pitcherFlyBalls"`
					PitcherLineDrives             int     `json:"pitcherLineDrives"`
					PitcherSacrificeBunts         int     `json:"pitcherSacrificeBunts"`
					Pitcher2SeamFastballs         int     `json:"pitcher2SeamFastballs"`
					Pitcher4SeamFastballs         int     `json:"pitcher4SeamFastballs"`
					PitcherCurveballs             int     `json:"pitcherCurveballs"`
					PitcherChangeups              int     `json:"pitcherChangeups"`
					PitcherCutters                int     `json:"pitcherCutters"`
					PitcherSliders                int     `json:"pitcherSliders"`
					PitcherSinkers                int     `json:"pitcherSinkers"`
					PitcherSplitters              int     `json:"pitcherSplitters"`
					PitcherSacrificeFlies         int     `json:"pitcherSacrificeFlies"`
					PitcherStrikeouts             int     `json:"pitcherStrikeouts"`
					PitchingAvg                   float64 `json:"pitchingAvg"`
					WalksAndHitsPerInningPitched  float64 `json:"walksAndHitsPerInningPitched"`
					Shutouts                      int     `json:"shutouts"`
					BattersHit                    int     `json:"battersHit"`
					PitcherIntentionalWalks       int     `json:"pitcherIntentionalWalks"`
					PitcherGroundOuts             int     `json:"pitcherGroundOuts"`
					PitcherFlyOuts                int     `json:"pitcherFlyOuts"`
					PitcherWildPitches            int     `json:"pitcherWildPitches"`
					Balks                         int     `json:"balks"`
					PitcherStolenBasesAllowed     int     `json:"pitcherStolenBasesAllowed"`
					PitcherCaughtStealing         int     `json:"pitcherCaughtStealing"`
					Pickoffs                      int     `json:"pickoffs"`
					PickoffAttempts               int     `json:"pickoffAttempts"`
					TotalBattersFaced             int     `json:"totalBattersFaced"`
					PitchesThrown                 int     `json:"pitchesThrown"`
					PitcherGroundOutToFlyOutRatio float64 `json:"pitcherGroundOutToFlyOutRatio"`
					PitcherOnBasePct              float64 `json:"pitcherOnBasePct"`
					PitcherSluggingPct            float64 `json:"pitcherSluggingPct"`
					PitcherOnBasePlusSluggingPct  float64 `json:"pitcherOnBasePlusSluggingPct"`
					StrikeoutsPer9Innings         float64 `json:"strikeoutsPer9Innings"`
					WalksAllowedPer9Innings       float64 `json:"walksAllowedPer9Innings"`
					HitsAllowedPer9Innings        float64 `json:"hitsAllowedPer9Innings"`
					StrikeoutsToWalksRatio        float64 `json:"strikeoutsToWalksRatio"`
					PitchesPerInning              float64 `json:"pitchesPerInning"`
				} `json:"pitching"`
				Fielding struct {
					InningsPlayed             float64 `json:"inningsPlayed"`
					TotalChances              int     `json:"totalChances"`
					FielderTagOuts            int     `json:"fielderTagOuts"`
					FielderForceOuts          int     `json:"fielderForceOuts"`
					FielderPutOuts            int     `json:"fielderPutOuts"`
					Assists                   int     `json:"assists"`
					Errors                    int     `json:"errors"`
					FielderDoublePlays        int     `json:"fielderDoublePlays"`
					FielderTriplePlays        int     `json:"fielderTriplePlays"`
					FielderStolenBasesAllowed int     `json:"fielderStolenBasesAllowed"`
					FielderCaughtStealing     int     `json:"fielderCaughtStealing"`
					FielderStolenBasePct      float64 `json:"fielderStolenBasePct"`
					PassedBalls               int     `json:"passedBalls"`
					FielderWildPitches        int     `json:"fielderWildPitches"`
					FieldingPct               float64 `json:"fieldingPct"`
					DefenceEfficiencyRatio    float64 `json:"defenceEfficiencyRatio"`
					OutsFaced                 int     `json:"outsFaced"`
				} `json:"fielding"`
				Standings struct {
					Wins            int     `json:"wins"`
					Losses          int     `json:"losses"`
					WinPct          float64 `json:"winPct"`
					GamesBack       float64 `json:"gamesBack"`
					RunsFor         int     `json:"runsFor"`
					RunsAgainst     int     `json:"runsAgainst"`
					RunDifferential int     `json:"runDifferential"`
				} `json:"standings"`
			} `json:"teamStats"`
			Players []struct {
				Player struct {
					ID           int    `json:"id"`
					FirstName    string `json:"firstName"`
					LastName     string `json:"lastName"`
					Position     string `json:"position"`
					JerseyNumber int    `json:"jerseyNumber"`
				} `json:"player"`
				PlayerStats []struct {
					Batting struct {
						AtBats                       int     `json:"atBats"`
						Runs                         int     `json:"runs"`
						Hits                         int     `json:"hits"`
						SecondBaseHits               int     `json:"secondBaseHits"`
						ThirdBaseHits                int     `json:"thirdBaseHits"`
						Homeruns                     int     `json:"homeruns"`
						EarnedRuns                   int     `json:"earnedRuns"`
						UnearnedRuns                 int     `json:"unearnedRuns"`
						RunsBattedIn                 int     `json:"runsBattedIn"`
						BatterWalks                  int     `json:"batterWalks"`
						BatterSwings                 int     `json:"batterSwings"`
						BatterStrikes                int     `json:"batterStrikes"`
						BatterStrikesFoul            int     `json:"batterStrikesFoul"`
						BatterStrikesMiss            int     `json:"batterStrikesMiss"`
						BatterStrikesLooking         int     `json:"batterStrikesLooking"`
						BatterTagOuts                int     `json:"batterTagOuts"`
						BatterForceOuts              int     `json:"batterForceOuts"`
						BatterPutOuts                int     `json:"batterPutOuts"`
						BatterGroundBalls            int     `json:"batterGroundBalls"`
						BatterFlyBalls               int     `json:"batterFlyBalls"`
						BatterLineDrives             int     `json:"batterLineDrives"`
						Batter2SeamFastballs         int     `json:"batter2SeamFastballs"`
						Batter4SeamFastballs         int     `json:"batter4SeamFastballs"`
						BatterCurveballs             int     `json:"batterCurveballs"`
						BatterChangeups              int     `json:"batterChangeups"`
						BatterCutters                int     `json:"batterCutters"`
						BatterSliders                int     `json:"batterSliders"`
						BatterSinkers                int     `json:"batterSinkers"`
						BatterSplitters              int     `json:"batterSplitters"`
						BatterStrikeouts             int     `json:"batterStrikeouts"`
						StolenBases                  int     `json:"stolenBases"`
						CaughtBaseSteals             int     `json:"caughtBaseSteals"`
						BatterStolenBasePct          float64 `json:"batterStolenBasePct"`
						BattingAvg                   float64 `json:"battingAvg"`
						BatterOnBasePct              float64 `json:"batterOnBasePct"`
						BatterSluggingPct            float64 `json:"batterSluggingPct"`
						BatterOnBasePlusSluggingPct  float64 `json:"batterOnBasePlusSluggingPct"`
						BatterIntentionalWalks       int     `json:"batterIntentionalWalks"`
						HitByPitch                   int     `json:"hitByPitch"`
						BatterSacrificeBunts         int     `json:"batterSacrificeBunts"`
						BatterSacrificeFlies         int     `json:"batterSacrificeFlies"`
						TotalBases                   int     `json:"totalBases"`
						ExtraBaseHits                int     `json:"extraBaseHits"`
						BatterDoublePlays            int     `json:"batterDoublePlays"`
						BatterTriplePlays            int     `json:"batterTriplePlays"`
						BatterGroundOuts             int     `json:"batterGroundOuts"`
						BatterFlyOuts                int     `json:"batterFlyOuts"`
						BatterGroundOutToFlyOutRatio float64 `json:"batterGroundOutToFlyOutRatio"`
						PitchesFaced                 int     `json:"pitchesFaced"`
						PlateAppearances             int     `json:"plateAppearances"`
						LeftOnBase                   int     `json:"leftOnBase"`
					} `json:"batting"`
					Fielding struct {
						InningsPlayed             float64 `json:"inningsPlayed"`
						TotalChances              int     `json:"totalChances"`
						FielderTagOuts            int     `json:"fielderTagOuts"`
						FielderForceOuts          int     `json:"fielderForceOuts"`
						FielderPutOuts            int     `json:"fielderPutOuts"`
						OutsFaced                 int     `json:"outsFaced"`
						Assists                   int     `json:"assists"`
						Errors                    int     `json:"errors"`
						FielderDoublePlays        int     `json:"fielderDoublePlays"`
						FielderTriplePlays        int     `json:"fielderTriplePlays"`
						FielderStolenBasesAllowed int     `json:"fielderStolenBasesAllowed"`
						FielderCaughtStealing     int     `json:"fielderCaughtStealing"`
						FielderStolenBasePct      float64 `json:"fielderStolenBasePct"`
						PassedBalls               int     `json:"passedBalls"`
						FielderWildPitches        int     `json:"fielderWildPitches"`
						FieldingPct               float64 `json:"fieldingPct"`
						RangeFactor               float64 `json:"rangeFactor"`
					} `json:"fielding"`
					Miscellaneous struct {
						GamesStarted int `json:"gamesStarted"`
					} `json:"miscellaneous"`
				} `json:"playerStats"`
			} `json:"players"`
		} `json:"home"`
	} `json:"stats"`
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
			TeamColoursHex      []string `json:"teamColoursHex"`
			SocialMediaAccounts []struct {
				MediaType string `json:"mediaType"`
				Value     string `json:"value"`
			} `json:"socialMediaAccounts"`
			OfficialLogoImageSrc string `json:"officialLogoImageSrc"`
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
		PlayerReferences []struct {
			ID              int    `json:"id"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			PrimaryPosition string `json:"primaryPosition"`
			JerseyNumber    int    `json:"jerseyNumber"`
			CurrentTeam     struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"currentTeam"`
			CurrentRosterStatus string      `json:"currentRosterStatus"`
			CurrentInjury       interface{} `json:"currentInjury"`
			Height              string      `json:"height"`
			Weight              int         `json:"weight"`
			BirthDate           string      `json:"birthDate"`
			Age                 int         `json:"age"`
			BirthCity           string      `json:"birthCity"`
			BirthCountry        string      `json:"birthCountry"`
			Rookie              bool        `json:"rookie"`
			HighSchool          interface{} `json:"highSchool"`
			College             string      `json:"college"`
			Handedness          struct {
				Bats   string `json:"bats"`
				Throws string `json:"throws"`
			} `json:"handedness"`
			OfficialImageSrc    string `json:"officialImageSrc"`
			SocialMediaAccounts []struct {
				MediaType string `json:"mediaType"`
				Value     string `json:"value"`
			} `json:"socialMediaAccounts"`
		} `json:"playerReferences"`
		PlayerStatReferences []struct {
			Category     string `json:"category"`
			FullName     string `json:"fullName"`
			Description  string `json:"description"`
			Abbreviation string `json:"abbreviation"`
			Type         string `json:"type"`
		} `json:"playerStatReferences"`
		TeamStatReferences []struct {
			Category     string `json:"category"`
			FullName     string `json:"fullName"`
			Description  string `json:"description"`
			Abbreviation string `json:"abbreviation"`
			Type         string `json:"type"`
		} `json:"teamStatReferences"`
	} `json:"references"`
}

//*Retrieved from mysportsfeeds*//
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