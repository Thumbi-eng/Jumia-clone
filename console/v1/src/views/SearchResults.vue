<template>
  <v-container class="py-6">
    <!-- Search Query Header -->
    <div v-if="productStore.searchQuery" class="mb-6">
      <h1 class="text-h5 font-weight-medium mb-2">
        Search Results for "{{ productStore.searchQuery }}"
      </h1>
      <p class="text-body-2 text-grey-darken-1">
        {{ productStore.totalProducts }}
        {{ productStore.totalProducts === 1 ? "product" : "products" }} found
      </p>
    </div>

    <!-- Loading State -->
    <div v-if="productStore.loading" class="text-center py-12">
      <v-progress-circular indeterminate color="orange" size="64" />
      <p class="text-body-1 mt-4">Searching products...</p>
    </div>

    <!-- Error State -->
    <v-alert
      v-else-if="productStore.error"
      type="error"
      variant="tonal"
      closable
      class="mb-4"
      @click:close="productStore.clearError"
    >
      {{ productStore.error }}
    </v-alert>

    <!-- No Results -->
    <div
      v-else-if="!productStore.loading && !productStore.hasSearchResults"
      class="text-center py-12"
    >
      <v-icon size="120" color="grey-lighten-1">mdi-magnify</v-icon>
      <h2 class="text-h5 mt-4 mb-2">No products found</h2>
      <p class="text-body-1 text-grey-darken-1 mb-6">
        We couldn't find any products matching your search.
      </p>
      <v-btn color="orange" variant="flat" size="large" to="/">
        Browse All Products
      </v-btn>
    </div>

    <!-- Search Results Grid -->
    <div v-else>
      <v-row>
        <v-col
          v-for="product in productStore.searchResults"
          :key="product.id"
          cols="12"
          sm="6"
          md="4"
          lg="3"
          xl="2"
        >
          <v-card hover class="product-card" @click="goToProduct(product.id)">
            <v-img
              :src="product.image_url || '/images/placeholder.jpg'"
              :alt="product.name"
              aspect-ratio="1"
              cover
              class="bg-grey-lighten-3"
            >
              <!-- Discount Badge -->
              <v-chip
                v-if="product.discount_percentage > 0"
                color="orange"
                size="small"
                class="ma-2"
                label
              >
                -{{ product.discount_percentage }}%
              </v-chip>
            </v-img>

            <v-card-text class="pb-2">
              <h3 class="text-body-2 font-weight-medium mb-1 text-truncate-2">
                {{ product.name }}
              </h3>

              <div class="d-flex align-center gap-2 mb-2">
                <span class="text-h6 font-weight-bold text-grey-darken-4">
                  KSh {{ formatPrice(calculateDiscountedPrice(product)) }}
                </span>
                <span
                  v-if="product.discount_percentage > 0"
                  class="text-caption text-decoration-line-through text-grey"
                >
                  KSh {{ formatPrice(product.price) }}
                </span>
              </div>

              <!-- Stock Status -->
              <v-chip
                v-if="product.stock === 0"
                color="red"
                size="x-small"
                variant="flat"
              >
                Out of Stock
              </v-chip>
              <v-chip
                v-else-if="product.stock < 10"
                color="orange"
                size="x-small"
                variant="flat"
              >
                Only {{ product.stock }} left
              </v-chip>
              <v-chip v-else color="green" size="x-small" variant="flat">
                In Stock
              </v-chip>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Pagination -->
      <div
        v-if="productStore.totalPages > 1"
        class="d-flex justify-center mt-8"
      >
        <v-pagination
          v-model="productStore.currentPage"
          :length="productStore.totalPages"
          :total-visible="7"
          color="orange"
          @update:model-value="onPageChange"
        />
      </div>
    </div>
  </v-container>
</template>

<script setup>
import { onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useProductStore } from "@/stores/products";

const route = useRoute();
const router = useRouter();
const productStore = useProductStore();

// Perform search on mount and when query parameter changes
onMounted(() => {
  const query = route.query.q;
  if (query) {
    productStore.searchProducts(query);
  }
});

watch(
  () => route.query.q,
  (newQuery) => {
    if (newQuery) {
      productStore.searchProducts(newQuery);
    }
  }
);

function onPageChange(page) {
  if (route.query.q) {
    productStore.searchProducts(route.query.q, page);
  }
}

function goToProduct(productId) {
  // Navigate to product detail page
  router.push(`/product/${productId}`);
}

function calculateDiscountedPrice(product) {
  if (!product.discount_percentage || product.discount_percentage === 0) {
    return product.price;
  }
  return product.price * (1 - product.discount_percentage / 100);
}

function formatPrice(price) {
  return new Intl.NumberFormat("en-KE", {
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  }).format(price);
}
</script>

<style scoped>
.product-card {
  cursor: pointer;
  transition: all 0.3s ease;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1) !important;
}

.text-truncate-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.4;
  max-height: 2.8em;
}
</style>
