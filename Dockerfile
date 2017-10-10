FROM scratch

COPY /shortener /app/

CMD ["/app/shortener"]
