package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	router.HandleFunc("/", homePage).Methods("GET")               //? link : 127.0.0.1:8080/
	router.HandleFunc("/getTask/{id}", getTask).Methods("GET")    //? link : 127.0.0.1:8080/getTask/1
	router.HandleFunc("/getAllTasks", getAllTasks).Methods("GET") //? link : 127.0.0.1:8080/getAllTasks
	router.HandleFunc("/createTask", createTask).Methods("POST")
	router.HandleFunc("/updateTask/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/deleteTask/{id}", deleteTask).Methods("DELETE")
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
	data := mux.Vars(r)
	flag := false
	for index, value := range tasks {
		if data["id"] == value.ID {
			json.NewEncoder(w).Encode(tasks[index])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
	}
}

//? get all tasks route
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All tasks") //? link : 127.0.0.1:8080/getAllTasks
	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(tasks)
}

//? create task route
func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
}

//? update task route
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
}

//? delete task route
func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
}

/////////////////////////////////////////////////////////
//? initialize tasks
func allTaskList() {
	task1 := Tasks{
		ID:         "1",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       "2020-01-01",
	}
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       "2020-01-02",
	}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	fmt.Println(tasks)
}

//? struct
type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
	Date       string `json:"date"`
}
