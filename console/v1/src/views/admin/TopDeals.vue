<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12">
        <h1 class="text-h4 mb-4">Top Deals Management</h1>
      </v-col>
    </v-row>

    <!-- Stats Cards -->
    <v-row>
      <v-col cols="12" md="3">
        <v-card color="primary">
          <v-card-text>
            <div class="text-h5">{{ stats.totalDeals }}</div>
            <div class="text-caption">Total Top Deals</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card color="success">
          <v-card-text>
            <div class="text-h5">{{ stats.hotDeals }}</div>
            <div class="text-caption">Hot Deals</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card color="info">
          <v-card-text>
            <div class="text-h5">{{ stats.clearance }}</div>
            <div class="text-caption">Clearance Items</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card color="warning">
          <v-card-text>
            <div class="text-h5">{{ stats.limitedOffer }}</div>
            <div class="text-caption">Limited Offers</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Products Table -->
    <v-row class="mt-4">
      <v-col cols="12">
        <v-card>
          <v-card-title>
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="search"
                  label="Search products"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  density="compact"
                  hide-details
                  clearable
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6" class="text-right">
                <v-btn color="primary" @click="loadProducts">
                  <v-icon left>mdi-refresh</v-icon>
                  Refresh
                </v-btn>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table
            :headers="headers"
            :items="filteredProducts"
            :loading="loading"
            item-value="id"
          >
            <template v-slot:item.image_url="{ item }">
              <v-img
                :src="item.image_url"
                height="60"
                width="60"
                class="ma-2"
              ></v-img>
            </template>

            <template v-slot:item.price="{ item }">
              KSh {{ item.price.toLocaleString() }}
            </template>

            <template v-slot:item.is_top_deal="{ item }">
              <v-chip
                :color="item.is_top_deal ? 'success' : 'default'"
                size="small"
              >
                {{ item.is_top_deal ? "Yes" : "No" }}
              </v-chip>
            </template>

            <template v-slot:item.deal_type="{ item }">
              <v-chip
                v-if="item.deal_type"
                :color="getDealTypeColor(item.deal_type)"
                size="small"
              >
                {{ formatDealType(item.deal_type) }}
              </v-chip>
              <span v-else>-</span>
            </template>

            <template v-slot:item.deal_priority="{ item }">
              <v-chip v-if="item.is_top_deal" size="small">
                {{ item.deal_priority || 0 }}
              </v-chip>
              <span v-else>-</span>
            </template>

            <template v-slot:item.actions="{ item }">
              <v-btn
                icon
                size="small"
                variant="text"
                @click="openDealDialog(item)"
              >
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>

    <!-- Deal Dialog -->
    <v-dialog v-model="dealDialog" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="text-h5">Manage Top Deal</span>
        </v-card-title>

        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <h3>{{ selectedProduct?.name }}</h3>
                <p class="text-caption">{{ selectedProduct?.category }}</p>
              </v-col>

              <v-col cols="12">
                <v-switch
                  v-model="dealForm.is_top_deal"
                  label="Mark as Top Deal"
                  color="primary"
                  hide-details
                ></v-switch>
              </v-col>

              <v-col cols="12" v-if="dealForm.is_top_deal">
                <v-select
                  v-model="dealForm.deal_type"
                  :items="dealTypes"
                  label="Deal Type"
                  variant="outlined"
                  density="compact"
                  :rules="[(v) => !!v || 'Deal type is required']"
                ></v-select>
              </v-col>

              <v-col cols="12" v-if="dealForm.is_top_deal">
                <v-text-field
                  v-model.number="dealForm.deal_priority"
                  label="Deal Priority (higher = displayed first)"
                  type="number"
                  variant="outlined"
                  density="compact"
                  hint="Higher numbers appear first"
                  :rules="[(v) => v >= 0 || 'Priority must be 0 or greater']"
                ></v-text-field>
              </v-col>

              <v-col cols="12">
                <v-alert type="info" variant="tonal" density="compact">
                  <div class="text-caption">
                    <strong>Deal Types:</strong><br />
                    • Hot Deal: Popular products with high demand<br />
                    • Clearance: Products being cleared out<br />
                    • Limited Offer: Time-sensitive special offers<br />
                    • Top Deal: General featured deals
                  </div>
                </v-alert>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn text @click="dealDialog = false">Cancel</v-btn>
          <v-btn color="primary" :loading="saving" @click="saveDeal">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

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
import axios from "axios";

const API_URL = "http://localhost:8080/api/v1";

const products = ref([]);
const loading = ref(false);
const search = ref("");
const dealDialog = ref(false);
const selectedProduct = ref(null);
const saving = ref(false);

const dealForm = ref({
  is_top_deal: false,
  deal_type: "",
  deal_priority: 0,
});

const dealTypes = [
  { title: "Hot Deal", value: "hot_deal" },
  { title: "Clearance", value: "clearance" },
  { title: "Limited Offer", value: "limited_offer" },
  { title: "Top Deal", value: "top_deal" },
];

const snackbar = ref({
  show: false,
  message: "",
  color: "success",
});

const headers = [
  { title: "Image", key: "image_url", sortable: false },
  { title: "Name", key: "name" },
  { title: "Category", key: "category" },
  { title: "Price", key: "price" },
  { title: "Stock", key: "stock" },
  { title: "Is Top Deal", key: "is_top_deal" },
  { title: "Deal Type", key: "deal_type" },
  { title: "Priority", key: "deal_priority" },
  { title: "Actions", key: "actions", sortable: false },
];

const stats = computed(() => {
  const deals = products.value.filter((p) => p.is_top_deal);
  return {
    totalDeals: deals.length,
    hotDeals: deals.filter((p) => p.deal_type === "hot_deal").length,
    clearance: deals.filter((p) => p.deal_type === "clearance").length,
    limitedOffer: deals.filter((p) => p.deal_type === "limited_offer").length,
  };
});

const filteredProducts = computed(() => {
  if (!search.value) return products.value;

  const searchLower = search.value.toLowerCase();
  return products.value.filter(
    (product) =>
      product.name.toLowerCase().includes(searchLower) ||
      product.category.toLowerCase().includes(searchLower)
  );
});

const loadProducts = async () => {
  loading.value = true;
  try {
    const response = await axios.get(`${API_URL}/products`, {
      params: { page: 1, page_size: 100 },
    });
    if (response.data.success) {
      products.value = response.data.products || [];
    }
  } catch (error) {
    console.error("Error loading products:", error);
    showSnackbar("Error loading products", "error");
  } finally {
    loading.value = false;
  }
};

const openDealDialog = (product) => {
  selectedProduct.value = product;
  dealForm.value = {
    is_top_deal: product.is_top_deal || false,
    deal_type: product.deal_type || "",
    deal_priority: product.deal_priority || 0,
  };
  dealDialog.value = true;
};

const saveDeal = async () => {
  if (dealForm.value.is_top_deal && !dealForm.value.deal_type) {
    showSnackbar("Please select a deal type", "error");
    return;
  }

  saving.value = true;
  try {
    const updateData = {
      ...selectedProduct.value,
      is_top_deal: dealForm.value.is_top_deal,
      deal_type: dealForm.value.is_top_deal ? dealForm.value.deal_type : "",
      deal_priority: dealForm.value.is_top_deal
        ? dealForm.value.deal_priority
        : 0,
    };

    const response = await axios.put(
      `${API_URL}/products/${selectedProduct.value.id}`,
      updateData
    );

    if (response.data.success) {
      showSnackbar("Deal updated successfully", "success");
      dealDialog.value = false;
      await loadProducts();
    } else {
      showSnackbar(response.data.message || "Error updating deal", "error");
    }
  } catch (error) {
    console.error("Error saving deal:", error);
    showSnackbar("Error saving deal", "error");
  } finally {
    saving.value = false;
  }
};

const getDealTypeColor = (dealType) => {
  const colors = {
    hot_deal: "error",
    clearance: "warning",
    limited_offer: "info",
    top_deal: "success",
  };
  return colors[dealType] || "default";
};

const formatDealType = (dealType) => {
  return dealType
    .split("_")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(" ");
};

const showSnackbar = (message, color = "success") => {
  snackbar.value = {
    show: true,
    message,
    color,
  };
};

onMounted(() => {
  loadProducts();
});
</script>
