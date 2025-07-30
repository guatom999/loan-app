deploy-backend:
	kubectl apply -f ./k8s/backend-configmap.yaml
	kubectl apply -f ./k8s/backend-deployment.yaml
	kubectl apply -f ./k8s/backend-hcpa.yaml
	kubectl apply -f ./k8s/backend-service.yaml

deploy-frontend:
	kubectl apply -f ./k8s/frontend-deployment.yaml
	kubectl apply -f ./k8s/frontend-service.yaml

deploy-delete-all:
	kubectl delete -f ./k8s/backend-deployment.yaml
	kubectl delete -f ./k8s/backend-service.yaml
	kubectl delete -f ./k8s/backend-configmap.yaml
	kubectl delete -f ./k8s/frontend-deployment.yaml
	kubectl delete -f ./k8s/frontend-service.yaml
	kubectl delete -f ./k8s/frontend-configmap.yaml

deploy-databases:
	kubectl apply -f ./k8s/postgres-deployment.yaml
	kubectl apply -f ./k8s/postgres-service.yaml
	kubectl apply -f ./k8s/postgres-configmap.yaml

.PHONY:  docker-build-backend
docker-build-backend:
	@echo "Building backend image..."
	docker build -t badzboss/loan-app-api:latest ./api && docker push badzboss/loan-app-api:latest
	

.PHONY: docker-build-frontend
docker-build-frontend:
	@echo "Building frontend image..."
	docker build -t badzboss/loan-app-web:latest ./web && docker push badzboss/loan-app-web:latest
