FROM alpine:latest

LABEL app="bidding"
LABEL maintainer="manigandan.jeff@gmail.com"
LABEL version="1.4.2"
LABEL description="Ad Request Auction System."

RUN mkdir -p /app && apk update && apk add --no-cache ca-certificates
WORKDIR /app
# This require the project to be built first before copying,
# else docker build will fail
COPY bidding /app/

ENV ENV=dev
ENV PORT=8080
ENV AUCTIONEER_HOST=http://localhost:8080

EXPOSE 8080
CMD /app/bidding
