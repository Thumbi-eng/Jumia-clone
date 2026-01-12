# Deploy to Google Cloud Platform - FREE Setup

# Uses Cloud Run (free tier) + External Free Database

## Prerequisites

1. Create GCP account (get $300 credit): https://cloud.google.com/free
2. Install gcloud CLI:

```bash
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init
```

3. Set your project:

```bash
export PROJECT_ID="your-project-id"
gcloud config set project $PROJECT_ID
```

4. Enable required APIs:

```bash
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

## Step 1: Set Up Free Database (Neon.tech)

1. Go to https://neon.tech (free PostgreSQL)
2. Sign up and create project "jumia-clone"
3. Create two databases:
   - `jumia_users`
   - `jumia_products`
4. Copy connection strings (you'll need them later)

Example connection string:

```
postgres://username:password@ep-xxx.us-east-2.aws.neon.tech/jumia_users
```

## Step 2: Create Production Dockerfiles

All Dockerfiles are already created in each service directory.

## Step 3: Build and Deploy Services

### Deploy User Service

```bash
cd services/user-service

# Build and submit to Cloud Build
gcloud builds submit --tag gcr.io/$PROJECT_ID/user-service

# Deploy to Cloud Run (FREE tier)
gcloud run deploy user-service \
  --image gcr.io/$PROJECT_ID/user-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 50051 \
  --memory 256Mi \
  --set-env-vars "DB_HOST=YOUR_NEON_HOST,DB_PORT=5432,DB_NAME=jumia_users,DB_USER=YOUR_USER,DB_PASSWORD=YOUR_PASSWORD,DB_SSLMODE=require,GRPC_PORT=50051,JWT_SECRET=your-secret-key"

# Save the service URL
export USER_SERVICE_URL=$(gcloud run services describe user-service --region us-central1 --format 'value(status.url)')
echo "User Service: $USER_SERVICE_URL"
```

### Deploy Product Service

```bash
cd ../product-service

gcloud builds submit --tag gcr.io/$PROJECT_ID/product-service

gcloud run deploy product-service \
  --image gcr.io/$PROJECT_ID/product-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 50052 \
  --memory 256Mi \
  --set-env-vars "DB_HOST=YOUR_NEON_HOST,DB_PORT=5432,DB_NAME=jumia_products,DB_USER=YOUR_USER,DB_PASSWORD=YOUR_PASSWORD,DB_SSLMODE=require,GRPC_PORT=50052"

export PRODUCT_SERVICE_URL=$(gcloud run services describe product-service --region us-central1 --format 'value(status.url)')
echo "Product Service: $PRODUCT_SERVICE_URL"
```

### Deploy Cart Service

```bash
cd ../cart-service

gcloud builds submit --tag gcr.io/$PROJECT_ID/cart-service

gcloud run deploy cart-service \
  --image gcr.io/$PROJECT_ID/cart-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 50053 \
  --memory 256Mi \
  --set-env-vars "GRPC_PORT=50053"

export CART_SERVICE_URL=$(gcloud run services describe cart-service --region us-central1 --format 'value(status.url)')
echo "Cart Service: $CART_SERVICE_URL"
```

### Deploy Order Service

```bash
cd ../order-service

gcloud builds submit --tag gcr.io/$PROJECT_ID/order-service

gcloud run deploy order-service \
  --image gcr.io/$PROJECT_ID/order-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 50054 \
  --memory 256Mi \
  --set-env-vars "GRPC_PORT=50054"

export ORDER_SERVICE_URL=$(gcloud run services describe order-service --region us-central1 --format 'value(status.url)')
echo "Order Service: $ORDER_SERVICE_URL"
```

### Deploy API Gateway

```bash
cd ../../api-gateway

gcloud builds submit --tag gcr.io/$PROJECT_ID/api-gateway

gcloud run deploy api-gateway \
  --image gcr.io/$PROJECT_ID/api-gateway \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --memory 512Mi \
  --set-env-vars "USER_SERVICE_URL=$USER_SERVICE_URL,PRODUCT_SERVICE_URL=$PRODUCT_SERVICE_URL,CART_SERVICE_URL=$CART_SERVICE_URL,ORDER_SERVICE_URL=$ORDER_SERVICE_URL,PORT=8080,JWT_SECRET=your-secret-key"

export API_GATEWAY_URL=$(gcloud run services describe api-gateway --region us-central1 --format 'value(status.url)')
echo "API Gateway: $API_GATEWAY_URL"
```

## Step 4: Update Frontend

Update your frontend `.env`:

```bash
cd ../../console/v1
cat > .env << EOF
VITE_API_BASE_URL=$API_GATEWAY_URL
VITE_FIREBASE_API_KEY=your_firebase_key
VITE_FIREBASE_AUTH_DOMAIN=your_domain
# ... other Firebase vars
EOF
```

## Step 5: Deploy Frontend (Also Free!)

### Option A: Firebase Hosting (Free)

```bash
npm install -g firebase-tools
firebase login
firebase init hosting
npm run build
firebase deploy
```

### Option B: Vercel (Free)

```bash
npm i -g vercel
vercel --prod
```

## Cost Monitoring

Check your Cloud Run usage:

```bash
# View requests
gcloud run services describe SERVICE_NAME --region us-central1 --format 'value(status.traffic)'

# Check if you're within free tier
gcloud billing accounts list
```

## Automation Script

Save this as `deploy-all.sh`:

```bash
#!/bin/bash
set -e

PROJECT_ID="your-project-id"
REGION="us-central1"

echo "🚀 Deploying Jumia Clone to GCP Cloud Run..."

# Array of services
declare -A services
services=(
  ["user-service"]="50051"
  ["product-service"]="50052"
  ["cart-service"]="50053"
  ["order-service"]="50054"
)

# Deploy microservices
for service in "${!services[@]}"; do
  port="${services[$service]}"
  echo "📦 Deploying $service on port $port..."

  cd "services/$service"
  gcloud builds submit --tag "gcr.io/$PROJECT_ID/$service"
  gcloud run deploy "$service" \
    --image "gcr.io/$PROJECT_ID/$service" \
    --platform managed \
    --region "$REGION" \
    --allow-unauthenticated \
    --port "$port" \
    --memory 256Mi
  cd ../..
done

# Deploy API Gateway
echo "📦 Deploying API Gateway..."
cd api-gateway
gcloud builds submit --tag "gcr.io/$PROJECT_ID/api-gateway"
gcloud run deploy api-gateway \
  --image "gcr.io/$PROJECT_ID/api-gateway" \
  --platform managed \
  --region "$REGION" \
  --allow-unauthenticated \
  --port 8080 \
  --memory 512Mi

echo "✅ Deployment complete!"
echo "API Gateway URL: $(gcloud run services describe api-gateway --region $REGION --format 'value(status.url)')"
```

Make it executable:

```bash
chmod +x deploy-all.sh
./deploy-all.sh
```

## Staying Within Free Tier

**Free limits:**

- 2M requests/month
- 360K GB-seconds compute
- 180K vCPU-seconds

**Tips to stay free:**

1. Use `--memory 256Mi` (minimum)
2. Set `--min-instances 0` (scales to zero)
3. Use `--cpu 1` (default)
4. Monitor usage in Cloud Console

## Alternative: f1-micro VM Setup

If you prefer a VM (also free):

```bash
# Create free VM
gcloud compute instances create jumia-backend \
  --machine-type f1-micro \
  --zone us-central1-a \
  --image-family ubuntu-2204-lts \
  --image-project ubuntu-os-cloud \
  --boot-disk-size 30GB

# SSH and setup
gcloud compute ssh jumia-backend --zone us-central1-a

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Clone and run
git clone YOUR_REPO
cd jumia-clone-backend
sudo docker-compose up -d
```

## Cleanup (if needed)

Delete all Cloud Run services:

```bash
gcloud run services delete user-service --region us-central1 -q
gcloud run services delete product-service --region us-central1 -q
gcloud run services delete cart-service --region us-central1 -q
gcloud run services delete order-service --region us-central1 -q
gcloud run services delete api-gateway --region us-central1 -q
```

Delete images:

```bash
gcloud container images delete gcr.io/$PROJECT_ID/user-service
gcloud container images delete gcr.io/$PROJECT_ID/product-service
# ... etc
```
