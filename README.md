# golang-parse-worker
REST based worker/dispatch example in GoLang

#Application
 starts a Go Server. and listens on /work route for post requests
 
# Queues/Channels

* WorkerQueue  ( this communicates between Dispatch and Worker ) 
* WorkQueue  ( this is a workers assignments )
* Quit Channel ( all workers listen for a quit/stop call )

# Starting workers
 each work instantiates itself and adds itself as available to  the WorkerQueue.  
 It will then wait on a Go subrouting, essentially block itself, until the worker is selected from the pool of workers
 a task is then placed on its WorkQueue.
  
# Dispatching Work

The handler publishes the work onto the WorkQueue. The dispatch pulls a worker from the WorkerQueue and assigns the work to the
workers stack.