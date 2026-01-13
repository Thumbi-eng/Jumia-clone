<template>
  <v-card class="mx-auto mb-8" flat>
    <v-card-title class="d-flex align-center pa-4">
      <v-icon icon="mdi-grid" class="mr-2" color="orange"></v-icon>
      <span class="text-h6 font-weight-bold">All Products</span>
    </v-card-title>

    <v-divider></v-divider>

    <v-card-text class="pa-4">
      <v-row dense>
        <v-col
          v-for="product in products"
          :key="product.id"
          cols="6"
          sm="4"
          md="3"
          lg="2"
        >
          <v-card
            @click="goToProduct(product.id)"
            class="product-card"
            elevation="0"
            border
            hover
          >
            <v-img
              :src="product.image"
              :alt="product.name"
              height="180"
              cover
              class="bg-grey-lighten-4"
            >
              <template v-if="product.discount" v-slot:placeholder>
                <v-chip color="orange" size="small" class="ma-2" label>
                  -{{ product.discount }}%
                </v-chip>
              </template>
            </v-img>

            <v-card-text class="pa-3">
              <div class="product-name text-body-2 mb-2">
                {{ product.name }}
              </div>

              <div class="d-flex align-center mb-1">
                <span class="text-h6 font-weight-bold orange--text">
                  KSh {{ formatPrice(product.price) }}
                </span>
              </div>

              <div v-if="product.originalPrice" class="d-flex align-center">
                <span
                  class="text-caption text-decoration-line-through text-grey mr-2"
                >
                  KSh {{ formatPrice(product.originalPrice) }}
                </span>
                <span class="text-caption orange--text">
                  -{{ product.discount }}%
                </span>
              </div>

              <div v-if="product.rating" class="d-flex align-center mt-2">
                <v-rating
                  :model-value="product.rating"
                  density="compact"
                  size="x-small"
                  color="orange"
                  readonly
                  half-increments
                ></v-rating>
                <span class="text-caption text-grey ml-1">
                  ({{ product.reviews }})
                </span>
              </div>

              <v-chip
                v-if="product.inStock"
                color="success"
                size="x-small"
                class="mt-2"
                variant="tonal"
              >
                In Stock
              </v-chip>
              <v-chip
                v-else
                color="error"
                size="x-small"
                class="mt-2"
                variant="tonal"
              >
                Out of Stock
              </v-chip>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Load More Button -->
      <div v-if="hasMore" class="text-center mt-6">
        <v-btn
          color="orange"
          variant="outlined"
          size="large"
          @click="loadMore"
          :loading="loading"
        >
          Load More Products
        </v-btn>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const products = ref([]);
const loading = ref(false);
const hasMore = ref(true);
const page = ref(1);

// Sample products data
const allProducts = [
  {
    id: 1,
    name: "Samsung Galaxy A54 5G - 256GB",
    price: 42999,
    originalPrice: 52999,
    discount: 19,
    image: "/images/phone1.jpg",
    rating: 4.5,
    reviews: 324,
    inStock: true,
  },
  {
    id: 2,
    name: "Infinix Note 30 Pro - 256GB",
    price: 28499,
    originalPrice: 34999,
    discount: 19,
    image: "/images/phone2.jpg",
    rating: 4.3,
    reviews: 156,
    inStock: true,
  },
  {
    id: 3,
    name: "Redmi Note 13 Pro+ 5G",
    price: 36999,
    originalPrice: 44999,
    discount: 18,
    image: "/images/phone3.jpg",
    rating: 4.6,
    reviews: 289,
    inStock: true,
  },
  {
    id: 4,
    name: "Tecno Camon 20 Premier",
    price: 32999,
    image: "/images/phone4.jpg",
    rating: 4.2,
    reviews: 98,
    inStock: true,
  },
  {
    id: 5,
    name: 'HP 15.6" Laptop - Intel Core i5',
    price: 54999,
    originalPrice: 69999,
    discount: 21,
    image: "/images/laptop1.jpg",
    rating: 4.4,
    reviews: 167,
    inStock: true,
  },
  {
    id: 6,
    name: "Dell Inspiron 15 - AMD Ryzen 5",
    price: 48999,
    originalPrice: 59999,
    discount: 18,
    image: "/images/laptop2.jpg",
    rating: 4.5,
    reviews: 234,
    inStock: true,
  },
  {
    id: 7,
    name: "Sony WH-1000XM5 Headphones",
    price: 29999,
    originalPrice: 39999,
    discount: 25,
    image: "/images/headphones1.jpg",
    rating: 4.8,
    reviews: 512,
    inStock: true,
  },
  {
    id: 8,
    name: "JBL Flip 6 Bluetooth Speaker",
    price: 12499,
    originalPrice: 15999,
    discount: 22,
    image: "/images/speaker1.jpg",
    rating: 4.6,
    reviews: 345,
    inStock: true,
  },
  {
    id: 9,
    name: 'Samsung 55" 4K Smart TV',
    price: 64999,
    originalPrice: 84999,
    discount: 24,
    image: "/images/tv1.jpg",
    rating: 4.7,
    reviews: 189,
    inStock: true,
  },
  {
    id: 10,
    name: 'LG 43" Full HD Smart TV',
    price: 34999,
    originalPrice: 44999,
    discount: 22,
    image: "/images/tv2.jpg",
    rating: 4.4,
    reviews: 267,
    inStock: true,
  },
  {
    id: 11,
    name: "Apple Watch Series 9",
    price: 48999,
    image: "/images/watch1.jpg",
    rating: 4.7,
    reviews: 423,
    inStock: true,
  },
  {
    id: 12,
    name: "Samsung Galaxy Watch 6",
    price: 35999,
    originalPrice: 42999,
    discount: 16,
    image: "/images/watch2.jpg",
    rating: 4.5,
    reviews: 198,
    inStock: true,
  },
  {
    id: 13,
    name: "Canon EOS M50 Camera",
    price: 74999,
    originalPrice: 89999,
    discount: 17,
    image: "/images/camera1.jpg",
    rating: 4.8,
    reviews: 156,
    inStock: false,
  },
  {
    id: 14,
    name: "PlayStation 5 Console",
    price: 69999,
    image: "/images/ps5.jpg",
    rating: 4.9,
    reviews: 678,
    inStock: false,
  },
  {
    id: 15,
    name: "Xbox Series X",
    price: 64999,
    image: "/images/xbox.jpg",
    rating: 4.8,
    reviews: 534,
    inStock: true,
  },
  {
    id: 16,
    name: "Logitech MX Master 3S Mouse",
    price: 8499,
    originalPrice: 10999,
    discount: 23,
    image: "/images/mouse1.jpg",
    rating: 4.7,
    reviews: 289,
    inStock: true,
  },
  {
    id: 17,
    name: "Mechanical Gaming Keyboard RGB",
    price: 6499,
    originalPrice: 8999,
    discount: 28,
    image: "/images/keyboard1.jpg",
    rating: 4.4,
    reviews: 412,
    inStock: true,
  },
  {
    id: 18,
    name: "iPad Air 5th Gen - 64GB",
    price: 74999,
    originalPrice: 89999,
    discount: 17,
    image: "/images/ipad1.jpg",
    rating: 4.8,
    reviews: 234,
    inStock: true,
  },
];

onMounted(() => {
  loadProducts();
});

const loadProducts = () => {
  // Load first 12 products initially
  const itemsPerPage = 12;
  const start = (page.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;

  products.value = allProducts.slice(0, end);
  hasMore.value = end < allProducts.length;
};

const loadMore = () => {
  loading.value = true;
  page.value++;

  setTimeout(() => {
    loadProducts();
    loading.value = false;
  }, 500);
};

const formatPrice = (price) => {
  return price.toLocaleString();
};

const goToProduct = (id) => {
  router.push(`/product/${id}`);
};
</script>

<style scoped>
.product-card {
  cursor: pointer;
  transition: all 0.3s ease;
  height: 100%;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1) !important;
}

.product-name {
  height: 40px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  line-height: 1.4;
}

.orange--text {
  color: #ff9800 !important;
}
</style>
