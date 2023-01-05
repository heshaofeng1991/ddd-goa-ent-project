FROM golang:1.19-buster

WORKDIR /

COPY ./ddd-goa-ent-project /ddd-goa-ent-project
COPY ./resources /resources

EXPOSE 80

# USER nonroot:nonroot

ENTRYPOINT ["/ddd-goa-ent-project"]
