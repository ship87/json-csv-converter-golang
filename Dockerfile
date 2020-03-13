FROM alpine:3.11

ENV JSON_CSV_CONVERTER_GOLANG_DIRECTORY_DOWNLOAD=download
ENV JSON_CSV_CONVERTER_GOLANG_PREFIX_FILE=result-*.csv
ENV JSON_CSV_CONVERTER_GOLANG_APP_URL=http://localhost:8081
ENV JSON_CSV_CONVERTER_GOLANG_APP_PORT=8081

RUN apk add --no-cache ca-certificates

WORKDIR /public

COPY ./service /service

EXPOSE 8081

CMD ["./../service"]