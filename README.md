# review-chatbot
The Review Chatbot Application is designed to empower small and medium businesses (SMBs) to collect personal reviews from their customers. The application leverages Next.js for the frontend, Go for the backend, and PostgreSQL for data persistence. The chatbot is dynamically triggered upon customers' sentiment after buying a product.

## How to run
Build and deployment is managed by [Tilt](https://tilt.dev/), a tool that allows quick iteration in local environments.
A quick summary in their own words:
> Tilt automates all the steps from a code change to a new process: watching files, building container images, and bringing your environment up-to-date. Think `docker build && kubectl apply` or `docker-compose up`.
#### Getting set up 
First install [Docker-Desktop](https://www.docker.com/products/docker-desktop/) and enable the Kubernetes engine.

Then install Tilt
`curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash`

Clone + run:
```
git clone https://github.com/atharvapatil4/review-chatbot.git
cd review-chatbot
tilt up
```
Navigate to `http://localhost:10350/`, where you should see all the services coming up. After `next-js-frontend` is deployed, head to `http://localhost:3001/` to interact with the application.

## Architecture
![review-chatbot-design](https://github.com/atharvapatil4/review-chatbot/assets/46949208/d2bb336f-b1ae-4404-bebd-2bfd73b4c8da)


