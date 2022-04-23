FROM aristorinjuang/go-builder AS BUILDER
COPY . .
RUN go build -o app

FROM aristorinjuang/go-production
ENV TZ=Asia/Jakarta
COPY --from=BUILDER /app/app .
CMD ["./app"]