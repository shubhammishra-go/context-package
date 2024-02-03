# Context in Go

A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
context can also store key-value pairs..
It provides a powerful toolset for managing concurrent operations. It enables the propagation of cancellation signals, deadlines, and values across goroutines, ensuring that related operations can gracefully terminate when necessary. With context, you can create a hierarchy of goroutines and pass important information down the chain.

# Why context need ?

Context provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutines.

so it is often used in concurrent and networked applications to manage the lifecycle and cancellation of operations.


# When to use Context ?

Request IDs for goroutine and function calls that are a component of an HTTP request call.

Information retrieval or data recovery or data fetching errors when working with a database.

When you wish to stop an operation in the middle of it - A HTTP request should be ended since the client has disconnected.

When you wish to stop an operation before a specific time - for example, a cron that needs to be aborted in 5 minutes if it is not done.

Obtain additional information about the environment they're running in, and usually pass that information to the functions they also call.


# To Create Context

there exists two ways to create context

## context.Background() Context

Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. 
It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests. 
It is often used when no specific context is available or needed.


## context.TODO() Context


TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter). 


# to assign Deadlines to a context

there are some ways to assign deadlines or values to parent context.


## context.WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

WithDeadline , is used to specify a completion date for the context and stops automatically after that date has passed.

WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d. If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent. The returned [Context.Done] channel is closed when the deadline expires, when the returned cancel function is called, or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete. 


## context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)

WithCancel method returns a copy of the parent context along with a cancel function; invoking the cancel function releases resources connected with the context and should be called as soon as operations in the Context type are finally completed.

WithCancel returns a copy of parent with a new Done channel. The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete. 


## context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout, allows a program to continue where it might otherwise hang, giving the end user a better experience.
Timeouts/Deadlines can be used to protect against corrupt or failing dependencies. It accepts a brief period as a parameter, along with the parent context, and terminates the function if it runs beyond the timeout period.


## context.WithValue(parent Context, key, val any) Context

WithValue accepts a parent context and returns a context copy. As a result, rather than overwriting the value, it creates a new duplicate with a new key-value pair.
Therefore, you should only use WithValue with data that is within a specific request scope.

Passing function arguments or values that will be modified later can result in the formation of many context variables, which will significantly increase your memory usage.

WithValue returns a copy of parent in which the value associated with key is val.

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.

The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context. Users of WithValue should define their own types for keys. To avoid allocating when assigning to an interface{}, context keys often have concrete type struct{}. Alternatively, exported context key variables' static type should be a pointer or interface. 

Passing information like request ID, user authentication, or configuration settings to functions without changing their signatures.
Storing and retrieving context-specific data in middleware or middleware-like patterns.


## context.AfterFunc(ctx Context, f func()) (stop func() bool)

AfterFunc arranges to call f in its own goroutine after ctx is done (cancelled or timed out). If ctx is already done, AfterFunc calls f immediately in its own goroutine.

Multiple calls to AfterFunc on a context operate independently; one does not replace another.

Calling the returned stop function stops the association of ctx with f. It returns true if the call stopped f from being run. If stop returns false, either the context is done and f has been started in its own goroutine; or f was already stopped. The stop function does not wait for f to complete before returning. If the caller needs to know whether f is completed, it must coordinate with f explicitly.

If ctx has a "AfterFunc(func()) func() bool" method, AfterFunc will use it to schedule the call. 

# Checking a exciting context "ctx"

there are some ways to check wheather context finished its dead

## <-ctx.Done()
Done() that can be used to determine whether or not a context has ended.
This method is used to identify when a context has finished executing and context value to add values to new contexts and to retrieve them in other functions using the Value method.

## Context.Cause(c Context) error

Cause returns a non-nil error explaining why c was canceled. The first cancellation of c or one of its parents sets the cause. If that cancellation happened via a call to CancelCauseFunc(err), then Cause returns err. Otherwise Cause(c) returns the same value as c.Err(). Cause returns nil if c has not been canceled yet. 


# About CancelFunc() function

A CancelFunc tells an operation to abandon its work. A CancelFunc does not wait for the work to stop. A CancelFunc may be called by multiple goroutines simultaneously. After the first call, subsequent calls to a CancelFunc do nothing.


# References

```https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go```

```https://pkg.go.dev/context#Context```

```https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea```

```https://www.scaler.com/topics/golang/golang-context/```

```https://www.educative.io/answers/golang-context```