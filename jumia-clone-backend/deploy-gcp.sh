#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Jumia Clone - Cloud Run Deployment Script${NC}\n"

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo -e "${RED}❌ gcloud CLI not found. Install it from: https://cloud.google.com/sdk/docs/install${NC}"
    exit 1
fi

# Get project ID
PROJECT_ID=$(gcloud config get-value project 2>/dev/null)
if [ -z "$PROJECT_ID" ]; then
    echo -e "${YELLOW}⚠️  No project set. Please set your GCP project:${NC}"
    echo "gcloud config set project YOUR_PROJECT_ID"
    exit 1
fi

REGION="us-central1"

echo -e "${GREEN}📋 Project: $PROJECT_ID${NC}"
echo -e "${GREEN}📍 Region: $REGION${NC}\n"

# Prompt for database credentials
echo -e "${YELLOW}📊 Database Setup${NC}"
echo "You need a PostgreSQL database. Options:"
echo "1. Neon.tech (Free): https://neon.tech"
echo "2. Supabase (Free): https://supabase.com"
echo "3. Cloud SQL (Paid): ~\$10/month"
echo ""

read -p "Enter DB Host (e.g., ep-xyz.aws.neon.tech): " DB_HOST
read -p "Enter DB Name for users (default: jumia_users): " DB_NAME_USERS
DB_NAME_USERS=${DB_NAME_USERS:-jumia_users}
read -p "Enter DB Name for products (default: jumia_products): " DB_NAME_PRODUCTS
DB_NAME_PRODUCTS=${DB_NAME_PRODUCTS:-jumia_products}
read -p "Enter DB User: " DB_USER
read -sp "Enter DB Password: " DB_PASSWORD
echo ""
read -p "Enter DB Port (default: 5432): " DB_PORT
DB_PORT=${DB_PORT:-5432}
read -p "Enter DB SSL Mode (default: require): " DB_SSLMODE
DB_SSLMODE=${DB_SSLMODE:-require}
read -p "Enter JWT Secret (for authentication): " JWT_SECRET

echo ""
echo -e "${GREEN}✅ Configuration complete!${NC}\n"

# Enable required services
echo -e "${YELLOW}🔧 Enabling required GCP services...${NC}"
gcloud services enable run.googleapis.com
gcloud services enable cloudbuild.googleapis.com
gcloud services enable containerregistry.googleapis.com
echo -e "${GREEN}✅ Services enabled${NC}\n"

# Function to deploy a service
deploy_service() {
    local SERVICE_NAME=$1
    local SERVICE_DIR=$2
    local PORT=$3
    local ENV_VARS=$4
    
    echo -e "${YELLOW}📦 Building and deploying $SERVICE_NAME...${NC}"
    
    cd "$SERVICE_DIR"
    
    # Build and push image
    gcloud builds submit --tag "gcr.io/$PROJECT_ID/$SERVICE_NAME" --quiet
    
    # Deploy to Cloud Run
    gcloud run deploy "$SERVICE_NAME" \
        --image "gcr.io/$PROJECT_ID/$SERVICE_NAME" \
        --platform managed \
        --region "$REGION" \
        --allow-unauthenticated \
        --port "$PORT" \
        --memory 512Mi \
        --cpu 1 \
        --min-instances 0 \
        --max-instances 10 \
        --set-env-vars "$ENV_VARS" \
        --quiet
    
    # Get service URL
    SERVICE_URL=$(gcloud run services describe "$SERVICE_NAME" --region "$REGION" --format 'value(status.url)')
    echo -e "${GREEN}✅ $SERVICE_NAME deployed: $SERVICE_URL${NC}\n"
    
    cd - > /dev/null
    echo "$SERVICE_URL"
}

# Deploy User Service
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
echo -e "${YELLOW}Deploying User Service${NC}"
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
USER_SERVICE_URL=$(deploy_service "user-service" "services/user-service" "50051" \
    "DB_HOST=$DB_HOST,DB_PORT=$DB_PORT,DB_NAME=$DB_NAME_USERS,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,DB_SSLMODE=$DB_SSLMODE,GRPC_PORT=50051,JWT_SECRET=$JWT_SECRET")

# Deploy Product Service
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
echo -e "${YELLOW}Deploying Product Service${NC}"
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
PRODUCT_SERVICE_URL=$(deploy_service "product-service" "services/product-service" "50052" \
    "DB_HOST=$DB_HOST,DB_PORT=$DB_PORT,DB_NAME=$DB_NAME_PRODUCTS,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,DB_SSLMODE=$DB_SSLMODE,GRPC_PORT=50052")

# Deploy Cart Service
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
echo -e "${YELLOW}Deploying Cart Service${NC}"
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
CART_SERVICE_URL=$(deploy_service "cart-service" "services/cart-service" "50053" "GRPC_PORT=50053")

# Deploy Order Service
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
echo -e "${YELLOW}Deploying Order Service${NC}"
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
ORDER_SERVICE_URL=$(deploy_service "order-service" "services/order-service" "50054" "GRPC_PORT=50054")

# Deploy API Gateway
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
echo -e "${YELLOW}Deploying API Gateway${NC}"
echo -e "${YELLOW}═══════════════════════════════════════${NC}"
API_GATEWAY_URL=$(deploy_service "api-gateway" "api-gateway" "8080" \
    "USER_SERVICE_URL=$USER_SERVICE_URL,PRODUCT_SERVICE_URL=$PRODUCT_SERVICE_URL,CART_SERVICE_URL=$CART_SERVICE_URL,ORDER_SERVICE_URL=$ORDER_SERVICE_URL,JWT_SECRET=$JWT_SECRET")

# Summary
echo -e "\n${GREEN}═══════════════════════════════════════${NC}"
echo -e "${GREEN}🎉 Deployment Complete!${NC}"
echo -e "${GREEN}═══════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Service URLs:${NC}"
echo -e "User Service:    ${GREEN}$USER_SERVICE_URL${NC}"
echo -e "Product Service: ${GREEN}$PRODUCT_SERVICE_URL${NC}"
echo -e "Cart Service:    ${GREEN}$CART_SERVICE_URL${NC}"
echo -e "Order Service:   ${GREEN}$ORDER_SERVICE_URL${NC}"
echo -e "API Gateway:     ${GREEN}$API_GATEWAY_URL${NC}"

echo -e "\n${YELLOW}📝 Next Steps:${NC}"
echo "1. Update your frontend .env file:"
echo -e "   ${GREEN}VITE_API_BASE_URL=$API_GATEWAY_URL${NC}"
echo ""
echo "2. Test the API:"
echo -e "   ${GREEN}curl $API_GATEWAY_URL/health${NC}"
echo ""
echo "3. Monitor services:"
echo -e "   ${GREEN}gcloud run services list --region $REGION${NC}"
echo ""
echo "4. View logs:"
echo -e "   ${GREEN}gcloud run services logs read SERVICE_NAME --region $REGION${NC}"

# Save URLs to file
cat > deployment-urls.txt << EOF
Deployment completed at: $(date)
Project: $PROJECT_ID
Region: $REGION

Service URLs:
- User Service: $USER_SERVICE_URL
- Product Service: $PRODUCT_SERVICE_URL
- Cart Service: $CART_SERVICE_URL
- Order Service: $ORDER_SERVICE_URL
- API Gateway: $API_GATEWAY_URL

Frontend .env update:
VITE_API_BASE_URL=$API_GATEWAY_URL
EOF

echo -e "\n${GREEN}✅ Deployment URLs saved to: deployment-urls.txt${NC}"
