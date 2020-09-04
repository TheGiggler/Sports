package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sports/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var updateCount int

var authHeader = "ODA5MWU4OGUtZDQ2Ni00YTdlLTljNTUtZTE2MTZhOk1ZU1BPUlRTRkVFRFM="

//test http request
func ProcessBoxScore(game models.GameSummary, kafkaChan chan interface{}, quitChannel chan bool, wg *sync.WaitGroup) {

	currentGameSummary := game
	refreshTimer := time.Tick(10000 * time.Millisecond)
	defer wg.Done()
	id := strconv.Itoa(game.GameID)
	tmp := "http://api.mysportsfeeds.com/v2.1/pull/mlb/2020-regular/games/{game}/boxscore.json"

	uri := strings.Replace(tmp, "{game}", id, -1)
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Basic "+authHeader)
	for {
		select {

		case <-refreshTimer:

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			if resp.StatusCode == 204 {
				msg := "Game " + id + " has not started"
				log.Println(string(msg))
				kafkaChan <- string(msg)
				return

			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			boxScore := &models.Boxscore{}
			newerr := json.Unmarshal([]byte(body), boxScore)
			if newerr != nil {
				log.Fatal(newerr)
			}
			//bodyBytes := []byte(body)
			//boxJson := string(bodyBytes[:])
			//fmt.Printf("myString:%v\n", boxJson)

			newGameSummary := GetGameSummaryFromBoxScore(boxScore)
			//todo: use a hash
			if newGameSummary.HomeTeamScore > currentGameSummary.HomeTeamScore || newGameSummary.AwayTeamScore > currentGameSummary.AwayTeamScore || newGameSummary.TimePeriod > currentGameSummary.TimePeriod {

				currentGameSummary = newGameSummary
				var gameStatus = "Inning: " + boxScore.Scoring.CurrentInningHalf + " " + strconv.Itoa(boxScore.Scoring.CurrentInning)

				if boxScore.GameSummary.PlayedStatus == "COMPLETED_PENDING_REVIEW" || boxScore.GameSummary.PlayedStatus == "COMPLETED" {
					gameStatus = "FINAL"
				}
				gameDesc := "GameID: " + id + " " + boxScore.GameSummary.AwayTeam.Abbreviation +
					" " + strconv.Itoa(boxScore.Scoring.AwayScoreTotal) + " @ " + boxScore.GameSummary.HomeTeam.Abbreviation + " " + strconv.Itoa(boxScore.Scoring.HomeScoreTotal) + "  " + gameStatus
				log.Println("Here's a big update!")
				log.Println(string(gameDesc))
			}
		case <-quitChannel:
			log.Printf("Quitting Game " + id)
			return
		default:

			//log.Println(string("Doing nothing for GameID " + id))
		}

	}
}

func GetGameSummaryFromBoxScore(boxScore *models.Boxscore) models.GameSummary {

	var gameSummary *models.GameSummary
	gameSummary = new(models.GameSummary)
	gameSummary.HomeTeamScore = boxScore.Scoring.HomeScoreTotal
	gameSummary.AwayTeamScore = boxScore.Scoring.AwayScoreTotal
	gameSummary.TimePeriod = boxScore.Scoring.CurrentInning
	return *gameSummary

}

func GetGamesForToday() []models.GameSummary {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//can do this, but not here!
	//	year, month, day := time.Now().Date()

	//get current date//
	current := time.Now()
	year, month, day := time.Now().Date()

	// var year int = current.Year()
	//var month int = int(current.Month())
	//var day int = current.Day()

	//var year = "2019"
	//var month = "7"
	//var day = "15"

	//if len(month) == 1 {
	//	month = "0" + month
	//}

	//if len(day) == 1 {
	//	day = "0" + day

	//}

	gameDate := time.Date(year, current.Month(), day, 0, 0, 0, 0, current.Location())
	fmt.Printf("Game Date %v\n", gameDate)

	//try to get from mongo first

	gamesFromDb, ok := GetGamesFromDb(mongoClient, gameDate)
	fmt.Printf("Games found:%v\n", gamesFromDb)

	if ok {
		return *gamesFromDb
	}

	uriDateString := GetUriDateString(year, int(month), day)
	//games not in db ... get from service and save to db
	//tmp := "https://api.mysportsfeeds.com/v2.1/pull/nhl/2018-2019/date/{gameDate}/games.json"
	tmp := "https://api.mysportsfeeds.com/v2.1/pull/mlb/2020-regular/date/{gameDate}/games.json"
	uri := strings.Replace(tmp, "{gameDate}", uriDateString, -1)
	//uri:="http://api.mysportsfeeds.com/v2.1/pull/mlb/2020-regular/date/20200821/games.json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Basic "+authHeader)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	gameFeed := &models.GameFeed{}
	newerr := json.Unmarshal([]byte(body), gameFeed)
	if newerr != nil {
		log.Fatal(newerr)
	}
	bodyBytes := []byte(body)
	myString := string(bodyBytes[:])
	fmt.Printf("myString:%v\n", myString)

	//	newerr := json.Unmarshal([]byte(body), &gameFeed)
	//	if newerr != nil {
	//	log.Fatal(newerr)
	//	}

	gameFeed.GameDate = gameDate
	//persist to mongo
	collection := mongoClient.Database("schedule").Collection("games")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, gameFeed)

	fmt.Printf("inserted:%v\n", res.InsertedID)

	games := FeedGamesToGameSummaries(gameFeed)
	return games
}

func GetUriDateString(year int, month int, day int) (dateString string) {

	yearStr := strconv.Itoa(year)
	monthStr := strconv.Itoa(month)
	dayStr := strconv.Itoa(day)

	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}

	if len(dayStr) == 1 {
		dayStr = "0" + dayStr

	}

	return yearStr + monthStr + dayStr
}

func GetGamesFromDb(client *mongo.Client, gameDate time.Time) (games *[]models.GameSummary, ok bool) {

	//func (coll *Collection) FindOne(ctx context.Context, filter interface{},
	// opts ...*options.FindOneOptions) *SingleResult

	collection := client.Database("schedule").Collection("games")
	dbGames := models.GameFeed{}

	gameDoc := bson.D{{"gamedate", gameDate}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	dbErr := collection.FindOne(ctx, gameDoc).Decode(&dbGames)

	if len(dbGames.Games) == 0 || dbErr != nil {
		return nil, false
	}

	games2 := FeedGamesToGameSummaries(&dbGames)

	return &games2, true
}

func FeedGamesToGameSummaries(feed *models.GameFeed) (gameModels []models.GameSummary) {
	games := []models.GameSummary{}
	for _, feedGame := range feed.Games {
		g := models.GameSummary{GameID: feedGame.Schedule.ID, HomeTeamID: feedGame.Schedule.HomeTeam.ID, HomeTeamName: feedGame.Schedule.HomeTeam.Abbreviation,
			AwayTeamID: feedGame.Schedule.AwayTeam.ID, AwayTeamName: feedGame.Schedule.AwayTeam.Abbreviation, HomeTeamScore: 0, AwayTeamScore: 0, TimePeriod: 0}
		games = append(games, g)

	}

	return games
}

func kafkaTest() {

}

func main() {

	//	updateChan := make(chan string)
	kafkaChan := make(chan interface{})

	quit := make(chan bool, 2)

	fmt.Printf("Welcome to Sports!\n")

	var wg sync.WaitGroup
	//var mux = sync.Mutex{}
	var games = GetGamesForToday()
	//need to get game list first ... needs to run synchronously?
	gameMap := GetGamesMap(games)
	fmt.Printf(gameMap[0].AwayTeamName)

	wg.Add(2)
	go PublishToKafka(kafkaChan, quit, &wg) //pass pointers to Channels?
	go ProcessGames(gameMap, quit, kafkaChan, &wg)
	//go ProcessGameData(updateChan, quit, &wg)

	wg.Wait()

	//for true {

	fmt.Printf("Sending quit\n")
	quit <- true
	quit <- true
	fmt.Print("Quits sent\n")
	time.Sleep(time.Second * 5)
	fmt.Printf("Sports is over!\n")
}

func ProcessGames(games map[int]models.GameSummary, quitChannel chan bool, kafkaChannel chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	//for true {
	for _, game := range games {

		wg.Add(1)
		go ProcessBoxScore(game, kafkaChannel, quitChannel, wg)
	}
	//}

}

func GetGamesMap(games []models.GameSummary) map[int]models.GameSummary {
	var gamesMap = make(map[int]models.GameSummary)

	for _, game := range games {
		gamesMap[game.GameID] = game
	}
	return gamesMap
}

func PublishToKafka(kafkaChan chan interface{}, quit chan bool, wg *sync.WaitGroup) {

	//containerAddress := "172.17.0.2:2181"

	for true {
		select {
		case newEvent := <-kafkaChan:
			//we would publish to kafka here
			fmt.Println("Published to KAFKA %v/n", newEvent)

		case <-quit:
			fmt.Println("Quitting PublishToKafka")
			wg.Done()
			return

		default:
			//fmt.Println("no KAFKA activity")
			return
		}
	}
}

//*Analyze game data for changes to scores or game time.  Publish changes to Kafka*//
func ProcessGameData(ch chan string, quit chan bool, wg *sync.WaitGroup) {

	//defer wg.Done()
	timer := time.NewTimer(time.Second * 6)
	for true {
		select {
		case newEvent := <-ch:
			fmt.Println("received", newEvent)
		case <-timer.C:
			fmt.Sprint("Timer expired")
		case <-quit:
			fmt.Println("Quitting OutputGameResult")
			//wg.Done()
			return

		default:
			//fmt.Println("no activity")
		}
	}

}

func GetLivePlayByPlay(gameID int) models.PlayByPlay {

	return models.PlayByPlay{}

}
