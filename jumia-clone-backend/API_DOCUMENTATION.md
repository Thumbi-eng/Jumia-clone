# Jumia Clone API Gateway

API Gateway running on `http://localhost:8080`

## Available Endpoints

### User Service (Authentication)

#### Register User

```bash
POST /api/v1/users/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+254700000000"
}
```

#### Login

```bash
POST /api/v1/users/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Get User (Requires Auth)

```bash
GET /api/v1/users/:id
Authorization: Bearer <token>
```

#### Update User (Requires Auth)

```bash
PUT /api/v1/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "first_name": "Jane",
  "last_name": "Doe",
  "phone": "+254700000001"
}
```

#### Verify Token

```bash
POST /api/v1/auth/verify
Content-Type: application/json

{
  "token": "<jwt_token>"
}
```

---

### Product Service

#### List Products

```bash
GET /api/v1/products?page=1&page_size=20
```

#### Search Products

```bash
GET /api/v1/products/search?q=iphone&page=1&page_size=20
```

#### Get Products by Category

```bash
GET /api/v1/products/category?category=Electronics&page=1&page_size=20
```

#### Get Single Product

```bash
GET /api/v1/products/:id
```

#### Create Product

```bash
POST /api/v1/products
Content-Type: application/json

{
  "name": "iPhone 15 Pro",
  "description": "Latest Apple iPhone",
  "price": 999.99,
  "category": "Electronics",
  "stock": 50,
  "image_url": "https://example.com/image.jpg",
  "brand": "Apple",
  "discount_percentage": 10
}
```

#### Update Product

```bash
PUT /api/v1/products/:id
Content-Type: application/json

{
  "name": "iPhone 15 Pro Max",
  "price": 1099.99,
  "stock": 30
}
```

#### Delete Product

```bash
DELETE /api/v1/products/:id
```

---

### Cart Service

#### Add to Cart

```bash
POST /api/v1/cart/add
Content-Type: application/json

{
  "user_id": "user-uuid",
  "product_id": "product-uuid",
  "product_name": "iPhone 15 Pro",
  "quantity": 2,
  "price": 999.99,
  "image_url": "https://example.com/image.jpg"
}
```

#### Update Cart Item

```bash
PUT /api/v1/cart/update
Content-Type: application/json

{
  "user_id": "user-uuid",
  "product_id": "product-uuid",
  "quantity": 5
}
```

#### Remove from Cart

```bash
POST /api/v1/cart/remove
Content-Type: application/json

{
  "user_id": "user-uuid",
  "product_id": "product-uuid"
}
```

#### Get Cart

```bash
GET /api/v1/cart/:user_id
```

#### Clear Cart

```bash
DELETE /api/v1/cart/:user_id
```

---

### Order Service

#### Create Order

```bash
POST /api/v1/orders
Content-Type: application/json

{
  "user_id": "user-uuid",
  "items": [
    {
      "product_id": "product-uuid",
      "product_name": "iPhone 15 Pro",
      "quantity": 2,
      "price": 999.99
    }
  ],
  "shipping_address": "123 Main St, Nairobi, Kenya",
  "payment_method": "M-Pesa"
}
```

#### Get Order

```bash
GET /api/v1/orders/:id
```

#### List User Orders

```bash
GET /api/v1/orders?user_id=user-uuid&page=1&page_size=10
```

#### Update Order Status

```bash
PUT /api/v1/orders/:id/status
Content-Type: application/json

{
  "status": "processing"
}
```

#### Cancel Order

```bash
POST /api/v1/orders/:id/cancel
Content-Type: application/json

{
  "user_id": "user-uuid"
}
```

---

## CORS Configuration

The API Gateway is configured to accept requests from:

- `http://localhost:3000` (React default)
- `http://localhost:5173` (Vite default)

## Services

- User Service: `localhost:50051` (gRPC)
- Product Service: `localhost:50052` (gRPC)
- Cart Service: `localhost:50053` (gRPC)
- Order Service: `localhost:50054` (gRPC)
- API Gateway: `localhost:8080` (HTTP)

## Frontend Integration

You can now use these HTTP endpoints from your Vue.js frontend. Example using fetch:

```javascript
// Get products
const response = await fetch(
  "http://localhost:8080/api/v1/products?page=1&page_size=20"
);
const data = await response.json();
console.log(data.products);

// Add to cart
const cartResponse = await fetch("http://localhost:8080/api/v1/cart/add", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    user_id: userId,
    product_id: productId,
    product_name: product.name,
    quantity: 1,
    price: product.price,
    image_url: product.image_url,
  }),
});
const cartData = await cartResponse.json();
```
