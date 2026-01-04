<template>
  <v-container fluid class="pa-6 mt-7">
    <v-row>
      <v-col cols="12">
        <!-- Header -->
        <div class="d-flex justify-space-between align-center mb-6 mt-4">
          <div>
            <h1 class="text-h4 font-weight-bold">Flash Sales Management</h1>
            <p class="text-body-2 text-grey-darken-1 mt-1">
              Select products and set timelines for flash sales
            </p>
          </div>
        </div>

        <!-- Success/Error Alerts -->
        <v-alert
          v-if="successMessage"
          type="success"
          variant="tonal"
          closable
          class="mb-4"
          @click:close="successMessage = ''"
        >
          {{ successMessage }}
        </v-alert>

        <v-alert
          v-if="error"
          type="error"
          variant="tonal"
          closable
          class="mb-4"
          @click:close="error = ''"
        >
          {{ error }}
        </v-alert>

        <!-- Stats Cards -->
        <v-row class="mb-6">
          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="red-lighten-4" class="mr-4">
                    <v-icon color="red">mdi-flash</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-caption text-grey">Active Flash Sales</div>
                    <div class="text-h5 font-weight-bold">
                      {{ activeFlashSalesCount }}
                    </div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>

          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="orange-lighten-4" class="mr-4">
                    <v-icon color="orange">mdi-clock-outline</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-caption text-grey">Ending Soon</div>
                    <div class="text-h5 font-weight-bold">
                      {{ endingSoonCount }}
                    </div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>

          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="green-lighten-4" class="mr-4">
                    <v-icon color="green">mdi-package-variant</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-caption text-grey">Available Products</div>
                    <div class="text-h5 font-weight-bold">
                      {{ products.length }}
                    </div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>

          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="blue-lighten-4" class="mr-4">
                    <v-icon color="blue">mdi-percent</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-caption text-grey">Avg. Discount</div>
                    <div class="text-h5 font-weight-bold">
                      {{ avgDiscount }}%
                    </div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>

        <!-- Products Table -->
        <v-card>
          <v-card-title class="d-flex align-center py-4">
            <v-icon class="mr-2">mdi-flash</v-icon>
            Products
            <v-spacer />
            <v-text-field
              v-model="search"
              density="compact"
              placeholder="Search products..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              hide-details
              single-line
              class="mr-4"
              style="max-width: 300px"
            />
            <v-btn color="orange" variant="flat" @click="loadProducts">
              <v-icon class="mr-1">mdi-refresh</v-icon>
              Refresh
            </v-btn>
          </v-card-title>

          <v-data-table
            :headers="headers"
            :items="filteredProducts"
            :loading="loading"
            :items-per-page="10"
          >
            <!-- Image -->
            <template #item.image_url="{ item }">
              <v-avatar size="50" rounded="lg" class="my-2">
                <v-img :src="item.image_url" cover />
              </v-avatar>
            </template>

            <!-- Flash Sale Status -->
            <template #item.is_flash_sale="{ item }">
              <v-chip
                :color="item.is_flash_sale ? 'red' : 'grey'"
                variant="flat"
                size="small"
              >
                <v-icon start>
                  {{ item.is_flash_sale ? "mdi-flash" : "mdi-flash-off" }}
                </v-icon>
                {{ item.is_flash_sale ? "Active" : "Inactive" }}
              </v-chip>
            </template>

            <!-- Price Display -->
            <template #item.price="{ item }">
              <div>
                <div class="font-weight-bold">
                  KSh {{ formatPrice(item.price) }}
                </div>
                <div v-if="item.is_flash_sale" class="text-caption text-red">
                  Sale: KSh {{ formatPrice(item.flash_sale_price) }}
                </div>
              </div>
            </template>

            <!-- End Time -->
            <template #item.flash_sale_end_time="{ item }">
              <div v-if="item.is_flash_sale && item.flash_sale_end_time">
                <div class="text-caption">
                  {{ formatDateTime(item.flash_sale_end_time) }}
                </div>
                <div class="text-caption text-grey">
                  {{ getTimeRemaining(item.flash_sale_end_time) }}
                </div>
              </div>
              <span v-else class="text-grey">—</span>
            </template>

            <!-- Actions -->
            <template #item.actions="{ item }">
              <v-btn
                :color="item.is_flash_sale ? 'grey' : 'red'"
                size="small"
                variant="tonal"
                @click="openFlashSaleDialog(item)"
              >
                <v-icon start>
                  {{ item.is_flash_sale ? "mdi-pencil" : "mdi-flash" }}
                </v-icon>
                {{ item.is_flash_sale ? "Edit" : "Add to Flash Sale" }}
              </v-btn>
              <v-btn
                v-if="item.is_flash_sale"
                color="red"
                size="small"
                variant="text"
                class="ml-2"
                @click="removeFromFlashSale(item)"
              >
                Remove
              </v-btn>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>

    <!-- Flash Sale Dialog -->
    <v-dialog v-model="flashSaleDialog" max-width="600">
      <v-card>
        <v-card-title class="bg-red-lighten-5 d-flex align-center">
          <v-icon class="mr-2">mdi-flash</v-icon>
          {{ selectedProduct?.is_flash_sale ? "Edit" : "Add to" }} Flash Sale
        </v-card-title>

        <v-card-text class="pt-6">
          <v-form ref="flashSaleForm">
            <!-- Product Info -->
            <div class="mb-4 d-flex align-center">
              <v-avatar size="60" rounded="lg" class="mr-3">
                <v-img :src="selectedProduct?.image_url" cover />
              </v-avatar>
              <div>
                <div class="font-weight-bold">{{ selectedProduct?.name }}</div>
                <div class="text-caption text-grey">
                  {{ selectedProduct?.category }}
                </div>
              </div>
            </div>

            <v-divider class="mb-4" />

            <v-row>
              <!-- Original Price (Read-only) -->
              <v-col cols="12" md="6">
                <v-text-field
                  :model-value="formatPrice(selectedProduct?.price)"
                  label="Original Price (KSh)"
                  variant="outlined"
                  readonly
                  prepend-inner-icon="mdi-currency-usd"
                />
              </v-col>

              <!-- Flash Sale Price -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="flashSaleData.flash_sale_price"
                  label="Flash Sale Price (KSh) *"
                  variant="outlined"
                  type="number"
                  step="0.01"
                  :rules="flashSalePriceRules"
                  prepend-inner-icon="mdi-tag"
                  required
                />
              </v-col>

              <!-- Discount Percentage (Calculated) -->
              <v-col cols="12">
                <v-alert type="info" variant="tonal" density="compact">
                  <div class="d-flex align-center">
                    <v-icon start>mdi-percent</v-icon>
                    <strong>{{ calculatedDiscount }}% OFF</strong>
                    <span class="ml-2 text-grey">
                      (Save KSh {{ formatPrice(calculatedSavings) }})
                    </span>
                  </div>
                </v-alert>
              </v-col>

              <!-- End Date & Time -->
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="flashSaleData.end_date"
                  label="End Date *"
                  variant="outlined"
                  type="date"
                  :rules="[(v) => !!v || 'End date is required']"
                  prepend-inner-icon="mdi-calendar"
                  :min="todayDate"
                  required
                />
              </v-col>

              <v-col cols="12" md="6">
                <v-text-field
                  v-model="flashSaleData.end_time"
                  label="End Time *"
                  variant="outlined"
                  type="time"
                  :rules="[(v) => !!v || 'End time is required']"
                  prepend-inner-icon="mdi-clock-outline"
                  required
                />
              </v-col>

              <!-- Initial Stock (for progress tracking) -->
              <v-col cols="12">
                <v-text-field
                  v-model.number="flashSaleData.initial_stock"
                  label="Flash Sale Stock Quantity *"
                  variant="outlined"
                  type="number"
                  :rules="stockRules"
                  prepend-inner-icon="mdi-package-variant"
                  hint="This will be used to track flash sale progress"
                  persistent-hint
                  required
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions>
          <v-spacer />
          <v-btn
            variant="text"
            @click="flashSaleDialog = false"
            :disabled="saving"
          >
            Cancel
          </v-btn>
          <v-btn
            color="red"
            variant="flat"
            :loading="saving"
            @click="saveFlashSale"
          >
            <v-icon start>mdi-flash</v-icon>
            Save Flash Sale
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from "vue";

const API_BASE = "http://localhost:8080/api/v1";

// State
const loading = ref(false);
const saving = ref(false);
const error = ref("");
const successMessage = ref("");
const products = ref([]);
const search = ref("");
const flashSaleDialog = ref(false);
const selectedProduct = ref(null);
const flashSaleForm = ref(null);

// Flash Sale Data
const flashSaleData = reactive({
  flash_sale_price: 0,
  end_date: "",
  end_time: "",
  initial_stock: 0,
});

// Table Headers
const headers = [
  { title: "Image", key: "image_url", sortable: false },
  { title: "Product", key: "name" },
  { title: "Category", key: "category" },
  { title: "Price", key: "price" },
  { title: "Stock", key: "stock" },
  { title: "Flash Sale", key: "is_flash_sale" },
  { title: "Ends At", key: "flash_sale_end_time" },
  { title: "Actions", key: "actions", sortable: false },
];

// Validation Rules
const flashSalePriceRules = [
  (v) => (v !== null && v !== "") || "Flash sale price is required",
  (v) => v > 0 || "Price must be greater than 0",
  (v) =>
    !selectedProduct.value ||
    v < selectedProduct.value.price ||
    "Flash sale price must be less than original price",
];

const stockRules = [
  (v) => (v !== null && v !== "") || "Stock is required",
  (v) => v > 0 || "Stock must be greater than 0",
  (v) =>
    !selectedProduct.value ||
    v <= selectedProduct.value.stock ||
    "Flash sale stock cannot exceed available stock",
];

// Computed
const todayDate = computed(() => {
  const today = new Date();
  return today.toISOString().split("T")[0];
});

const filteredProducts = computed(() => {
  if (!search.value) return products.value;
  const query = search.value.toLowerCase();
  return products.value.filter(
    (p) =>
      p.name.toLowerCase().includes(query) ||
      p.category.toLowerCase().includes(query) ||
      p.brand?.toLowerCase().includes(query)
  );
});

const activeFlashSalesCount = computed(() => {
  return products.value.filter(
    (p) => p.is_flash_sale && isFlashSaleActive(p.flash_sale_end_time)
  ).length;
});

const endingSoonCount = computed(() => {
  const now = new Date();
  const sixHours = 6 * 60 * 60 * 1000;
  return products.value.filter((p) => {
    if (!p.is_flash_sale || !p.flash_sale_end_time) return false;
    const endTime = new Date(p.flash_sale_end_time);
    const diff = endTime - now;
    return diff > 0 && diff <= sixHours;
  }).length;
});

const avgDiscount = computed(() => {
  const flashSales = products.value.filter(
    (p) => p.is_flash_sale && p.flash_sale_price
  );
  if (flashSales.length === 0) return 0;
  const totalDiscount = flashSales.reduce((sum, p) => {
    const discount = ((p.price - p.flash_sale_price) / p.price) * 100;
    return sum + discount;
  }, 0);
  return Math.round(totalDiscount / flashSales.length);
});

const calculatedDiscount = computed(() => {
  if (!selectedProduct.value || !flashSaleData.flash_sale_price) return 0;
  const discount =
    ((selectedProduct.value.price - flashSaleData.flash_sale_price) /
      selectedProduct.value.price) *
    100;
  return Math.round(discount);
});

const calculatedSavings = computed(() => {
  if (!selectedProduct.value || !flashSaleData.flash_sale_price) return 0;
  return selectedProduct.value.price - flashSaleData.flash_sale_price;
});

// Methods
onMounted(() => {
  loadProducts();
});

async function loadProducts() {
  loading.value = true;
  error.value = "";

  try {
    const response = await fetch(`${API_BASE}/products?page=1&page_size=1000`);

    if (!response.ok) {
      throw new Error("Failed to load products");
    }

    const data = await response.json();
    products.value = data.products || [];
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
}

function openFlashSaleDialog(product) {
  selectedProduct.value = product;
  flashSaleDialog.value = true;

  // Pre-fill if editing existing flash sale
  if (product.is_flash_sale) {
    flashSaleData.flash_sale_price = product.flash_sale_price || 0;
    flashSaleData.initial_stock = product.initial_stock || product.stock;

    if (product.flash_sale_end_time) {
      const endDate = new Date(product.flash_sale_end_time);
      flashSaleData.end_date = endDate.toISOString().split("T")[0];
      flashSaleData.end_time = endDate.toTimeString().substring(0, 5);
    }
  } else {
    // Reset for new flash sale
    flashSaleData.flash_sale_price = Math.round(product.price * 0.7); // Suggest 30% off
    flashSaleData.initial_stock = product.stock;
    flashSaleData.end_date = "";
    flashSaleData.end_time = "23:59";
  }
}

async function saveFlashSale() {
  const { valid } = await flashSaleForm.value.validate();
  if (!valid) return;

  saving.value = true;
  error.value = "";

  try {
    // Combine date and time
    const endDateTime = new Date(
      `${flashSaleData.end_date}T${flashSaleData.end_time}`
    );

    const updateData = {
      ...selectedProduct.value,
      is_flash_sale: true,
      flash_sale_price: flashSaleData.flash_sale_price,
      flash_sale_end_time: endDateTime.toISOString(),
      initial_stock: flashSaleData.initial_stock,
      stock: flashSaleData.initial_stock, // Update stock to match
    };

    const response = await fetch(
      `${API_BASE}/products/${selectedProduct.value.id}`,
      {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updateData),
      }
    );

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || "Failed to save flash sale");
    }

    successMessage.value = `Flash sale added to "${selectedProduct.value.name}"`;
    flashSaleDialog.value = false;
    await loadProducts();

    setTimeout(() => {
      successMessage.value = "";
    }, 5000);
  } catch (err) {
    error.value = err.message;
  } finally {
    saving.value = false;
  }
}

async function removeFromFlashSale(product) {
  if (!confirm(`Remove "${product.name}" from flash sale?`)) return;

  saving.value = true;
  error.value = "";

  try {
    const updateData = {
      ...product,
      is_flash_sale: false,
      flash_sale_price: 0,
      flash_sale_end_time: null,
    };

    const response = await fetch(`${API_BASE}/products/${product.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(updateData),
    });

    if (!response.ok) {
      throw new Error("Failed to remove flash sale");
    }

    successMessage.value = `Removed "${product.name}" from flash sale`;
    await loadProducts();

    setTimeout(() => {
      successMessage.value = "";
    }, 3000);
  } catch (err) {
    error.value = err.message;
  } finally {
    saving.value = false;
  }
}

function formatPrice(price) {
  if (!price) return "0.00";
  return Number(price).toLocaleString("en-KE", {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  });
}

function formatDateTime(dateString) {
  if (!dateString) return "";
  const date = new Date(dateString);
  return date.toLocaleDateString("en-KE", {
    month: "short",
    day: "numeric",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

function getTimeRemaining(endTime) {
  if (!endTime) return "";
  const now = new Date();
  const end = new Date(endTime);
  const diff = end - now;

  if (diff <= 0) return "Ended";

  const hours = Math.floor(diff / (1000 * 60 * 60));
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));

  if (hours > 24) {
    const days = Math.floor(hours / 24);
    return `${days}d ${hours % 24}h left`;
  }
  return `${hours}h ${minutes}m left`;
}

function isFlashSaleActive(endTime) {
  if (!endTime) return false;
  return new Date(endTime) > new Date();
}
</script>

<style scoped>
.v-data-table {
  font-size: 14px;
}
</style>
