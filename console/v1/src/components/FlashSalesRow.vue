<template>
  <v-card
    class="flash-wrapper rounded-lg"
    elevation="0"
    style="max-width: 1550px; min-width: 1540px"
  >
    <!-- Header -->
    <div class="flash-header d-flex align-center justify-space-between px-4">
      <div class="d-flex align-center">
        <v-icon color="white" class="mr-2">mdi-flash</v-icon>
        <span class="text-white font-weight-bold">
          Flash Sales | {{ flashSalesCount }} Deals Live Now
        </span>
      </div>

      <div class="text-white font-weight-medium">
        Time Left:
        <span class="ml-1 countdown-timer">{{ timeLeft }}</span>
      </div>

      <v-btn
        variant="text"
        color="white"
        size="small"
        to="/admin/flash-sales"
        v-if="isAdmin"
      >
        Manage Flash Sales →
      </v-btn>
      <v-btn variant="text" color="white" size="small" v-else>
        See All →
      </v-btn>
    </div>

    <!-- Products -->
    <div
      v-if="loading"
      class="flash-products d-flex justify-center align-center px-3 py-8"
    >
      <v-progress-circular indeterminate color="orange" />
    </div>

    <div
      v-else-if="items.length === 0"
      class="flash-products px-3 py-8 text-center"
    >
      <div class="text-grey">
        <v-icon size="48" color="grey-lighten-1">mdi-flash-off</v-icon>
        <div class="mt-2">No active flash sales at the moment</div>
        <v-btn
          v-if="isAdmin"
          color="orange"
          variant="flat"
          size="small"
          class="mt-4"
          to="/admin/flash-sales"
        >
          Create Flash Sale
        </v-btn>
      </div>
    </div>

    <div v-else class="flash-products d-flex px-3 py-4">
      <div
        v-for="item in items"
        :key="item.id"
        class="flash-item mr-3"
        style="cursor: pointer"
        @click="goToProduct(item.id)"
      >
        <v-img :src="item.image" aspect-ratio="1" cover class="rounded mb-2" />

        <div class="text-body-2 truncate">
          {{ item.title }}
        </div>

        <div class="font-weight-bold text-red">
          KSh {{ formatPrice(item.price) }}
        </div>

        <div class="text-caption text-grey">
          <s>KSh {{ formatPrice(item.oldPrice) }}</s>
          <v-chip size="x-small" color="red" variant="flat" class="ml-1">
            -{{ item.discount }}%
          </v-chip>
        </div>

        <div class="text-caption mt-1">{{ item.left }} items left</div>

        <v-progress-linear
          :model-value="item.progress"
          height="6"
          rounded
          color="orange"
          class="mt-1"
        />
      </div>
    </div>
  </v-card>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from "vue";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";

const API_BASE = "http://localhost:8080/api/v1";
const router = useRouter();

const authStore = useAuthStore();
const loading = ref(false);
const timeLeft = ref("Loading...");
const items = ref([]);
const flashSalesCount = ref(0);
let timerInterval = null;
let flashSaleEndTime = null;

const isAdmin = computed(() => {
  // You can implement admin check here
  return authStore.isAuthenticated;
});

onMounted(() => {
  fetchFlashSales();
  startCountdown();
});

onUnmounted(() => {
  if (timerInterval) {
    clearInterval(timerInterval);
  }
});

async function fetchFlashSales() {
  loading.value = true;

  try {
    const response = await fetch(`${API_BASE}/products?page=1&page_size=100`);

    if (!response.ok) {
      throw new Error("Failed to load flash sales");
    }

    const data = await response.json();
    const allProducts = data.products || [];

    // Filter for active flash sales only
    const now = new Date();
    const flashSaleProducts = allProducts.filter((p) => {
      if (!p.is_flash_sale || !p.flash_sale_end_time) return false;
      const endTime = new Date(p.flash_sale_end_time);
      return endTime > now;
    });

    // Sort by end time (ending soon first) and limit to 6
    flashSaleProducts.sort((a, b) => {
      return new Date(a.flash_sale_end_time) - new Date(b.flash_sale_end_time);
    });

    flashSalesCount.value = flashSaleProducts.length;

    // Find the earliest end time for countdown
    if (flashSaleProducts.length > 0) {
      flashSaleEndTime = new Date(flashSaleProducts[0].flash_sale_end_time);
    }

    // Map to component format (take first 6)
    items.value = flashSaleProducts.slice(0, 6).map((product) => {
      const sold = product.initial_stock - product.stock;
      const progress =
        product.initial_stock > 0
          ? Math.round((sold / product.initial_stock) * 100)
          : 0;

      const discount = Math.round(
        ((product.price - product.flash_sale_price) / product.price) * 100
      );

      return {
        id: product.id,
        title: product.name,
        price: product.flash_sale_price,
        oldPrice: product.price,
        left: product.stock,
        progress: progress,
        image: product.image_url || "/images/placeholder.jpg",
        discount: discount,
      };
    });
  } catch (err) {
    console.error("Error fetching flash sales:", err);
  } finally {
    loading.value = false;
  }
}

function startCountdown() {
  updateCountdown(); // Initial update

  // Update every second
  timerInterval = setInterval(() => {
    updateCountdown();
  }, 1000);
}

function updateCountdown() {
  if (!flashSaleEndTime) {
    timeLeft.value = "No active sales";
    return;
  }

  const now = new Date();
  const diff = flashSaleEndTime - now;

  if (diff <= 0) {
    timeLeft.value = "Sale Ended";
    // Refresh flash sales to update list
    fetchFlashSales();
    return;
  }

  const hours = Math.floor(diff / (1000 * 60 * 60));
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));
  const seconds = Math.floor((diff % (1000 * 60)) / 1000);

  timeLeft.value = `${String(hours).padStart(2, "0")}h : ${String(
    minutes
  ).padStart(2, "0")}m : ${String(seconds).padStart(2, "0")}s`;
}

function formatPrice(price) {
  if (!price) return "0";
  return Number(price).toLocaleString("en-KE", {
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  });
}

function goToProduct(productId) {
  router.push(`/product/${productId}`);
}
</script>

<style scoped>
.flash-wrapper {
  background: #e50000;
}

.flash-header {
  height: 48px;
  background: #e50000;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.flash-products {
  background: #fff;
  overflow-x: auto;
  min-height: 280px;
}

.flash-item {
  width: 170px;
  flex-shrink: 0;
}

.truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.countdown-timer {
  font-family: "Courier New", monospace;
  font-weight: 700;
  letter-spacing: 1px;
}
</style>
