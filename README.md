# Pratice for website crawler

## Get Started

- clone and cd to the project
```bash
git clone https://github.com/The-fthe/go-crawler

cd ./go-crawler

```
- run the command
```go
// go run ./main.go www.exammple.com 5 10
go run ./main.go <www.example.com> <maxConcurrency> <maxPages>
```

## Feature
- stdout url-link found count report

## Idea for extending the project
- Make the script run on a timer and deploy it to a server. Have it email you every so often with a report.
- Add more robust error checking so that you can crawl larger sites without issues.
- Count external links, as well as internal links, and add them to the report
- Save the report as a CSV spreadsheet rather than printing it to the console
- Use a graphics library to create an image that shows the links between the pages as a graph visualization
- Make requests concurrently to speed up the crawling process
