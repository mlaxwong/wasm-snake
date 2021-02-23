FROM golang:1.14.3-alpine 

WORKDIR /app
COPY . . 
RUN apk add --no-cache git
RUN go get -u github.com/shurcooL/goexec

WORKDIR /app/main
RUN GOOS=js GOARCH=wasm go build -o ../static/main.wasm

WORKDIR /app/static
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
CMD ["goexec", "http.ListenAndServe(`:80`, http.FileServer(http.Dir(`.`)))"]