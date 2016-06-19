package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/erasche/gologme/types"
	util "github.com/erasche/gologme/util"
	"github.com/gorilla/mux"
)

var dateLayout = "2006-01-02"

// Events lists the last N recorded events
func RecentWindows(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now().Unix()
	t0 := t1 - (60 * 15)
	windowData := golog.ExportWindowLogsByRange(t0, t1)
	if len(windowData) > 30 {
		windowData = windowData[len(windowData)-30:]
	}

	js, err := json.MarshalIndent(windowData, "", "  ")
	if err != nil {
		// handle
	}
	w.Write(js)
}

// Events lists events for a given day
func Events(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var tm time.Time
	i, err := strconv.ParseInt(vars["date"], 10, 64)
	if err != nil {
		if strings.Count(vars["date"], "-") == 2 {
			tm, err = time.Parse(dateLayout, vars["date"])
			if err != nil {
				return
			}
		} else {
			return
		}
	} else {
		tm = time.Unix(i, 0)
	}

	eventData := golog.ExportEventsByDate(tm)
	js, err := json.MarshalIndent(eventData, "", "  ")
	if err != nil {
		// handle
	}
	w.Write(js)
}

type eventListJSON struct {
	FName string `json:"fname"`
	T0    int    `json:"t0"`
	T1    int    `json:"t1"`
}

// ExportList produces custom json struct required by ulogme UI
func ExportList(w http.ResponseWriter, r *http.Request) {
	minTime, maxTime := golog.RecordedDataRange()
	// add a day to maxTime
	maxTime.Add(time.Hour * 12)
	// Convert to u7am
	iminTime := util.Ulogme7amTime(minTime)
	imaxTime := util.Ulogme7amTime(maxTime)
	//{"fname": "/api/events/1448197200", "t0": 1448197200, "t1": 1448283600}
	eventList := make([]*eventListJSON, 0)
	var aDayInSeconds = int64((time.Hour * 24).Seconds())
	var i int64
	for i = iminTime; i < imaxTime; i += aDayInSeconds {
		eventList = append(eventList, &eventListJSON{
			FName: fmt.Sprintf("/api/events/%d", int(i)),
			T0:    int(i),
			T1:    int(i + aDayInSeconds),
		})
	}
	js, _ := json.MarshalIndent(eventList, "", "  ")
	w.Write(js)
}

func DataUpload(w http.ResponseWriter, r *http.Request) {
	//var logs
	decoder := json.NewDecoder(r.Body)
	logData := new(types.DataLogRequest)
	err := decoder.Decode(&logData)

	if err != nil {
		//log.Error(fmt.Sprintf("Error unmarshalling posted data %s", err))
		http.Error(w, "Invalid Route Data", http.StatusBadRequest)
		return
	}

	golog.Log(logData)
}

func AddNote(w http.ResponseWriter, r *http.Request) {

}

func AddBlog(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	logData := new(types.NoteBlogRequest)
	err := decoder.Decode(&logData)
	if err != nil {
		http.Error(w, "Invalid Route Data", http.StatusBadRequest)
		return
	}

	tm, err := time.Parse(dateLayout, logData.Date)
	if err != nil {
		http.Error(w, "Invalid Route Data (Date)", http.StatusBadRequest)
		return
	}

	err = golog.CreateBlog(logData.User, logData.ApiKey, tm, logData.Message)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(logData, "", "  ")
	if err != nil {
		// handle
	}
	w.Write(js)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
