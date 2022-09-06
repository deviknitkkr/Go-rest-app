FROM scratch
WORKDIR /app
COPY main "main"
CMD ["/app/main"]
