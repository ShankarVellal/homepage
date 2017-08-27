FROM scratch
ADD ./go/src/main/main /
CMD ["/main"]
EXPOSE 443
