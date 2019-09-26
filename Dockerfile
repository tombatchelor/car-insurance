FROM golang:1.13
RUN mkdir /car-insurance 
WORKDIR /car-insurance  
COPY . .  
RUN go build 
CMD ["/car-insurance/car-insurance"]