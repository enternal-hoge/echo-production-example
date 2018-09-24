# echo-production-example
Example App using golang, echo, sqlx, and so on.


# Requiment
+ mysql or mariadb

# Note
+ This application use go 1.11.
+ don't use glide or dep, use go modules.

set environment variables.
```
GO111MODULE=on
```

on first build, down load dependency modules automatically.
```
$ cd echo-production-example
$ go build
```


# Running
```
$ cd echo-production-example
$ go run main.go
```

# Running［fresh］
[fresh](https://github.com/pilu/fresh)

if you want to use live reloading.
```
$ cd echo-production-example
$ fresh
```

# Structure
```
main -> route > facade -> dao
```

# HTTP Request Interceptor

http request interceptor before route 

common/intercepter.go
```
...

func TraceLogInterceptor() e.MiddlewareFunc {
	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {
			/* before processing route execution */
      
      //TODO : variable processes.
      
			err := next(c)
			/* after processing route execution */

			return err
		}
	}
...
```


# Lisence
MIT

# Author
Kenichi Murakata
