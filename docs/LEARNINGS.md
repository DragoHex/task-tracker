- [flags](https://pkg.go.dev/flag) package can be used for command line flag parsing.
- Resource used [digitalocean](https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go)


## Help command using standard library.
- `flag.Usage` is a variable of type `func(){}` and can be overriden to define custom help response.
- help flag can be declared as a bool var. And check if an address has been assigned to execute `flag.Usage()`

## Makefile
- Makefile are meant for file tarkgets. Running a `make install` mean to execute some recipe to create the file named `install`. So if a file already exists in the directory make would not execute the recipe.
- So, we `.PHONY: install` before the series of command to execute without expectation of a file.
- To run a command in a quite mode prefix it with `@`.
- If not prefixed with `@`, by default the recipe would be printed on the terminal, which can be hidden using `.SILENT:`. If not particular target set this would silent all the make commands.
- One more way is to add `$(V)SILENT:` in the Makefile. Then all the make commands would be executed in quite mode. And if we execute them with a env variable `V` (which we are using here), then all the commands without `@` prefix would be printed. Ex: `V=1 make bin`
- To continue after a command failure, prefix it with `-`.
- To avoid printing error for a command failure add `|| true` or `2> /dev/null` at the end of the particular command.`2> /dev/null` at the end of the particular command.
