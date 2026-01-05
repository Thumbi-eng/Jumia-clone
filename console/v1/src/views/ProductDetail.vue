<template>
  <v-container fluid class="pa-0 mt-7 mb-4">
    <v-container class="py-4 mt-8 bg-white" style="max-width: 1200px">
      <v-row v-if="!loading && product">
        <!-- Left Column - Product Images -->
        <v-col cols="12" md="5">
          <v-card elevation="1" rounded="lg">
            <!-- Main Image -->
            <v-img
              :src="selectedImage || product.image_url"
              aspect-ratio="1"
              cover
              class="rounded-t-lg"
            ></v-img>

            <!-- Thumbnail Images -->
            <div class="d-flex pa-2 gap-2">
              <v-card
                v-for="(img, index) in productImages"
                :key="index"
                :class="{ 'border-primary': selectedImage === img }"
                class="cursor-pointer"
                style="width: 60px; height: 60px"
                elevation="0"
                @click="selectedImage = img"
              >
                <v-img :src="img" cover height="100%"></v-img>
              </v-card>
            </div>

            <!-- Share Section -->
            <div class="pa-4">
              <div class="text-caption text-grey mb-2">SHARE THIS PRODUCT</div>
              <div class="d-flex gap-2">
                <v-btn icon size="small" variant="outlined">
                  <v-icon>mdi-facebook</v-icon>
                </v-btn>
                <v-btn icon size="small" variant="outlined">
                  <v-icon>mdi-twitter</v-icon>
                </v-btn>
                <v-btn icon size="small" variant="outlined">
                  <v-icon>mdi-whatsapp</v-icon>
                </v-btn>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Right Column - Product Details -->
        <v-col cols="12" md="7">
          <v-row>
            <!-- Product Info -->
            <v-col cols="12" md="8">
              <div class="mb-2">
                <v-chip size="small" color="grey" class="mr-2">{{
                  product.brand
                }}</v-chip>
                <v-chip v-if="product.in_stock" size="small" color="success"
                  >In Stock</v-chip
                >
                <v-chip v-else size="small" color="error">Out of Stock</v-chip>
              </div>

              <h1 class="text-h5 mb-3">{{ product.name }}</h1>

              <div class="d-flex align-center mb-3">
                <v-rating
                  :model-value="4.5"
                  density="compact"
                  size="small"
                  readonly
                  color="orange"
                  half-increments
                ></v-rating>
                <span class="text-caption ml-2">(77 verified ratings)</span>
              </div>

              <v-divider class="my-3"></v-divider>

              <!-- Price -->
              <div class="mb-4">
                <div class="text-h4 font-weight-bold mb-1">
                  KSh {{ product.final_price.toLocaleString() }}
                </div>
                <div
                  v-if="product.discount_percentage > 0"
                  class="d-flex align-center gap-2"
                >
                  <span class="text-h6 text-grey text-decoration-line-through">
                    KSh {{ product.price.toLocaleString() }}
                  </span>
                  <v-chip color="orange" size="small">
                    -{{ Math.round(product.discount_percentage) }}%
                  </v-chip>
                </div>
              </div>

              <!-- Flash Sale Info -->
              <v-alert
                v-if="product.is_flash_sale && product.is_flash_sale_active"
                color="error"
                variant="tonal"
                density="compact"
                class="mb-4"
              >
                <div class="d-flex align-center justify-space-between">
                  <span class="font-weight-bold">⚡ Flash Sale</span>
                  <span
                    >Ends: {{ formatDate(product.flash_sale_end_time) }}</span
                  >
                </div>
              </v-alert>

              <!-- Deal Badge -->
              <v-chip
                v-if="product.is_top_deal"
                :color="getDealColor(product.deal_type)"
                class="mb-4"
                label
              >
                {{ formatDealType(product.deal_type) }}
              </v-chip>

              <!-- Add to Cart Button -->
              <v-btn
                block
                color="orange"
                size="x-large"
                class="mb-3 text-none"
                :disabled="!product.in_stock"
                @click="addToCart"
              >
                <v-icon left>mdi-cart</v-icon>
                Add to Cart
              </v-btn>

              <!-- Promotions -->
              <v-card variant="outlined" class="mb-4">
                <v-card-text>
                  <div class="text-caption font-weight-bold mb-2">
                    PROMOTIONS
                  </div>
                  <div class="d-flex align-center">
                    <v-icon color="primary" size="small" class="mr-2"
                      >mdi-wallet-giftcard</v-icon
                    >
                    <span class="text-body-2"
                      >Easy and safe payments via the JumiaPay App.</span
                    >
                  </div>
                </v-card-text>
              </v-card>
            </v-col>

            <!-- Delivery & Returns -->
            <v-col cols="12" md="4">
              <v-card variant="outlined" class="mb-3">
                <v-card-title class="text-body-1"
                  >DELIVERY & RETURNS</v-card-title
                >
                <v-card-text>
                  <div class="mb-3">
                    <div class="text-caption text-grey mb-1">
                      Choose your location
                    </div>
                    <v-select
                      v-model="selectedCity"
                      :items="cities"
                      variant="outlined"
                      density="compact"
                      placeholder="Nyeri"
                    ></v-select>
                    <v-select
                      v-model="selectedTown"
                      :items="towns"
                      variant="outlined"
                      density="compact"
                      placeholder="Nyeri Town"
                    ></v-select>
                  </div>

                  <div class="mb-3">
                    <v-icon size="small" class="mr-2">mdi-store</v-icon>
                    <span class="text-body-2">Pickup Station</span>
                    <div class="text-caption text-grey ml-6">
                      Delivery Fees KSh 170
                    </div>
                    <div class="text-caption ml-6">
                      Ready for pickup between 05 January & 11 January
                    </div>
                  </div>

                  <div>
                    <v-icon size="small" class="mr-2"
                      >mdi-keyboard-return</v-icon
                    >
                    <span class="text-body-2">Return Policy</span>
                    <div class="text-caption ml-6">
                      Easy Return, Quick Refund.
                      <a href="#" class="text-primary">Details</a>
                    </div>
                  </div>
                </v-card-text>
              </v-card>

              <!-- Seller Information -->
              <v-card variant="outlined">
                <v-card-title class="text-body-1"
                  >SELLER INFORMATION</v-card-title
                >
                <v-card-text>
                  <div class="mb-2">
                    <div class="font-weight-bold">{{ product.brand }}</div>
                    <div class="text-caption">
                      {{ product.stock }} units available
                    </div>
                  </div>
                  <v-btn
                    variant="outlined"
                    size="small"
                    color="orange"
                    block
                    class="text-none"
                  >
                    Follow
                  </v-btn>
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>

          <!-- Product Details Tabs -->
          <v-card variant="outlined" class="mt-4">
            <v-tabs v-model="tab" bg-color="grey-lighten-4">
              <v-tab value="details">
                <v-icon left>mdi-text-box</v-icon>
                Product Details
              </v-tab>
              <v-tab value="specifications">
                <v-icon left>mdi-format-list-bulleted</v-icon>
                Specifications
              </v-tab>
              <v-tab value="feedback">
                <v-icon left>mdi-comment</v-icon>
                Customer Feedback
              </v-tab>
            </v-tabs>

            <v-window v-model="tab">
              <v-window-item value="details">
                <v-card-text>
                  <h3 class="mb-3">Product Details</h3>
                  <div class="text-body-2" style="white-space: pre-line">
                    {{ product.description }}
                  </div>

                  <div v-if="product.category" class="mt-4">
                    <v-chip size="small" class="mr-2">
                      Category: {{ product.category }}
                    </v-chip>
                  </div>
                </v-card-text>
              </v-window-item>

              <v-window-item value="specifications">
                <v-card-text>
                  <h3 class="mb-3">Specifications</h3>
                  <v-table>
                    <tbody>
                      <tr>
                        <td class="font-weight-bold">Brand</td>
                        <td>{{ product.brand }}</td>
                      </tr>
                      <tr>
                        <td class="font-weight-bold">Category</td>
                        <td>{{ product.category }}</td>
                      </tr>
                      <tr>
                        <td class="font-weight-bold">SKU</td>
                        <td>{{ product.id }}</td>
                      </tr>
                      <tr>
                        <td class="font-weight-bold">Stock</td>
                        <td>{{ product.stock }} units</td>
                      </tr>
                      <tr v-if="product.discount_percentage > 0">
                        <td class="font-weight-bold">Discount</td>
                        <td>{{ Math.round(product.discount_percentage) }}%</td>
                      </tr>
                    </tbody>
                  </v-table>
                </v-card-text>
              </v-window-item>

              <v-window-item value="feedback">
                <v-card-text>
                  <h3 class="mb-3">Customer Feedback</h3>
                  <v-alert type="info" variant="tonal">
                    No reviews yet. Be the first to review this product!
                  </v-alert>
                </v-card-text>
              </v-window-item>
            </v-window>
          </v-card>
        </v-col>
      </v-row>

      <!-- Loading State -->
      <v-row v-if="loading">
        <v-col cols="12" md="5">
          <v-skeleton-loader
            type="image, list-item-two-line"
          ></v-skeleton-loader>
        </v-col>
        <v-col cols="12" md="7">
          <v-skeleton-loader
            type="heading, paragraph, button"
          ></v-skeleton-loader>
        </v-col>
      </v-row>

      <!-- Error State -->
      <v-alert v-if="error" type="error" class="my-4">
        {{ error }}
      </v-alert>
    </v-container>

    <!-- Snackbar -->
    <v-snackbar v-model="snackbar.show" :color="snackbar.color">
      {{ snackbar.message }}
      <template v-slot:actions>
        <v-btn variant="text" @click="snackbar.show = false">Close</v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useCartStore } from "@/stores/cart";
import axios from "axios";

const route = useRoute();
const router = useRouter();
const cartStore = useCartStore();
const API_URL = "http://localhost:8080/api/v1";

const product = ref(null);
const loading = ref(false);
const error = ref(null);
const selectedImage = ref(null);
const tab = ref("details");
const selectedCity = ref("Nyeri");
const selectedTown = ref("Nyeri Town");

const cities = ["Nairobi", "Mombasa", "Kisumu", "Nakuru", "Eldoret", "Nyeri"];
const towns = ["Nyeri Town", "Karatina", "Othaya", "Mukurweini"];

const snackbar = ref({
  show: false,
  message: "",
  color: "success",
});

const productImages = computed(() => {
  if (!product.value) return [];
  // For now, just show the main image
  // In production, you'd have multiple images from the backend
  return [product.value.image_url];
});

const loadProduct = async () => {
  loading.value = true;
  error.value = null;
  try {
    const productId = route.params.id;
    const response = await axios.get(`${API_URL}/products/${productId}`);

    if (response.data.success) {
      product.value = response.data.product;
      selectedImage.value = product.value.image_url;
    } else {
      error.value = response.data.message || "Product not found";
    }
  } catch (err) {
    console.error("Error loading product:", err);
    error.value = "Failed to load product. Please try again.";
  } finally {
    loading.value = false;
  }
};

const addToCart = () => {
  if (!product.value) return;

  cartStore.addToCart(product.value, 1);
  showSnackbar("Product added to cart!", "success");

  // Optional: Navigate to cart after a delay
  setTimeout(() => {
    router.push('/cart');
  }, 1000);
};

const getDealColor = (dealType) => {
  const colors = {
    hot_deal: "error",
    clearance: "warning",
    limited_offer: "info",
    top_deal: "success",
  };
  return colors[dealType] || "primary";
};

const formatDealType = (dealType) => {
  const formats = {
    hot_deal: "🔥 Hot Deal",
    clearance: "🏷️ Clearance Sale",
    limited_offer: "⏰ Limited Offer",
    top_deal: "⭐ Top Deal",
  };
  return formats[dealType] || dealType;
};

const formatDate = (dateString) => {
  if (!dateString) return "";
  const date = new Date(dateString);
  return date.toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const showSnackbar = (message, color = "success") => {
  snackbar.value = {
    show: true,
    message,
    color,
  };
};

onMounted(() => {
  loadProduct();
});
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}

.border-primary {
  border: 2px solid rgb(var(--v-theme-primary)) !important;
}

.gap-2 {
  gap: 8px;
}
</style>
