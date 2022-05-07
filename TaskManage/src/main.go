package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var tasks []Tasks

func main() {

	fmt.Println("started")
	allTaskList()
	//?
	handlRoutes()
}

//? create routes
func handlRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")                     //? link : 127.0.0.1:8080/
	router.HandleFunc("/getTask/{id}", getTask).Methods("GET")          //? link : 127.0.0.1:8080/getTask/1
	router.HandleFunc("/getAllTasks", getAllTasks).Methods("GET")       //? link : 127.0.0.1:8080/getAllTasks
	router.HandleFunc("/createTask", createTask).Methods("POST")        //? link : 127.0.0.1:8080/createTask
	router.HandleFunc("/updateTask/{id}", updateTask).Methods("PUT")    //? link : 127.0.0.1:8080/updateTask/1
	router.HandleFunc("/deleteTask/{id}", deleteTask).Methods("DELETE") //? link : 127.0.0.1:8080/deleteTask/1
	//?
	log.Fatal(http.ListenAndServe(":8080", router))
}

///? route methods

//? home route
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/
}

//? get task route
func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get task") //? link : 127.0.0.1:8080/getTask/1
	_params := mux.Vars(r)  //? get data from url like /1
	_flag := false
	for index, value := range tasks {
		if _params["id"] == value.ID {
			json.NewEncoder(w).Encode(tasks[index]) //? encode
			_flag = true
			break
		}
	}
	if _flag == false {
		errorMessage(w, r, "task not found")
	}
}

//? get all tasks route
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All tasks") //? link : 127.0.0.1:8080/getAllTasks
	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(tasks) //? encode
}

//? create task route
func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Header().Set("Content-Type", "application/json")
	var _task Tasks
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_taskLength := len(tasks)                  //? task length
	_task.ID = strconv.Itoa(_taskLength)       //? initialize a incremented id
	_task.Date = time.Now()
	if _task.TaskName == "" {
		errorMessage(w, r, "Task name can't be empty")
	} else if _task.TaskDetail == "" {
		errorMessage(w, r, "Task detail can't be empty")
	} else {
		tasks = append(tasks, _task) //? add new task

		if len(tasks) != _taskLength { //? that means, new task add in the slice
			sendMessage(w, r, "success")
		} else {
			sendMessage(w, r, "failed")
		}
	}
}

/*
//? update task route (1st way)
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/updateTask/1
	w.Header().Set("Content-Type", "application/json")
	_params := mux.Vars(r) //? get data from url like /1
	var _task Tasks
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_id, _ := strconv.Atoi(_params["id"])      //? id from paramiter
	_data := tasks[_id]

	if _task.TaskName != "" || _task.TaskDetail != "" {
		if _task.TaskName != "" {
			_data.TaskName = _task.TaskName
		}
		if _task.TaskDetail != "" {
			_data.TaskDetail = _task.TaskDetail
		}
		_data.Date = time.Now()
		tasks[_id] = _data
		sendMessage(w, r, "success")
	} else {
		sendMessage(w, r, "Task can't be empty")
	}
}
*/

//? update task route (2nd way)
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/updateTask/1
	w.Header().Set("Content-Type", "application/json")
	_params := mux.Vars(r)                //? get data from url like /1
	_id, _ := strconv.Atoi(_params["id"]) //? id from paramiter
	var _task Tasks
	_ = json.NewDecoder(r.Body).Decode(&_task) //? decode (r.Body return body)
	_data := &tasks[_id]

	if _task.TaskName != "" || _task.TaskDetail != "" {
		if _task.TaskName != "" {
			_data.TaskName = _task.TaskName
		}
		if _task.TaskDetail != "" {
			_data.TaskDetail = _task.TaskDetail
		}
		_data.Date = time.Now()
		tasks[_id] = *_data
		sendMessage(w, r, "success")
	} else {
		sendMessage(w, r, "Task can't be empty")
	}
}

//? delete task route
func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/deleteTask/1
	w.Header().Set("Content-Type", "application/json")
	_id := mux.Vars(r)["id"] //? get data from url like /1
	_flag := false
	for index, value := range tasks {
		if value.ID == _id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			_flag = true
			sendMessage(w, r, "success")
		}
	}
	if _flag == false {
		errorMessage(w, r, "Invalid id")
	}
}

///? send message
func sendMessage(w http.ResponseWriter, r *http.Request, message string) {
	json.NewEncoder(w).Encode(map[string]string{"message": message}) //? success message
}

//? error message
func errorMessage(w http.ResponseWriter, r *http.Request, message string) {
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

/////////////////////////////////////////////////////////
//? initialize tasks
func allTaskList() {
	task1 := Tasks{
		ID:         "0",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       time.Now(),
	}
	task2 := Tasks{
		ID:         "1",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       time.Now(),
	}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	fmt.Println(tasks)
}

//? struct
type Tasks struct {
	ID         string    `json:"id"`
	TaskName   string    `json:"taskName"`
	TaskDetail string    `json:"taskDetail"`
	Date       time.Time `json:"date"`
}
