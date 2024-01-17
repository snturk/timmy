# Timmy

Timmy is a simple time management utility for the command line.

For now, it only supports adding entries, starting and stopping them. It will also calculate the total time spent on each entry.

In the future, I plan to add more features like editing entries, deleting entries, and fetching entries from a remote server like Toggl.

## Installation

TODO

## Usage


### Define environment variables

Timmy uses environment variables to connect and sync to toggl and place time entries. You can define these variables in your `.bashrc` or `.zshrc` file.

```bash
# Timmy environment variables
# Optional, by default it will use $HOME/.timmy-entries
export TIMMY_PATH="path/to/.timmy-entries"
export TOGGL_WORKSPACE_ID="WORKSPACE_ID"
export TOGGL_API_TOKEN="API_TOKEN"
```
### Commands

`start` starts a new time entry with a task name `--task`.

```bash
timmy start --task "My task"
```

`stop` stops the current time entry.

```bash
timmy stop
```

`current` shows the current time entry, if any.

```bash
timmy current
```

`today` shows all time entries for today.

```bash
timmy today
```

`sync` syncs non-synced time entries with toggl, if you defined `TOGGL_WORKSPACE_ID` and `TOGGL_API_TOKEN` environment variables.

```bash
timmy sync
```

Just build and run the executable. It will create a `.timmy-entries` directory in your home directory and store all entries there.

```bash
git clone https://github.com/snturk/timmy.git
cd timmy
go build
./timmy --help
```

## Contributing

Pull requests and issues are always welcome. 