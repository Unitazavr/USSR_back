FROM alpine
WORKDIR .
COPY  . .
EXPOSE 8087
CMD ["./counter_back"]

