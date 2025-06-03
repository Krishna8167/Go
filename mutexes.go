package main

/*
mutexes in Go :

Mutexes allows us to lock access to data. this ensures that we can control which goroutines can access certain
data at which time.

Go's standard library provides a built-in implementation of a mutex with sync.Mutex type and its two methods:

Lock()
Unlock()

We can protect a block of code by surrounding it with a call to lock amd Unlock as shown
on the protected() method below.

It's good practice to structure the protected code within a function so that defer can be used to ensure that we never
forget to unlock the mutex.

func protected(){
mux.Lock()
defer mux.Unlock()
// the rest of the function is protected
// any other calls to mux.lock() will block
}

Maps are not thread-safe
maps are not safe for concurrent use! if you have multiple goroutines accessing the same map, and atleast one of them is
writing to the map, you must lock your maos with a mutex.

mutex is short for mutual exclusion, because a mutex excludes different threads(or goroutines) from accessing the same data
at the same time.

the principle problem that mutexes help us avoid is the concurrent read/write problem. this problem arises when 1 thread is writing
to a variable while another thread is reading from that same variable at the same time.
1 thread can be lock a mutex at once.
use mutexes to safely access a data structure concurrently.

RW Mutex :
the standard library also exposes a sync.RWMutex

in addition to these methods: Lock() and UNlock()
the sync.RWMutex also has these methods:
RLock()
RUnlock()

The sync.Mutex can help with performance if we have a read-intensive process. Many goroutines can safely read from the map
at the same time(multiple RLock() calls can happen simultaneously). However, only one goroutine can hold a Lock() and all
RLock()'s will also be excluded.

only 1 writer can access a RWMutex at once.
infinite readers can access a RWMutex at once.
no one use reader and writers in RWMutexes at the same time.

*/
