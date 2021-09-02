FROM golang AS build

WORKDIR /gameserver-cli

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /build/gameserver-cli -v ./

FROM scratch

COPY --from=build /build/gameserver-cli ./

ENTRYPOINT [ "./gameserver-cli" ]