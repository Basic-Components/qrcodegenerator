FROM alpine:3.10
ADD bin/linux-amd64/qrcodegenerator /code/qrcodegenerator
WORKDIR /code
CMD ["./qrcodegenerator"]