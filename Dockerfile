# Build Step
FROM golang:1.18-alpine AS builder

# Dependencies
RUN apk update && apk add --no-cache upx make git
COPY --from=mwader/static-ffmpeg:4.4.0 /ffmpeg /tmp/ffmpeg
RUN upx --best --lzma /tmp/ffmpeg

# Source
WORKDIR $GOPATH/src/github.com/apollgo/apollgo
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN make tmp
RUN upx --best --lzma /tmp/apollgo

# Last Step
FROM gcr.io/distroless/static
COPY --from=builder /tmp/apollgo /go/bin/apollgo

ENTRYPOINT ["/go/bin/apollgo"]