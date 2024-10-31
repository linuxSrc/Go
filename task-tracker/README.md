# Task Tracker

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Installation

To use Task Tracker, you need to have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/linuxSrc/Go
cd task-tracker
```

Install the required dependencies:

```sh
go get -u github.com/spf13/cobra
```

## Usage

Build the application:

```sh
go build -o task-tracker
```

Run the application:

```sh
./task-tracker
```

### Commands

- `add [task]`: Add a new task.
- `update [id] [task]`: Update an existing task.
- `mark-in-progress [id]`: Mark a task as in-progress.
- `mark-done [id]`: Mark a task as done.
- `delete [id]`: Delete a task by its ID.
- `list [status]`: List all tasks. Optionally, filter by status (`done`, `in-progress`, `todo`).

### Examples

Add a new task:

```sh
./task-tracker add "Buy groceries"
```

Update an existing task:

```sh
./task-tracker update 1 "Buy groceries and cook dinner"
```

Mark a task as in-progress:

```sh
./task-tracker mark-in-progress 1
```

Mark a task as done:

```sh
./task-tracker mark-done 1
```

Delete a task:

```sh
./task-tracker delete 1
```

List all tasks:

```sh
./task-tracker list
```

List tasks by status:

```sh
./task-tracker list done
./task-tracker list in-progress
./task-tracker list todo
```

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) for the CLI framework.
