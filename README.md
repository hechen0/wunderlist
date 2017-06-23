#### Yet Another wunderlist api client for go.

use for [my personal website](https://hechen0.com).


## Usage

```go
    go get -u github.com/hechen0/wunderlist
```

in your project

```go
    import "github.com/hechen0/wunderlist"
```

#### init api client

```go
    client := Wunderlist.NewClient()
```

#### set auth token

the auth token is represented by a Auth struct in this package.

```go
    type Auth struct {
    	Token    string
    	ClientId string
    }
```

```go
    auth := &wunderlist.Auth{Token: "xxx", ClientId: "xxx"}
    client.SetAuth(auth)
```

#### get all the list current user have
```go
    list, err := client.Lists.All(context.Background())
    if err != nil {
        fmt.Printf("got error: %v", err)
    }
```

# FAQ
#### Why the context.Background() for ?

basically, the context package provide a union method to cancel a request,
whether a user cancel a request explicitly or you set a timeline, and the
request exceed the timeline etc, the resource that associated with that
canceled request will be released immediately，so you have more control
over your request. for more infomation: check [context](https://golang.org/pkg/context/)

#### How do i get the token ?

the wunderlist api support a simple oauth2 service, so you can register
your app at [https://developer.wunderlist.com](https://developer.wunderlist.com/),
then get the token use a oatuh2 library, or by hand :) .

#### why another wunderlist api for go ?

first, i just wanna collect all my todo data, so i write a simple get
request to retrieve all my todo data, but as time goes by, the simple get
request grow bigger and bigger cause i need more detail about my data, suck
as my file, my note, my reminder, so one day, i decide to write a full
fuctional wunderlist api client.

### LICENSE

only for personal use, for commercial use, please contact me at me#hechen0.com