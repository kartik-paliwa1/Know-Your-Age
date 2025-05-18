# Know-Your-Age
## A fun little project to find out how old you are — and how to Dockerize a Go app while you're at it!

This is a basic Go lang project that shows your age if you know your birth year. In this, you have to input your birth year, and it will take the current year automatically.
By using the formula Current Year - Birth Year, it shows your current age.
This project is mainly intended to learn how to containerize an application. It is not a Go-focused project but mainly Docker-focused instead.
A Dockerfile will be created to build a Docker image and run it inside a container.

Once containerized, the application can be run easily using Docker commands on any system that supports Docker, regardless of the host environment. This helps demonstrate the power of containerization and the portability it brings to software development.

The project will also serve as a base to practice Docker fundamentals like:

-> Writing a Dockerfile

-> Building a lightweight image for a Go application

-> Running the image as a container

-> Understanding how to manage input/output inside the container

Overall, this project provides a practical, beginner-friendly example of containerizing a simple application — making it a great starting point for learning Docker and DevOps workflows.


---

## Setup and Usage

### 📝 Step 1: Clone the repository:

   ```bash
   git clone https://github.com/kartik-paliwa1/Know-Your-Age
   cd Know-Your-Age
   ```

### 📝 Step 2: Create a `Dockerfile`:

Example `Dockerfile`:

```Dockerfile
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
```

---

### 🛠️ Step 3: Build the Docker Image

```bash
docker build -t age-calculator .
```

* `-t age-calculator`: Tags the image with a custom name
* `.`: Refers to the current directory as the Docker context

---

### ▶️ Step 4: Run the Docker Container


```bash
docker run -it age-calculator
```

When prompted, enter your birth year inside the container and see your age output.

---

Let me know if you'd like to add Docker volumes, environment variables, or build using a multi-stage setup!


3. When prompted, enter your birth year inside the container and see your age output.

More improvements and additions like logging, user input validation, or even deploying to cloud platforms can be explored in future iterations of this project.
