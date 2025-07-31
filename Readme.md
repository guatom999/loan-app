    # Loan App Workshop

This repository is a sample Loan Application system consisting of two main services:

- **API** (Go): Backend service for handling loan-related logic.
- **Web** (React + Tailwind CSS): Frontend user interface.

## Project Structure

```
api/        # Backend service (Go)
web/        # Frontend application (React)
k8s/        # Kubernetes manifests for deployment
```

## Tech Stack

- Go (API)
- React, Tailwind CSS (Web)
- Docker, Kubernetes

## Getting Started

### 1. Clone the repository

```sh
git clone <repo-url>
cd loan-app
```

### 2. Build & Run with Docker

**Backend**

```sh
cd api
docker build -t loan-api .
docker run --env-file .env -p 4000:4000 loan-api
```

**Frontend**

```sh
cd ../web
docker build -t loan-web .
docker run -p 3000:80 loan-web
```

### 3. Deploy with Kubernetes

```sh
make <option-for-k8s>
```

## Notes

- Make sure to set up the `.env` file for each service before running.
- For horizontal scaling, ensure that the metrics-server is installed in your Kubernetes cluster.

---

> Feel free to customize this README to better
