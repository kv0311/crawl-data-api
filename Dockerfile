FROM golang:alpine
COPY . . 
CMD ["./data-crawl"]