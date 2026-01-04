<template>
  <v-container fluid class="pa-6 mt-7 ">
    <v-row>
      <v-col cols="12">
        <!-- Header -->
        <div class="d-flex justify-space-between align-center mb-6 mt-4">
          <div>
            <h1 class="text-h4 font-weight-bold">Products Management</h1>
            <p class="text-body-2 text-grey-darken-1 mt-1">
              Manage all products in the catalog
            </p>
          </div>
          <v-btn
            color="orange"
            size="large"
            prepend-icon="mdi-plus"
            to="/admin/products/add"
          >
            Add New Product
          </v-btn>
        </div>

        <!-- Stats Cards -->
        <v-row class="mb-6">
          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="orange-lighten-4" class="mr-4">
                    <v-icon color="orange">mdi-package-variant</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-h5 font-weight-bold">
                      {{ productStore.totalProducts }}
                    </div>
                    <div class="text-body-2 text-grey-darken-1">
                      Total Products
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
                    <v-icon color="green">mdi-check-circle</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-h5 font-weight-bold">
                      {{ inStockCount }}
                    </div>
                    <div class="text-body-2 text-grey-darken-1">In Stock</div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>

          <v-col cols="12" sm="6" md="3">
            <v-card>
              <v-card-text>
                <div class="d-flex align-center">
                  <v-avatar color="red-lighten-4" class="mr-4">
                    <v-icon color="red">mdi-alert-circle</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-h5 font-weight-bold">
                      {{ lowStockCount }}
                    </div>
                    <div class="text-body-2 text-grey-darken-1">Low Stock</div>
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
                    <v-icon color="blue">mdi-sale</v-icon>
                  </v-avatar>
                  <div>
                    <div class="text-h5 font-weight-bold">
                      {{ onSaleCount }}
                    </div>
                    <div class="text-body-2 text-grey-darken-1">On Sale</div>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>

        <!-- Filters & Search -->
        <v-card class="mb-6">
          <v-card-text>
            <v-row>
              <v-col cols="12" md="4">
                <v-text-field
                  v-model="searchQuery"
                  prepend-inner-icon="mdi-magnify"
                  label="Search products..."
                  variant="outlined"
                  density="compact"
                  hide-details
                  clearable
                  @input="handleSearch"
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="selectedCategory"
                  :items="categories"
                  label="Category"
                  variant="outlined"
                  density="compact"
                  hide-details
                  clearable
                  @update:model-value="fetchProducts"
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="stockFilter"
                  :items="stockFilters"
                  label="Stock Status"
                  variant="outlined"
                  density="compact"
                  hide-details
                  clearable
                  @update:model-value="applyFilters"
                />
              </v-col>
              <v-col cols="12" md="2">
                <v-btn
                  block
                  color="grey-darken-1"
                  variant="outlined"
                  @click="resetFilters"
                >
                  Reset Filters
                </v-btn>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- Products Table -->
        <v-card>
          <!-- Loading State -->
          <div v-if="productStore.loading" class="text-center py-12">
            <v-progress-circular indeterminate color="orange" size="64" />
            <p class="text-body-1 mt-4">Loading products...</p>
          </div>

          <!-- Error State -->
          <v-alert
            v-else-if="productStore.error"
            type="error"
            variant="tonal"
            class="ma-4"
          >
            {{ productStore.error }}
          </v-alert>

          <!-- Products Table -->
          <v-data-table
            v-else
            :headers="headers"
            :items="filteredProducts"
            :items-per-page="10"
            class="elevation-0"
          >
            <!-- Product Column -->
            <template #[`item.product`]="{ item }">
              <div class="d-flex align-center py-2">
                <v-img
                  :src="item.image_url || '/images/placeholder.jpg'"
                  :alt="item.name"
                  width="60"
                  height="60"
                  cover
                  class="rounded mr-3"
                />
                <div>
                  <div class="font-weight-medium">{{ item.name }}</div>
                  <div class="text-caption text-grey-darken-1">
                    {{ item.brand || "No brand" }}
                  </div>
                </div>
              </div>
            </template>

            <!-- Category -->
            <template #[`item.category`]="{ item }">
              <v-chip size="small" color="primary" variant="tonal">
                {{ item.category }}
              </v-chip>
            </template>

            <!-- Price -->
            <template #[`item.price`]="{ item }">
              <div>
                <div class="font-weight-medium">
                  KSh {{ formatPrice(item.price) }}
                </div>
                <div
                  v-if="item.discount_percentage > 0"
                  class="text-caption text-green"
                >
                  {{ item.discount_percentage }}% off
                </div>
              </div>
            </template>

            <!-- Stock -->
            <template #[`item.stock`]="{ item }">
              <v-chip
                :color="getStockColor(item.stock)"
                size="small"
                variant="flat"
              >
                {{ item.stock }} units
              </v-chip>
            </template>

            <!-- Status -->
            <template #[`item.status`]="{ item }">
              <v-chip
                :color="item.stock > 0 ? 'success' : 'error'"
                size="small"
                variant="tonal"
              >
                {{ item.stock > 0 ? "Active" : "Out of Stock" }}
              </v-chip>
            </template>

            <!-- Actions -->
            <template #[`item.actions`]="{ item }">
              <div class="d-flex ga-2">
                <v-btn
                  icon="mdi-eye"
                  size="small"
                  variant="tonal"
                  color="blue"
                  @click="viewProduct(item)"
                />
                <v-btn
                  icon="mdi-pencil"
                  size="small"
                  variant="tonal"
                  color="orange"
                  @click="editProduct(item)"
                />
                <v-btn
                  icon="mdi-delete"
                  size="small"
                  variant="tonal"
                  color="red"
                  @click="confirmDelete(item)"
                />
              </div>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="deleteDialog" max-width="500">
      <v-card>
        <v-card-title class="text-h5">Delete Product?</v-card-title>
        <v-card-text>
          Are you sure you want to delete
          <strong>{{ productToDelete?.name }}</strong
          >? This action cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="deleteDialog = false">Cancel</v-btn>
          <v-btn
            color="red"
            variant="flat"
            :loading="deleting"
            @click="deleteProduct"
          >
            Delete
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useProductStore } from "@/stores/products";

const router = useRouter();
const productStore = useProductStore();

const searchQuery = ref("");
const selectedCategory = ref(null);
const stockFilter = ref(null);
const deleteDialog = ref(false);
const productToDelete = ref(null);
const deleting = ref(false);

const headers = [
  { title: "Product", key: "product", sortable: false },
  { title: "Category", key: "category" },
  { title: "Price", key: "price" },
  { title: "Stock", key: "stock" },
  { title: "Status", key: "status" },
  { title: "Actions", key: "actions", sortable: false },
];

const categories = [
  "Electronics",
  "Smartphones",
  "Laptops",
  "Fashion",
  "Home & Kitchen",
  "Sports",
  "Beauty",
  "Books",
];

const stockFilters = [
  { title: "All", value: null },
  { title: "In Stock", value: "in_stock" },
  { title: "Low Stock (< 10)", value: "low_stock" },
  { title: "Out of Stock", value: "out_of_stock" },
];

// Computed
const filteredProducts = computed(() => {
  let products = productStore.products;

  // Apply stock filter
  if (stockFilter.value === "in_stock") {
    products = products.filter((p) => p.stock > 10);
  } else if (stockFilter.value === "low_stock") {
    products = products.filter((p) => p.stock > 0 && p.stock <= 10);
  } else if (stockFilter.value === "out_of_stock") {
    products = products.filter((p) => p.stock === 0);
  }

  return products;
});

const inStockCount = computed(
  () => productStore.products.filter((p) => p.stock > 0).length
);

const lowStockCount = computed(
  () => productStore.products.filter((p) => p.stock > 0 && p.stock <= 10).length
);

const onSaleCount = computed(
  () => productStore.products.filter((p) => p.discount_percentage > 0).length
);

// Methods
async function fetchProducts() {
  if (selectedCategory.value) {
    await productStore.getProductsByCategory(selectedCategory.value);
  } else {
    await productStore.fetchProducts(1, 100); // Get more products for admin
  }
}

function handleSearch() {
  if (searchQuery.value && searchQuery.value.trim()) {
    productStore.searchProducts(searchQuery.value);
  } else {
    fetchProducts();
  }
}

function resetFilters() {
  searchQuery.value = "";
  selectedCategory.value = null;
  stockFilter.value = null;
  fetchProducts();
}

function applyFilters() {
  // Filters are reactive and computed automatically
}

function viewProduct(product) {
  // Navigate to product detail page
  router.push(`/product/${product.id}`);
}

function editProduct(product) {
  // Navigate to edit page (you can create this later)
  router.push(`/admin/products/edit/${product.id}`);
}

function confirmDelete(product) {
  productToDelete.value = product;
  deleteDialog.value = true;
}

async function deleteProduct() {
  deleting.value = true;
  try {
    const response = await fetch(
      `http://localhost:8080/api/v1/products/${productToDelete.value.id}`,
      { method: "DELETE" }
    );

    if (response.ok) {
      // Refresh products list
      await fetchProducts();
      deleteDialog.value = false;
      productToDelete.value = null;
    } else {
      alert("Failed to delete product");
    }
  } catch (err) {
    console.error("Delete error:", err);
    alert("Failed to delete product");
  } finally {
    deleting.value = false;
  }
}

function getStockColor(stock) {
  if (stock === 0) return "red";
  if (stock <= 10) return "orange";
  return "green";
}

function formatPrice(price) {
  return new Intl.NumberFormat("en-KE", {
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  }).format(price);
}

// Lifecycle
onMounted(() => {
  fetchProducts();
});
</script>

<style scoped>
.v-data-table {
  background-color: transparent;
}
</style>
