from golang:1.14-alpine as builder

RUN adduser -D -g 'flexo' flexo
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w -extldflags '-static'" -o /opt/flexo

from scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /opt/flexo /opt/flexo

EXPOSE 8080
USER flexo
ENTRYPOINT ["/opt/flexo", "run"]
