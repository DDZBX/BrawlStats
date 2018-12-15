package controllers

import (
	"github.com/revel/revel"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type App struct {
	*revel.Controller
}

var API_KEY string = "6RWCS0DHQ0IWVDZNGKTNFAIH13W0U"
var DodoSteamId string = "76561197980126328"

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SearchMe() revel.Result {

	//request API using credentials
	//var res, err = http.Get("https://api.brawlhalla.com/search?steamid=" + id + "&api_key=" + API_KEY)
	var res, err = http.Get("https://api.brawlhalla.com/search?steamid=76561197980126328&api_key=6RWCS0DHQ0IWVDZNGKTNFAIH13W0U")

	if err != nil {
		//handle error, for example return error page
	}

	//close body at the end of the method
	defer res.Body.Close()
	//read body
	body, err := ioutil.ReadAll(res.Body)
	var json = string(body)

	//return response
	return c.Render(res, err, json)
}

func (c App) SearchBySteamId() revel.Result {

	//retrieve steamId parameter
	var steamId = c.Params.Get("steamId")

	//request API using credentials
	var res, err = http.Get("https://api.brawlhalla.com/search?steamid=" + steamId + "&api_key=" + API_KEY)

	if err != nil {
		//handle error, for example return error page
	}

	//close body at the end of the method
	defer res.Body.Close()
	//read body
	body, err := ioutil.ReadAll(res.Body)
	var json = string(body)

	//extract brawlhalla_id from json and use it in next request
	var bid = gjson.Get(json, "brawlhalla_id").String()

	//Second request for player stats
	var resStats, errStats = http.Get("https://api.brawlhalla.com/player/" + bid + "/stats&api_key=" + API_KEY)
	//close body at the end of the method
	defer resStats.Body.Close()
	//read body
	bodyStats, errStats := ioutil.ReadAll(resStats.Body)
	var jsonStats = string(bodyStats)

	//return response
	return c.Render(steamId, res, err, json, resStats, errStats, jsonStats)

}
