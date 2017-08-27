FROM debian
ADD ./go/src/main/main /
CMD ["/main"]
EXPOSE 443
