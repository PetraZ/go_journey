This chapter is about to show you two easy tools/libs you could use when you want create a simple backend endpoint.

1, fx, dependency ingestion
https://godoc.org/go.uber.org/fx
why bother?]
many times you'll need to import library from here to there, and code could get messy, dependecy ingestion basically help you centralize the needed librarys and output it when needed.

2, gin, web framwork
there will be always optimization around framework, gin is a such framwork that build faster web reponse, presumablly.

bonus:
Why you choose/like golang?
After using golang for a while, mainly cause of easy use of concurrency programming go routine and channels, which are frequenct used in backend development and very powerful.




How to play:

go run main.go
curl localhost:8100/ping

result:
pong