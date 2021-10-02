FROM golang as alpha

WORKDIR /src

ADD . .

RUN make build


FROM golang

ENV PORT=8000

COPY --from=alpha /src/dist/rss /app/rss

EXPOSE $PORT

CMD ["/app/rss"]
