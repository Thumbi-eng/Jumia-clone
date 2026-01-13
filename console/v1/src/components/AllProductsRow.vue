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
              :src="getProductImage(product)"
              :alt="product.name"
              height="180"
              cover
              class="bg-grey-lighten-4"
            >
              <template v-slot:placeholder>
                <v-row class="fill-height ma-0" align="center" justify="center">
                  <v-progress-circular
                    indeterminate
                    color="grey-lighten-1"
                  ></v-progress-circular>
                </v-row>
              </template>

              <!-- Discount Badge -->
              <v-chip
                v-if="product.discount_percentage > 0"
                color="orange"
                size="small"
                class="ma-2"
                label
              >
                -{{ Math.round(product.discount_percentage) }}%
              </v-chip>
            </v-img>

            <v-card-text class="pa-3">
              <div class="product-name text-body-2 mb-2">
                {{ product.name }}
              </div>

              <div class="d-flex align-center mb-1">
                <span class="text-h6 font-weight-bold orange--text">
                  KSh {{ formatPrice(getFinalPrice(product)) }}
                </span>
              </div>

              <div
                v-if="product.discount_percentage > 0"
                class="d-flex align-center"
              >
                <span
                  class="text-caption text-decoration-line-through text-grey mr-2"
                >
                  KSh {{ formatPrice(product.price) }}
                </span>
                <span class="text-caption orange--text">
                  -{{ Math.round(product.discount_percentage) }}%
                </span>
              </div>

              <!-- Stock Status -->
              <v-chip
                v-if="isInStock(product)"
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

      <!-- Loading State -->
      <div
        v-if="productStore.loading && !products.length"
        class="text-center py-8"
      >
        <v-progress-circular
          indeterminate
          color="orange"
          size="64"
        ></v-progress-circular>
        <p class="text-grey mt-4">Loading products...</p>
      </div>

      <!-- Error State -->
      <div
        v-else-if="productStore.error && !products.length"
        class="text-center py-8"
      >
        <v-icon
          icon="mdi-alert-circle-outline"
          size="64"
          color="error"
          class="mb-4"
        ></v-icon>
        <p class="text-h6 mb-2">Failed to load products</p>
        <p class="text-grey mb-4">{{ productStore.error }}</p>
        <v-btn color="orange" variant="outlined" @click="loadProducts">
          <v-icon left>mdi-refresh</v-icon>
          Retry
        </v-btn>
      </div>

      <!-- Empty State -->
      <div
        v-else-if="!products.length && !productStore.loading"
        class="text-center py-8"
      >
        <v-icon
          icon="mdi-package-variant"
          size="64"
          color="grey"
          class="mb-4"
        ></v-icon>
        <p class="text-h6 text-grey">No products available</p>
      </div>

      <!-- Load More Button -->
      <div v-if="hasMore && products.length" class="text-center mt-6">
        <v-btn
          color="orange"
          variant="outlined"
          size="large"
          @click="loadMore"
          :loading="productStore.loading"
        >
          Load More Products
        </v-btn>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useProductStore } from "@/stores/products";

const router = useRouter();
const productStore = useProductStore();

const products = ref([]);
const page = ref(1);
const pageSize = ref(12);

// Computed
const hasMore = computed(() => {
  const totalLoaded = products.value.length;
  return totalLoaded < productStore.totalProducts;
});

onMounted(() => {
  loadProducts();
});

const loadProducts = async () => {
  try {
    const data = await productStore.fetchProducts(page.value, pageSize.value);

    if (data && data.products) {
      // Append new products to existing ones
      if (page.value === 1) {
        products.value = data.products;
      } else {
        products.value = [...products.value, ...data.products];
      }
    }
  } catch (err) {
    console.error("Error loading products:", err);
  }
};

const loadMore = async () => {
  page.value++;
  await loadProducts();
};

const formatPrice = (price) => {
  return price.toLocaleString();
};

const getFinalPrice = (product) => {
  if (product.discount_percentage > 0) {
    return product.price * (1 - product.discount_percentage / 100);
  }
  return product.price;
};

const isInStock = (product) => {
  return product.stock > 0 && product.is_active;
};

const getProductImage = (product) => {
  // If image_url exists and is not empty, use it
  if (product.image_url && product.image_url.trim() !== "") {
    return product.image_url;
  }
  // Fallback to placeholder image
  return "https://via.placeholder.com/300x300?text=No+Image";
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
