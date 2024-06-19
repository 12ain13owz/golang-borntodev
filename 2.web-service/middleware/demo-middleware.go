package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID         int     `json:"ID"`
	Name       string  `json:"Name"`
	Price      float64 `json:"Price"`
	Instructor string  `json:"Instructor"`
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

func getNextID() int {
	highestID := -1

	for _, course := range courseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}

	return highestID + 1
}

func findID(ID int) (*Course, int) {
	for i, course := range courseList {
		if course.ID == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	course, listItemIndex := findID(ID)

	if course == nil {
		http.Error(w, fmt.Sprintf("No course with id %d", ID), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.Write(courseJSON)

	case http.MethodPut:
		var updatedCourse Course
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(byteBody, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedCourse.ID != ID {
			w.WriteHeader((http.StatusBadRequest))
			return
		}

		course = &updatedCourse
		courseList[listItemIndex] = *course
		w.WriteHeader((http.StatusOK))
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
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

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler middle start")
		handler.ServeHTTP(w, r)
		fmt.Println("middle finised")
	})
}

func main() {
	courseItemHandler := http.HandlerFunc(courseHandler)
	courseListHandler := http.HandlerFunc(coursesHandler)

	http.Handle("/course/", middlewareHandler(courseItemHandler))
	http.Handle("/course", middlewareHandler(courseListHandler))
	http.ListenAndServe(":5000", nil)
}
