#Pull a base image which gives all required tools and libraries
FROM golang:1.20-alpine

#Create a folder where the app code will be stored
WORKDIR /app

#Copy the source code from your Host machine to your container
COPY . .

#Compile the application code
RUN go build -o age-calculator Main.go

#Run the application
CMD ["./age-calculator"]
