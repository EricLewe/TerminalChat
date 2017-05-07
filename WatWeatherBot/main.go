package WatWeatherBot

import (
    "encoding/json"
    "log"
    "fmt"
    "net"
    ipinfo "github.com/syohex/go-ipinfo"
    "strings"
    "regexp"
    "net/http"
    "io/ioutil"
)

//Auto generated struct from: https://mholt.github.io/json-to-go/
type WeatherMap struct {
    Cod string `json:"cod"`
    Message float64 `json:"message"`
    Cnt int `json:"cnt"`
    List []struct {
        Dt int `json:"dt"`
        Main struct {
            Temp float64 `json:"temp"`
            TempMin float64 `json:"temp_min"`
            TempMax float64 `json:"temp_max"`
            Pressure float64 `json:"pressure"`
            SeaLevel float64 `json:"sea_level"`
            GrndLevel float64 `json:"grnd_level"`
            Humidity int `json:"humidity"`
            TempKf float64 `json:"temp_kf"`
        } `json:"main"`
        Weather []struct {
            ID int `json:"id"`
            Main string `json:"main"`
            Description string `json:"description"`
            Icon string `json:"icon"`
        } `json:"weather"`
        Clouds struct {
            All int `json:"all"`
        } `json:"clouds"`
        Wind struct {
            Speed float64 `json:"speed"`
            Deg float64 `json:"deg"` } `json:"wind"`
        Rain struct { ThreeH float64 `json:"3h"`
        } `json:"rain"`
        Sys struct {
            Pod string `json:"pod"`
        } `json:"sys"`
        DtTxt string `json:"dt_txt"`
        Snow struct {
            ThreeH float64 `json:"3h"`
        } `json:"snow,omitempty"`
    } `json:"list"`
    City struct {
        ID int `json:"id"`
        Name string `json:"name"`
        Coord struct {
            Lat float64 `json:"lat"`
            Lon float64 `json:"lon"`
        } `json:"coord"`
        Country string `json:"country"`
    } `json:"city"`
}

func initWeatherMap(body []byte) (*WeatherMap, error) {
    var s = new(WeatherMap)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("failed to marshal json:", err)
    }
    return s, err
}

func GetCurrentWeather(ip string) (string, string) {
    wd := GetWeatherMap(ip)
    weather := wd.List[0].Weather[0]
    return weather.Main, weather.Description
}

func GetWeatherMap(ip string)(wd WeatherMap) {
    info := ipinfo.IPInfo(net.ParseIP(ip))
    var validID = regexp.MustCompile(`([-+]?[0-9]*\.?[0-9]+),([-+]?[0-9]*\.?[0-9]+)`)
    coords := []string{"0", "0"}
    if validID.MatchString(info.Location) {
        coords = strings.SplitN(info.Location, ",", -1)
    }
    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&APPID=2d744698c24036564a733d5c1ad358ca", coords[0], coords[1])
    log.Println(url)
    res, err := http.Get(url)//"http://api.openweathermap.org/data/2.5/forecast?id=524901&APPID=2d744698c24036564a733d5c1ad358ca")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("", res.StatusCode)
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }

    s, err := initWeatherMap([]byte(body))
    return *s
}

