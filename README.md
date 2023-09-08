# go_web_development

GO Web Development Labs and Projects

## Requirements

- [Go](https://golang.org/dl/) 1.16 or higher
- [Git](https://git-scm.com/downloads)

## Labs

### Lab 1

To execute Go binary file, you need to set the executable permission on the binary file first.

```bash
./lab1
```

Fo testing the server, you can use curl command.

```bash
curl -v localhost:8080/index.html
curl -v localhost:8080/about.html
```

Or you can use browser to access the server.

Finished code is in the `lab1` folder.

### Lab 2

To execute Go binary file, you need to set the executable permission on the binary file first.
The lab2 looks like the same as lab1, but it uses the `net/http` package to implement the server.
It's more simple than lab1.

```bash
./lab2
```

Fo testing the server, you can use curl command.

```bash
curl -v localhost:8080/index.html
curl -v localhost:8080/about.html
```

Or you can use browser to access the server.

Finished code is in the `lab2` folder.

### Lab 3

In the Lab3, we will use the `html/template` package to implement the server. It's is other comcept to implement the server.
The `html/template` package is used to generate HTML output safe against code injection. It provides the same interface as `text/template` package.

And we will use enviroment to define the cache for the template. It's a good way to improve the performance.

If you want execute without cache, change the enviroment variable `env` to `false` in the `main.go` file and recompile the binary file.
Recompile the binary file is necessary, because the enviroment variable is set when the binary file is compiled.
To recompile the binary file, you can use the following command.

```bash
go build .
```

Remenber that cache mode not reflect the change of the template file. You need to recompile the binary file to reflect the change.

```bash
./lab3
```

Fo testing the server, you can use curl command.

```bash
curl -v localhost:8080/index.html
curl -v localhost:8080/about.html
```

Or you can use browser to access the server.

Finished code is in the `lab3` folder.

### Lab 4

In the Lab4, we will use the `gorilla/mux` package to implement the server. It's is other comcept to implement the server.
The `gorilla/mux` package is a powerful URL router and dispatcher. It's a good way to implement the RESTful API.

```bash
./lab4
```

or

The parameters "-port" and "-env" are optional. The default port is 8080 and the default enviroment is "dev".
Cache mode is the same as Lab3.
Remenber that cache mode not reflect the change of the template file. You need to recompile the binary file to reflect the change.

```bash
./lab4 -port 3000 -env "prod"
```

Fo testing the server, you can use curl command.

```bash
curl -v localhost:8080
curl -v localhost:8080/about
curl -v localhost:8080/contact
```

Or you can use browser to access the server.

Finished code is in the `lab4` folder.
