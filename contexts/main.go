// In Go, the context package is used to store key-value pairs and manage request-scoped data. Contexts are typically passed through all layers of an application (such as transport, business, and data layers) to:

//     Control the lifecycle of processes (e.g., cancel operations if they take too long).
//     Limit the number of incoming requests.
//     Handle edge case failures more gracefully.

// There are two main types of contexts:

//     context.Background(): Used to initialize an empty root context. It is the top-level context and is generally used when there is no existing context.
//     context.TODO(): Used to create an empty context when you are unsure which context to use. It acts as a placeholder until the appropriate context is determined in the future.

// In addition to these, the context object serves as a supplemental container to store specific information throughout the request lifecycle.

package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "request-id", "12345")
	return ctx
}

func addNameInContext(ctx context.Context, name string) context.Context {
	ctx = context.WithValue(ctx, "request-name", name)
	return ctx
}
func retrieveId(ctx context.Context) string {
	rId := ctx.Value("request-id").(string)

	// If not found, it will return nil and a boolean false.
	return rId
}

func contextWithTimeOuts(ctx context.Context) {
	rName := ctx.Value("request-name").(string)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool..")
			fmt.Println("Name ", rName)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	deadline := time.Now().Add(5 * time.Second)

	fmt.Println("GO CONTEXT PRACTICES")

	ctx := context.Background()
	ctx = enrichContext(ctx)
	id := retrieveId(ctx)
	fmt.Println("Id ", id)

	// ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctxWithTimeout, cancel := context.WithDeadline(context.Background(), deadline)

	fmt.Println("Error before closing done channel: ", ctxWithTimeout.Err())
	defer cancel()
	// if we give the cancel without defer keyword. It will immediately cancel the context

	// cancel()

	ctxWithTimeout = addNameInContext(ctxWithTimeout, "Venukishore R")
	go contextWithTimeOuts(ctxWithTimeout)

	select {
	case <-ctxWithTimeout.Done():
		fmt.Println("Oh no, I have exceeded the timeline")
		// when done channel is closed context will return an error. Otherwise it won't
		fmt.Println("Error after closing done channel", ctxWithTimeout.Err())
	}
	time.Sleep(2 * time.Second)
}
