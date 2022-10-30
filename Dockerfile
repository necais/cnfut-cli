FROM alpine:3.16.2
RUN apk --no-cache add ca-certificates
COPY skbn /usr/local/bin/cnfut
RUN addgroup -g 1001 -S cnfut \
    && adduser -u 1001 -D -S -G cnfut cnfut
USER cnfut
WORKDIR /home/cnfut
CMD ["cnfut"]