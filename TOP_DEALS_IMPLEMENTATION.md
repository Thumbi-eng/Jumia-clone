# Top Deals Implementation Guide

## Overview

The Top Deals feature allows admins to classify products as different types of deals (hot deals, clearance, limited offers, etc.) and display them on the homepage.

## Backend Changes

### 1. Database Schema (Product Model)

Added three new fields to the Product model:

- `is_top_deal` (bool): Flag to mark product as a deal
- `deal_type` (string): Category of deal (hot_deal, clearance, limited_offer, top_deal)
- `deal_priority` (int): Display order priority (higher numbers appear first)

### 2. Repository Layer

**File**: `services/product-service/internal/repository/product_repository.go`

New methods:

```go
GetTopDeals(page, pageSize int) ([]*models.Product, int64, error)
GetDealsByType(dealType string, page, pageSize int) ([]*models.Product, int64, error)
```

Both methods:

- Filter by `is_active = true` and `is_top_deal = true`
- Order by `deal_priority DESC` (highest priority first)
- Support pagination
- GetDealsByType additionally filters by specific deal_type

### 3. Service Layer

**File**: `services/product-service/internal/service/product_service.go`

Updated methods:

- `CreateProduct()`: Added isTopDeal, dealType, dealPriority parameters
- `UpdateProduct()`: Added isTopDeal, dealType, dealPriority parameters
- `GetTopDeals()`: Validates pagination and calls repository
- `GetDealsByType()`: Validates deal type and pagination

### 4. Handler Layer

**File**: `services/product-service/internal/handler/product_handler.go`

- Updated `CreateProduct` to extract deal fields from proto request
- Updated `UpdateProduct` to handle deal field updates
- Updated `convertToProductData` to include deal fields in response
- Added `GetTopDeals` handler
- Added `GetDealsByType` handler

### 5. Protocol Buffers

**File**: `api-gateway/proto/product.proto`

New RPC methods:

```protobuf
rpc GetTopDeals(GetTopDealsRequest) returns (GetTopDealsResponse) {}
rpc GetDealsByType(GetDealsByTypeRequest) returns (GetDealsByTypeResponse) {}
```

Updated messages:

- CreateProductRequest: Added is_top_deal, deal_type, deal_priority
- UpdateProductRequest: Added is_top_deal, deal_type, deal_priority
- ProductData: Added is_top_deal, deal_type, deal_priority

### 6. API Gateway

**File**: `api-gateway/internal/handler/product_handler.go`

New endpoints:

- `GET /api/v1/products/top-deals?page=1&page_size=20`
- `GET /api/v1/products/deals?deal_type=hot_deal&page=1&page_size=20`

## Frontend Changes

### 1. Admin Top Deals Page

**File**: `console/v1/src/views/admin/TopDeals.vue`

Features:

- Stats cards showing total deals by type (Hot Deals, Clearance, Limited Offers)
- Searchable product table with all products
- Mark/unmark products as top deals
- Select deal type (hot_deal, clearance, limited_offer, top_deal)
- Set deal priority (higher = displayed first)
- Visual badges showing deal status and type

**Route**: `/admin/top-deals`

### 2. Updated TopDealsRow Component

**File**: `console/v1/src/components/TopDealsRow.vue`

Changes:

- Now fetches live data from `/api/v1/products/top-deals`
- Displays deal type badges (HOT, SALE, LIMITED, DEAL)
- Shows actual product images, prices, and discounts
- Loading skeleton while fetching
- Empty state message when no deals
- Click handler for product navigation (ready for product detail page)

## Deal Types

| Deal Type     | Badge Text | Badge Color | Use Case                          |
| ------------- | ---------- | ----------- | --------------------------------- |
| hot_deal      | HOT        | Red         | Popular products with high demand |
| clearance     | SALE       | Orange      | Products being cleared out        |
| limited_offer | LIMITED    | Blue        | Time-sensitive special offers     |
| top_deal      | DEAL       | Green       | General featured deals            |

## Usage Instructions

### Adding a Top Deal

1. Navigate to `/admin/top-deals`
2. Find the product you want to feature
3. Click the edit icon (pencil)
4. Toggle "Mark as Top Deal" switch
5. Select a deal type from dropdown
6. Set priority (higher numbers appear first, e.g., 100 = top priority)
7. Click "Save"

### Viewing Top Deals

- Homepage automatically displays top deals in the "Top Deals Row" section
- Deals are ordered by priority (highest first)
- Only active products marked as top deals appear
- Limited to 12 deals on homepage (configurable)

## API Endpoints

### Get Top Deals

```bash
GET /api/v1/products/top-deals
Query Params:
  - page: int (default: 1)
  - page_size: int (default: 20, max: 100)

Response:
{
  "success": true,
  "message": "Top deals retrieved successfully",
  "products": [...],
  "total": 10
}
```

### Get Deals by Type

```bash
GET /api/v1/products/deals
Query Params:
  - deal_type: string (hot_deal, clearance, limited_offer, top_deal)
  - page: int (default: 1)
  - page_size: int (default: 20, max: 100)

Response:
{
  "success": true,
  "message": "Deals retrieved successfully",
  "products": [...],
  "total": 5
}
```

### Update Product with Deal Fields

```bash
PUT /api/v1/products/:id
Body:
{
  "name": "Product Name",
  "price": 1000,
  "is_top_deal": true,
  "deal_type": "hot_deal",
  "deal_priority": 100,
  ...other fields
}
```

## Database Migration

The product service automatically migrates the database on startup. New fields added:

```sql
is_top_deal BOOLEAN DEFAULT false
deal_type VARCHAR(50)
deal_priority INTEGER DEFAULT 0
```

## Testing

1. **Backend Services Running**:

   - Product Service: Port 50052 ✓
   - API Gateway: Port 8080 ✓

2. **Test Top Deals Endpoint**:

   ```bash
   curl "http://localhost:8080/api/v1/products/top-deals?page=1&page_size=5"
   ```

3. **Create a Top Deal**:
   - Go to `/admin/products/add` and create a product
   - Or use existing products in `/admin/top-deals`
   - Mark as top deal with type and priority
   - Verify it appears on homepage

## Future Enhancements

1. **Deal Analytics**: Track views, clicks, and conversions per deal type
2. **Scheduled Deals**: Auto-enable/disable deals based on date/time
3. **Deal Templates**: Save deal configurations for quick reuse
4. **A/B Testing**: Test different deal types and priorities
5. **Deal Recommendations**: AI-suggested products to feature as deals
6. **Bulk Operations**: Mark multiple products as deals at once
7. **Deal Categories**: Create custom deal categories beyond the 4 types
8. **Performance Metrics**: Show which deal types perform best

## Notes

- Deal priority doesn't affect flash sales (separate system)
- Products can be both flash sale AND top deal simultaneously
- Deal type is required when marking product as top deal
- Unmarking as top deal clears deal_type and sets priority to 0
- Frontend automatically refreshes when deal status changes
