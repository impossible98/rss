FROM golang as alpha

WORKDIR /src

ADD . .

RUN make build


FROM golang

COPY --from=alpha /src/dist/rss /app/rss

CMD ["/app/rss"]
