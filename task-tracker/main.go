package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type TaskStore struct {
	Tasks map[string]string `json:"tasks"`
	LastID int 				 `json:"lastID"`
	MarkInProgress map[string]bool 	 `json:"mark"`
}


const dataFile = "tasks.json"

func loadTasks() TaskStore {
	var store TaskStore
	data, err := os.ReadFile(dataFile)

	if err != nil {
		return TaskStore{
			Tasks: make(map[string]string),
			LastID: 1,
			MarkInProgress: make(map[string]bool),
		}
	}

	err = json.Unmarshal(data, &store)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return TaskStore{
			Tasks: make(map[string]string),
			LastID: 1,
			MarkInProgress: make(map[string]bool),

		}
	}
	return store
}

func saveTasks (store TaskStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}



func main() {

	store := loadTasks()

	var rootCmd = &cobra.Command{
		Use: "task-tracker",
		Short: "A simple CLI for managing tasks",
	}

	var addCmd = &cobra.Command{
		Use: "add [task]",
		Short: "Add a new task",
		Run: func (cmd *cobra.Command, args []string)  {
			if len(args) < 1 {
				fmt.Println("Please provide a task to add.")
				return
			}
			task := args[0]
            strID := strconv.Itoa(store.LastID)
            store.Tasks[strID] = task
			store.MarkInProgress[strID] = false
			fmt.Printf("Task added successfully (ID: %d)", store.LastID)
			store.LastID++

			if err := saveTasks(store); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
		},
	}

	var updateCmd = &cobra.Command{
		Use: "update",
		Short: "Update the existing task",
		Run: func (cmd *cobra.Command, args []string)  {
            id := args[0]
            task := args[1]
            
            if _, exists := store.Tasks[id]; !exists {
                fmt.Printf("Error: task with ID %s not found\n", id)
                return
            }
            
            store.Tasks[id] = task
            fmt.Printf("Successfully updated the task with ID %s\n", id)
            
            if err := saveTasks(store); err != nil {
                fmt.Printf("Error saving tasks: %v\n", err)
            }
		},
	}

	var markCmd = &cobra.Command{
		Use: "mark-in-progress",
		Short: "mark the task in progress",
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			if _, exists := store.Tasks[id]; !exists {
				fmt.Printf("Error task with id %s not found\n", id)
			}
			store.MarkInProgress[id] = true
			if err := saveTasks(store); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
			fmt.Println("Successfully marked the task in progress.")
		},
	}

    var showCmd = &cobra.Command{
        Use:   "show",
        Short: "Show the list of available tasks with their ID",
        Run: func(cmd *cobra.Command, args []string) {
            if len(store.Tasks) == 0 {
                fmt.Println("No tasks available.")
                return
            }
            
            fmt.Println("Tasks:")
            for id, task := range store.Tasks {
                fmt.Printf("%s: %s\n", id, task)
            }
        },
    }

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(markCmd)

	if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}