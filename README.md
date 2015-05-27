This is presentation code.

railsapp - code of rails application
goapp - code of go application

railsapp

Rails app for this presentation is very straightforward.

goapp

Go app has 2 executables:
* http server - simple server that queries rails database and returns json.
  It shows how to easily write a go program that may substitute rails endpoint
* resque/sidekiq compatibile worker - thanks to go-worker project, its easy to reduce memory
  and time overhead for more complicated tasks

