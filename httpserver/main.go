// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// )

// type key string // we create a key with a custom type to avoid collision

// const keyServerAddress key = "serverAddr" //server key

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	//If queries are present, return true and the query value
// 	hasFirst := r.URL.Query().Has("first")
// 	first := r.URL.Query().Get("first")
// 	second := r.URL.Query().Get("second")
// 	hasSecond := r.URL.Query().Has("second")

// 	//Read query body
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("error reading body", err)
// 	}

// 	//reading entire body passed with -X POST flag
// 	//and -d body flag along with body string
// 	io.WriteString(w, "Root Page")
// 	fmt.Printf("%s:Got Root. first(%t):%s, second(%t):%s\n body: %s \n",
// 		ctx.Value(keyServerAddress), hasFirst, first, hasSecond, second, body)
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Println("Got Hello", ctx.Value(keyServerAddress)) //passing key
// 	// io.WriteString(w, "Hello Page\n")

// 	myName := r.PostFormValue("myName")
// 	if myName == "" {
// 		w.Header().Set("x-missing-field", "myName")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	io.WriteString(w, fmt.Sprintf("\nHello, %s\n", myName))
// }

// func main() {
// 	mux := http.NewServeMux() //initalizing multiplexer

// 	// HandleFunc registers the handler function for the given pattern.
// 	mux.HandleFunc("/", getRoot)
// 	mux.HandleFunc("/hello", getHello)

// 	ctx, cancelCtx := context.WithCancel(context.Background()) //empty context, utilized on ln 41 & 50

// 	//custom server
// 	//listens to port :3000
// 	//makes use of a key
// 	serverOne := &http.Server{
// 		Addr:    ":3000",
// 		Handler: mux,
// 		BaseContext: func(l net.Listener) context.Context { //feed key to server, listening port as context
// 			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
// 			return ctx
// 		},
// 	}
// 	fmt.Println("Server running")

// 	err := serverOne.ListenAndServe()
// 	if errors.Is(err, http.ErrServerClosed) {
// 		fmt.Println("Server 1 Closed")
// 	} else if err != nil {
// 		fmt.Println("Unable to start server 1")
// 	}
// 	cancelCtx()

// }

// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// )

// type key string // we create a key with a custom type to avoid collision

// const keyServerAddress key = "serverAddr" //server key

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Println("Got Root", ctx.Value(keyServerAddress)) //passing key
// 	io.WriteString(w, "Root Page")
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Println("Got Hello", ctx.Value(keyServerAddress)) //passing key
// 	io.WriteString(w, "Hello Page")
// }

// func main() {
// 	mux := http.NewServeMux() //initalizing multiplexer

// 	// HandleFunc registers the handler function for the given pattern.
// 	mux.HandleFunc("/", getRoot)
// 	mux.HandleFunc("/hello", getHello)

// 	ctx, cancelCtx := context.WithCancel(context.Background()) //empty context, utilized on ln 41 & 50
// 	//custom server
// 	//listens to port :3000
// 	//makes use of a key
// 	serverOne := &http.Server{
// 		Addr:    ":3000",
// 		Handler: mux,
// 		BaseContext: func(l net.Listener) context.Context { //feed key to server, listening port as context
// 			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
// 			return ctx
// 		},
// 	}

// 	//custom server
// 	//listens to port :3030
// 	//makes use of a key
// 	serverTwo := &http.Server{
// 		Addr:    ":3030",
// 		Handler: mux,
// 		BaseContext: func(l net.Listener) context.Context { // to listen to a specific port
// 			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
// 			return ctx
// 		},
// 	}

// 	fmt.Println("Server running")

// 	// create go function to listen to multiple
// 	// servers concurrently to avoid collision
// 	go func() {
// 		err := serverOne.ListenAndServe()
// 		if errors.Is(err, http.ErrServerClosed) {
// 			fmt.Println("Server 1 Closed")
// 		} else if err != nil {
// 			fmt.Println("Unable to start server 1")
// 		}
// 		cancelCtx()
// 	}()

// 	// create go function to listen to multiple
// 	// servers concurrently to avoid collision
// 	go func() {
// 		err := serverTwo.ListenAndServe()
// 		if errors.Is(err, http.ErrServerClosed) {
// 			fmt.Println("Server 2 Closed")
// 		} else if err != nil {
// 			fmt.Println("Unable to start server 2")
// 		}
// 		cancelCtx()
// 	}()
// 	<-ctx.Done()
// }

package main

import (
	// "context"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	// "os"
	// "net"
	"net/http"
	"time"
)

type key int

const (
	ctxKey     key = 1    //context key to avoid collision
	serverport key = 8000 // server port
)

func main() {
	go func() { //creating go routines to ensure server requests do not collide
		mux := http.NewServeMux()                                          //multiplexing for custom server and/or running multiple servers at once
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //takes a pattern and a handle

			fmt.Printf("server: %s\n", r.Method)
			fmt.Printf("query id: %s\n", r.URL.Query().Get("id"))
			fmt.Printf("server content type: %s\n", r.Header.Get("content-type"))
			fmt.Printf("server headers: \n")

			for headerName, headerValue := range r.Header { // Displaying header name nad contents
				fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
			}

			// reqBody, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	fmt.Printf("Unable to read request body: %s\n", err)
			// }

			// fmt.Printf("request body: %s\n", reqBody)
			// fmt.Println("Request Method:", r.Method) //Method: GET

			io.WriteString(w, "{message: \"Hello\"}") //writing to the response writer/body
			time.Sleep(1 * time.Second)
		})

		// ctx, ctxCancel := context.WithCancel(context.Background())
		server := http.Server{ //creating custom server
			Addr:    fmt.Sprintf(":%d", serverport),
			Handler: mux,
			// BaseContext: func(l net.Listener) context.Context {
			// 	ctx = context.WithValue(ctx, ctxKey, l.Addr().String())
			// 	return ctx
			// },
		}

		//Listing to the above server and handling incoming requests
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Server Closed")
			}
		}
		// ctxCancel()
	}()
	time.Sleep(100 * time.Millisecond) // additional time for server startup

	jsonbody := []byte(`{"client_message": "hello, server!"}`) //writing to server in the form of pseudo json
	bodyReader := bytes.NewReader(jsonbody)

	requestURL := fmt.Sprintf("http://localhost:%d?id=1234", serverport)
	// res, err := http.Get(requestURL)

	//Creating new request by passing method, server port, and byte content to be displayed
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Println("Unable to make server request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json") //specifying type of content to be passed

	//-------------------START: Sending request using default client-------------------//
	// res, err := http.DefaultClient.Do(req) //sending request to the server
	// if err != nil {
	// 	fmt.Println("Unable to fetch response", err)
	// 	return
	// }

	// go func() {
	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintf(w, `{message:"Hey There!"}`)
	// 		time.Sleep(35 * time.Second)
	// 	})
	// }()
	//-------------------END: Sending request using default client-------------------//

	client := http.Client{ //Creating client with timeout
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req) // response by server to client
	if err != nil {
		fmt.Println("Unable to fetch response", err)
		return
	}

	fmt.Println("client: got response!")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read response body", err)
		return
	}
	fmt.Printf("client: response body: %s\n", resBody)

	
}
