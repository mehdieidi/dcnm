FROM golang:1.17 as build

COPY . /src/dcnm

WORKDIR /src/dcnm

RUN go get github.com/gorilla/mux github.com/lib/pq

RUN CGO_ENABLED=0 GOOS=linux go build -o dcnm


FROM scratch as image

COPY --from=build /src/dcnm .

EXPOSE 8080

CMD ["/dcnm"]