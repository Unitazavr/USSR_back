FROM alpine
WORKDIR .
COPY  . .
CMD ["./counter_back"]

