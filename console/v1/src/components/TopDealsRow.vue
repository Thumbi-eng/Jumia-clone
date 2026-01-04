<template>
  <div>
    <!-- Hot Deals Section -->
    <v-card
      v-if="dealsByType.hot_deal.length > 0"
      elevation="0"
      rounded="lg"
      class="mb-4"
    >
      <div
        class="d-flex align-center justify-space-between px-4"
        style="height: 48px; background: #d32f2f; min-width: 1540px;"
      >
        <div class="font-weight-bold text-white">🔥 Hot Deals</div>
        <v-btn variant="text" color="white" size="small" class="text-none">
          See All →
        </v-btn>
      </div>

      <div class="d-flex flex-wrap pa-3">
        <v-card
          v-for="deal in dealsByType.hot_deal"
          :key="deal.id"
          elevation="0"
          class="pa-2 mr-2 mb-2"
          style="width: 180px; cursor: pointer"
          @click="goToProduct(deal.id)"
        >
          <div class="position-relative">
            <v-img
              :src="deal.image_url || '/images/products/placeholder.jpg'"
              aspect-ratio="1"
              cover
              class="rounded"
            />
            <v-chip
              v-if="deal.discount_percentage > 0"
              size="small"
              color="orange-lighten-5"
              class="position-absolute top-0 right-0 ma-1"
              text-color="orange-darken-2"
              label
            >
              -{{ Math.round(deal.discount_percentage) }}%
            </v-chip>
          </div>
          <div
            class="text-body-2 mt-2"
            style="line-height: 1.2; height: 34px; overflow: hidden"
          >
            {{ deal.name }}
          </div>
          <div class="font-weight-bold mt-1">
            KSh {{ deal.final_price.toLocaleString() }}
          </div>
          <div
            v-if="deal.discount_percentage > 0"
            class="text-caption text-grey text-decoration-line-through"
          >
            KSh {{ deal.price.toLocaleString() }}
          </div>
        </v-card>
      </div>
    </v-card>

    <!-- Clearance Section -->
    <v-card
      v-if="dealsByType.clearance.length > 0"
      elevation="0"
      rounded="lg"
      class="mb-4"
    >
      <div
        class="d-flex align-center justify-space-between px-4"
        style="height: 48px; background: #f57c00"
      >
        <div class="font-weight-bold text-white">🏷️ Clearance Sale</div>
        <v-btn variant="text" color="white" size="small" class="text-none">
          See All →
        </v-btn>
      </div>

      <div class="d-flex flex-wrap pa-3">
        <v-card
          v-for="deal in dealsByType.clearance"
          :key="deal.id"
          elevation="0"
          class="pa-2 mr-2 mb-2"
          style="width: 180px; cursor: pointer"
          @click="goToProduct(deal.id)"
        >
          <div class="position-relative">
            <v-img
              :src="deal.image_url || '/images/products/placeholder.jpg'"
              aspect-ratio="1"
              cover
              class="rounded"
            />
            <v-chip
              v-if="deal.discount_percentage > 0"
              size="small"
              color="orange-lighten-5"
              class="position-absolute top-0 right-0 ma-1"
              text-color="orange-darken-2"
              label
            >
              -{{ Math.round(deal.discount_percentage) }}%
            </v-chip>
          </div>
          <div
            class="text-body-2 mt-2"
            style="line-height: 1.2; height: 34px; overflow: hidden"
          >
            {{ deal.name }}
          </div>
          <div class="font-weight-bold mt-1">
            KSh {{ deal.final_price.toLocaleString() }}
          </div>
          <div
            v-if="deal.discount_percentage > 0"
            class="text-caption text-grey text-decoration-line-through"
          >
            KSh {{ deal.price.toLocaleString() }}
          </div>
        </v-card>
      </div>
    </v-card>

    <!-- Limited Offer Section -->
    <v-card
      v-if="dealsByType.limited_offer.length > 0"
      elevation="0"
      rounded="lg"
      class="mb-4"
    >
      <div
        class="d-flex align-center justify-space-between px-4"
        style="height: 48px; background: #0288d1"
      >
        <div class="font-weight-bold text-white">⏰ Limited Time Offers</div>
        <v-btn variant="text" color="white" size="small" class="text-none">
          See All →
        </v-btn>
      </div>

      <div class="d-flex flex-wrap pa-3">
        <v-card
          v-for="deal in dealsByType.limited_offer"
          :key="deal.id"
          elevation="0"
          class="pa-2 mr-2 mb-2"
          style="width: 180px; cursor: pointer"
          @click="goToProduct(deal.id)"
        >
          <div class="position-relative">
            <v-img
              :src="deal.image_url || '/images/products/placeholder.jpg'"
              aspect-ratio="1"
              cover
              class="rounded"
            />
            <v-chip
              v-if="deal.discount_percentage > 0"
              size="small"
              color="orange-lighten-5"
              class="position-absolute top-0 right-0 ma-1"
              text-color="orange-darken-2"
              label
            >
              -{{ Math.round(deal.discount_percentage) }}%
            </v-chip>
          </div>
          <div
            class="text-body-2 mt-2"
            style="line-height: 1.2; height: 34px; overflow: hidden"
          >
            {{ deal.name }}
          </div>
          <div class="font-weight-bold mt-1">
            KSh {{ deal.final_price.toLocaleString() }}
          </div>
          <div
            v-if="deal.discount_percentage > 0"
            class="text-caption text-grey text-decoration-line-through"
          >
            KSh {{ deal.price.toLocaleString() }}
          </div>
        </v-card>
      </div>
    </v-card>

    <!-- Top Deals Section -->
    <v-card
      v-if="dealsByType.top_deal.length > 0"
      elevation="0"
      rounded="lg"
      class="mb-4"
    >
      <div
        class="d-flex align-center justify-space-between px-4"
        style="height: 48px; background: #388e3c"
      >
        <div class="font-weight-bold text-white">
          ⭐ Krisi Na Jumia | Top Deals
        </div>
        <v-btn variant="text" color="white" size="small" class="text-none">
          See All →
        </v-btn>
      </div>

      <div class="d-flex flex-wrap pa-3">
        <v-card
          v-for="deal in dealsByType.top_deal"
          :key="deal.id"
          elevation="0"
          class="pa-2 mr-2 mb-2"
          style="width: 180px; cursor: pointer"
          @click="goToProduct(deal.id)"
        >
          <div class="position-relative">
            <v-img
              :src="deal.image_url || '/images/products/placeholder.jpg'"
              aspect-ratio="1"
              cover
              class="rounded"
            />
            <v-chip
              v-if="deal.discount_percentage > 0"
              size="small"
              color="orange-lighten-5"
              class="position-absolute top-0 right-0 ma-1"
              text-color="orange-darken-2"
              label
            >
              -{{ Math.round(deal.discount_percentage) }}%
            </v-chip>
          </div>
          <div
            class="text-body-2 mt-2"
            style="line-height: 1.2; height: 34px; overflow: hidden"
          >
            {{ deal.name }}
          </div>
          <div class="font-weight-bold mt-1">
            KSh {{ deal.final_price.toLocaleString() }}
          </div>
          <div
            v-if="deal.discount_percentage > 0"
            class="text-caption text-grey text-decoration-line-through"
          >
            KSh {{ deal.price.toLocaleString() }}
          </div>
        </v-card>
      </div>
    </v-card>

    <!-- Loading State -->
    <v-card v-if="loading" elevation="0" rounded="lg">
      <v-skeleton-loader
        type="card, divider, card-avatar, card-avatar"
      ></v-skeleton-loader>
    </v-card>

    <!-- Empty State -->
    <v-alert
      v-if="
        !loading && Object.values(dealsByType).every((arr) => arr.length === 0)
      "
      type="info"
      variant="tonal"
    >
      No top deals available at the moment. Check back later!
    </v-alert>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();
const API_URL = "http://localhost:8080/api/v1";

const rawDeals = ref([]);
const loading = ref(false);

// Group deals by type and sort by priority within each type
const dealsByType = computed(() => {
  const grouped = {
    hot_deal: [],
    clearance: [],
    limited_offer: [],
    top_deal: [],
  };

  rawDeals.value.forEach((deal) => {
    const dealType = deal.deal_type || "top_deal";
    if (grouped[dealType]) {
      grouped[dealType].push(deal);
    }
  });

  // Sort each group by priority (1, 2, 3, 4...)
  Object.keys(grouped).forEach((type) => {
    grouped[type].sort((a, b) => {
      const priorityA = a.deal_priority || 999;
      const priorityB = b.deal_priority || 999;
      return priorityA - priorityB;
    });
  });

  return grouped;
});

const loadTopDeals = async () => {
  loading.value = true;
  try {
    const response = await axios.get(`${API_URL}/products/top-deals`, {
      params: {
        page: 1,
        page_size: 50,
      },
    });

    if (response.data.success) {
      rawDeals.value = response.data.products || [];
    }
  } catch (error) {
    console.error("Error loading top deals:", error);
  } finally {
    loading.value = false;
  }
};

const goToProduct = (productId) => {
  // Navigate to product detail page
  // router.push(`/product/${productId}`)
  console.log("Navigate to product:", productId);
};

onMounted(() => {
  loadTopDeals();
});
</script>
