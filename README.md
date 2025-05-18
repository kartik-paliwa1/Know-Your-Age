# Know-Your-Age
A fun little project to find out how old you are — and how to Dockerize a Go app while you're at it!

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

Got it! Here’s your **Setup and Usage** section written in that style, tailored for your Go age calculator project starting from Docker setup:

---

## Setup and Usage

### Running Locally

1. Clone the repository:

   ```bash
   git clone [https://github.com/yourusername/know-your-age.git](https://github.com/kartik-paliwa1/Know-Your-Age)
   cd Know-Your-Age
   ```

2. Run the Go program:

   ```bash
   go run Main.go
   ```

3. Follow the prompt to enter your birth year and see your calculated age.

---

### Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t Know-Your-Age .
   ```

2. Run the container interactively:

   ```bash
   docker run -it Know-Your-Age
   ```

3. When prompted, enter your birth year inside the container and see your age output.

More improvements and additions like logging, user input validation, or even deploying to cloud platforms can be explored in future iterations of this project.
