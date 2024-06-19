package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var courseList []Course

func init() {
	CourseJSON := `[
    {
      "id": 101,
      "name": "Python",
      "price": 2590,
      "instructor": "BorntoDev"
    },
    {
      "id": 102,
      "name": "JavaScript",
      "price": 0,
      "instructor": "BorntoDev"
    },
    {
      "id": 103,
      "name": "SQL",
      "price": 200,
      "instructor": "BorntoDev"
    }
  ]`

	err := json.Unmarshal([]byte(CourseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)

	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)

	case http.MethodPost:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var newCourse Course
		bodyByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyByte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getNextID() int {
	highestID := -1

	for _, course := range courseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}

	return highestID + 1
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
