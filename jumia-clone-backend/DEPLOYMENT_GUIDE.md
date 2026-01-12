# Backend Deployment Guide - Google Cloud Platform

## Option 1: Google Cloud Run (Recommended for Getting Started)

Cloud Run is a fully managed serverless platform that runs containers. Perfect for microservices.

### Prerequisites

1. Install Google Cloud CLI:

```bash
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init
```

2. Authenticate:

```bash
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```

3. Enable required APIs:

```bash
gcloud services enable run.googleapis.com
gcloud services enable sql-component.googleapis.com
gcloud services enable sqladmin.googleapis.com
gcloud services enable containerregistry.googleapis.com
```

### Step 1: Set Up Cloud SQL (PostgreSQL)

```bash
# Create Cloud SQL instance for users database
gcloud sql instances create jumia-users-db \
    --database-version=POSTGRES_15 \
    --tier=db-f1-micro \
    --region=us-central1

# Create Cloud SQL instance for products database
gcloud sql instances create jumia-products-db \
    --database-version=POSTGRES_15 \
    --tier=db-f1-micro \
    --region=us-central1

# Set root password
gcloud sql users set-password postgres \
    --instance=jumia-users-db \
    --password=YOUR_SECURE_PASSWORD

gcloud sql users set-password postgres \
    --instance=jumia-products-db \
    --password=YOUR_SECURE_PASSWORD

# Create databases
gcloud sql databases create jumia_users --instance=jumia-users-db
gcloud sql databases create jumia_products --instance=jumia-products-db
```

### Step 2: Create Dockerfiles for Each Service

Each service needs a production Dockerfile. I'll create them below.

### Step 3: Build and Push Docker Images

```bash
# Set your project ID
export PROJECT_ID=YOUR_GCP_PROJECT_ID

# User Service
cd services/user-service
docker build -t gcr.io/$PROJECT_ID/user-service:latest .
docker push gcr.io/$PROJECT_ID/user-service:latest

# Product Service
cd ../product-service
docker build -t gcr.io/$PROJECT_ID/product-service:latest .
docker push gcr.io/$PROJECT_ID/product-service:latest

# Cart Service
cd ../cart-service
docker build -t gcr.io/$PROJECT_ID/cart-service:latest .
docker push gcr.io/$PROJECT_ID/cart-service:latest

# Order Service
cd ../order-service
docker build -t gcr.io/$PROJECT_ID/order-service:latest .
docker push gcr.io/$PROJECT_ID/order-service:latest

# API Gateway
cd ../../api-gateway
docker build -t gcr.io/$PROJECT_ID/api-gateway:latest .
docker push gcr.io/$PROJECT_ID/api-gateway:latest
```

### Step 4: Deploy to Cloud Run

```bash
# Deploy User Service
gcloud run deploy user-service \
    --image gcr.io/$PROJECT_ID/user-service:latest \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --add-cloudsql-instances jumia-users-db \
    --set-env-vars DB_HOST=/cloudsql/YOUR_PROJECT_ID:us-central1:jumia-users-db,DB_NAME=jumia_users,DB_USER=postgres,DB_PASSWORD=YOUR_PASSWORD,DB_PORT=5432

# Deploy Product Service
gcloud run deploy product-service \
    --image gcr.io/$PROJECT_ID/product-service:latest \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --add-cloudsql-instances jumia-products-db \
    --set-env-vars DB_HOST=/cloudsql/YOUR_PROJECT_ID:us-central1:jumia-products-db,DB_NAME=jumia_products,DB_USER=postgres,DB_PASSWORD=YOUR_PASSWORD,DB_PORT=5432

# Deploy Cart Service
gcloud run deploy cart-service \
    --image gcr.io/$PROJECT_ID/cart-service:latest \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated

# Deploy Order Service
gcloud run deploy order-service \
    --image gcr.io/$PROJECT_ID/order-service:latest \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated

# Deploy API Gateway (with environment variables for service URLs)
gcloud run deploy api-gateway \
    --image gcr.io/$PROJECT_ID/api-gateway:latest \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --set-env-vars USER_SERVICE_URL=USER_SERVICE_CLOUD_RUN_URL,PRODUCT_SERVICE_URL=PRODUCT_SERVICE_CLOUD_RUN_URL,CART_SERVICE_URL=CART_SERVICE_CLOUD_RUN_URL,ORDER_SERVICE_URL=ORDER_SERVICE_CLOUD_RUN_URL
```

### Step 5: Get Service URLs

```bash
gcloud run services list --platform managed
```

Update your frontend `.env` file with the API Gateway URL.

---

## Option 2: Google Kubernetes Engine (GKE) - Production Grade

Better for: Production environments, complex orchestration, fine-grained control

### Prerequisites

```bash
# Enable Kubernetes Engine API
gcloud services enable container.googleapis.com

# Install kubectl
gcloud components install kubectl
```

### Step 1: Create GKE Cluster

```bash
gcloud container clusters create jumia-cluster \
    --num-nodes=3 \
    --machine-type=e2-medium \
    --region=us-central1 \
    --enable-autoscaling \
    --min-nodes=1 \
    --max-nodes=5
```

### Step 2: Configure kubectl

```bash
gcloud container clusters get-credentials jumia-cluster --region=us-central1
```

### Step 3: Create Kubernetes Manifests

I'll create these in a `k8s/` folder:

- Deployments for each service
- Services for internal communication
- Ingress for external access
- ConfigMaps for configuration
- Secrets for sensitive data

### Step 4: Deploy to GKE

```bash
kubectl apply -f k8s/
```

---

## Option 3: Compute Engine (VMs) - Traditional Approach

Best for: Maximum control, existing VM management experience

```bash
# Create VM instance
gcloud compute instances create jumia-backend \
    --machine-type=e2-medium \
    --zone=us-central1-a \
    --image-family=ubuntu-2204-lts \
    --image-project=ubuntu-os-cloud \
    --boot-disk-size=20GB

# SSH into instance
gcloud compute ssh jumia-backend --zone=us-central1-a

# Install Docker and Docker Compose
sudo apt-get update
sudo apt-get install -y docker.io docker-compose
sudo usermod -aG docker $USER

# Clone your repository
git clone YOUR_REPO_URL
cd jumia-clone-backend

# Run with Docker Compose
sudo docker-compose up -d
```

---

## Cost Comparison

| Service                      | Cloud Run      | GKE             | Compute Engine |
| ---------------------------- | -------------- | --------------- | -------------- |
| Monthly Cost (light traffic) | $5-20          | $50-150         | $30-80         |
| Auto-scaling                 | ✅ Automatic   | ✅ Configure    | ❌ Manual      |
| Setup Complexity             | ⭐ Easy        | ⭐⭐⭐ Complex  | ⭐⭐ Medium    |
| Best For                     | Startups, MVPs | Production apps | Legacy apps    |

---

## Recommended Deployment Path

**For your use case, I recommend:**

1. **Start with Cloud Run** - quickest path to production
2. **Use Cloud SQL** for databases - managed, scalable
3. **Migrate to GKE** when you need:
   - Complex networking
   - Custom load balancing
   - Service mesh (Istio)
   - Multi-region deployments

---

## Environment Variables Management

### For Cloud Run:

```bash
# Create .env.production
cat > .env.production << EOF
DB_HOST=/cloudsql/PROJECT_ID:REGION:INSTANCE_NAME
DB_NAME=jumia_users
DB_USER=postgres
DB_PASSWORD=your_secure_password
JWT_SECRET=your_jwt_secret
GRPC_PORT=50051
EOF

# Deploy with env vars
gcloud run deploy SERVICE_NAME \
    --env-vars-file .env.production
```

### For GKE:

```bash
# Create Kubernetes secret
kubectl create secret generic jumia-secrets \
    --from-literal=db-password=YOUR_PASSWORD \
    --from-literal=jwt-secret=YOUR_JWT_SECRET
```

---

## Alternative: Railway, Render, or Fly.io

If you want simpler deployment than GCP:

### Railway (Easiest)

```bash
# Install Railway CLI
npm i -g @railway/cli

# Login and init
railway login
railway init
railway up
```

### Render

- Connect your GitHub repo
- Auto-deploys on push
- Free tier available

### Fly.io

```bash
# Install flyctl
curl -L https://fly.io/install.sh | sh

# Deploy
fly launch
fly deploy
```

---

## Next Steps

1. Choose deployment option (I recommend Cloud Run for now)
2. I'll create production Dockerfiles for all services
3. Set up CI/CD with GitHub Actions for automated deployments

Would you like me to set up Cloud Run deployment?
