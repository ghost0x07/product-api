FROM golang:1.16 as build

WORKDIR /app
COPY . .

RUN make build

###

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/product-api /app

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/app/product-api" ]