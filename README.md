# task-tracker
A CLI tool to track and manage your tasks.

# Build
- `make bin` will generate binaries within the `artifacts` directory.
- `tt` is the binary with code implemented using just the go standard library.
- `cobra-tt` is the binary coded using cobra framework.

# Use
Following are the commands, which are same for both the binaries.
## Add task
```shell
./tt add --task "description of the task"
```

## List task
To list all the tasks.
```shell
./tt list
```

To list tasks filtered via status (todo, in-progress, done) 
```shell
./tt list --status <todo/in-progress/done>
```

## Update task
```shell
./tt update --id <task_id> --task "description of the task"
```

## Mark task
To mark a task as ToDo
```shell
./tt mark-todo --id <task_id>
```

To mark a task as In Progress
```shell
./tt mark-in-progress --id <task_id>
```

To mark a task as Done
```shell
./tt mark-done --id <task_id>
```
