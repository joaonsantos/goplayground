package ratelimit

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrMaxReq = errors.New("max requests exceeded")

// RateLimiter represents an entity with rate limiting functionality.
type RateLimiter interface {
	// Records a call made by the user and returns an error if the limit is met or exceeded.
	RateLimit(clientID string) error
}

// RateLimiterImpl provides an implementation of rate limiting functionality.
type RateLimiterImpl struct {
	// ...
	IntervalSeconds int64 // time interval in seconds
	MaxRequests     int   // maximum number of requests allowed in interval

	Clients map[string]*CallList // client's calls
}

func NewClients() map[string]*CallList {
	return make(map[string]*CallList)
}

// CallList is list of times that each call is made
type CallList struct {
	*list.List
}

// initClientCalls initiates a new list of client calls with a call in the current time
func initClientCalls() *CallList {
	l := list.New()

	return &CallList{l}
}

// addCall adds the time of a call to the back of the list
func (cl *CallList) addCall() {
	cl.PushBack(time.Now())
}

// removeOldCalls removes all the calls before the given time
func (cl *CallList) removeOldCalls(rangeStart time.Time) {

	for t := cl.Front(); t != nil; {
		if t.Value.(time.Time).After(rangeStart) {
			return
		}

		next := t.Next()
		cl.Remove(t)
		t = next
	}
}

func (r *RateLimiterImpl) RateLimit(clientID string) error {

	// initiate call list for client
	if r.Clients[clientID] == nil {
		r.Clients[clientID] = initClientCalls()
	}

	// rangeStart is the time of the start of the interval in which there can only be max requests
	rangeStart := time.Now().Add(-time.Duration(r.IntervalSeconds * int64(time.Second)))

	// remove all calls that happen before rangeStart
	r.Clients[clientID].removeOldCalls(rangeStart)

	// check how many requests there are on the interval
	if r.Clients[clientID].Len() >= r.MaxRequests {
		return ErrMaxReq
	}

	// add current call
	r.Clients[clientID].addCall()

	return nil
}

func (r *RateLimiterImpl) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("IntervalSeconds: %d\n", r.IntervalSeconds))
	builder.WriteString(fmt.Sprintf("MaxReq: %d\n", r.MaxRequests))
	builder.WriteString("ClientIDS:\n")
	for clientID, callList := range r.Clients {
		var calls []string
		for t := callList.Front(); t != nil; t = t.Next() {
			calls = append(calls, t.Value.(time.Time).String())
		}

		builder.WriteString(fmt.Sprintf("- %s: {%s}\n", clientID, strings.Join(calls, "\n")))
	}

	return builder.String()
}
