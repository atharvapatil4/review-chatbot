FROM python:3.9-slim

WORKDIR /usr/src/app
RUN apt-get update && apt-get upgrade -y
RUN apt-get install libpq-dev gcc -y
RUN pip install psycopg2

# Copy the script and the assets into the directory
COPY db/seed.py .
COPY ../marketplace/public/assets/images ./assets

# Define the command to run the script
CMD ["python", "./seed.py"]