# Crawler

A simple go application, given a website it goes through all of the websites links (it will not leave the site) and prints all the found links with there occurrence.

## How to use it.
The application accepts 3 arguments, the website URL, the number of concurrent `crawlers` you want and the max number of found link.

### example
`./crawler "https://examples.com" 3 50`
with the above example we are going to crawl the examples.com having as max concurrency to 3 and the crawler will stop once he finished the website or when it found 50 links.
