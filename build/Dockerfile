FROM node:lts as alpha

WORKDIR /src

ADD . .

RUN make build-frontend


FROM golang as beta

WORKDIR /src

COPY --from=alpha /src .

RUN make build


FROM golang

ENV PORT=8000

COPY --from=beta /src/dist/rss /app/rss

EXPOSE $PORT

CMD ["/app/rss"]
