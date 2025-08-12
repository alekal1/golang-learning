## Concurrent Web Status Checker
Tool that checks the status (HTTP code or availability) of multiple websites concurrently using goroutines and channels.

## Technical Requirements

1. Use a goroutine for each HTTP request.

2. Use a chan to collect results.

3. Limit max concurrent requests (using buffered channels or semaphores).

4. Use sync.WaitGroup to wait for all routines to finish.

## Stretch Goals

1. Periodically check status every 5s (time.Ticker).


## Wanted output

```
[200] https://google.com

[200] https://github.com

[ERR] https://nonexistent.broken - Get "https://nonexistent.broken": no such host

[200] https://golang.org
```
