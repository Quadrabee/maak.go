FROM --platform=${BUILDPLATFORM} golang:1.14.3-alpine AS build

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
ARG TARGETOS
ARG TARGETARCH

RUN go get -u github.com/jteeuwen/go-bindata/...

COPY . .

RUN go-bindata -o make/bindata.go -pkg make make/templates
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/maak .

FROM scratch AS bin-unix
COPY --from=build /out/maak /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/maak /maak.exe

FROM bin-${TARGETOS} AS bin
