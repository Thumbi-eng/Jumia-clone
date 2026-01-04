<template>
  <v-container class="pa-6 mt-7" style="max-width: 1200px">
    <!-- Page Header -->
    <div class="d-flex justify-space-between align-center mb-6 mt-3">
      <div>
        <h1 class="text-h4 font-weight-bold mb-2">Product Management</h1>
        <p class="text-body-1 text-grey-darken-1">
          Add and manage products in your store
        </p>
      </div>
      <v-btn
        color="orange"
        variant="flat"
        prepend-icon="mdi-format-list-bulleted"
        @click="showProductsList = !showProductsList"
      >
        {{ showProductsList ? "Add Product" : "View All Products" }}
      </v-btn>
    </div>

    <!-- Add Product Form -->
    <v-card v-show="!showProductsList" elevation="2" class="mb-6">
      <v-card-title class="bg-orange-lighten-5 py-4">
        <v-icon class="mr-2">mdi-package-variant-plus</v-icon>
        Add New Product
      </v-card-title>

      <v-card-text class="pa-6">
        <!-- Success Alert -->
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

        <!-- Error Alert -->
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

        <v-form ref="form" @submit.prevent="handleAddProduct">
          <v-row>
            <!-- Product Name -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="productData.name"
                label="Product Name *"
                variant="outlined"
                prepend-inner-icon="mdi-tag-outline"
                :rules="nameRules"
                :disabled="loading"
                required
              />
            </v-col>

            <!-- Category -->
            <v-col cols="12" md="6">
              <v-select
                v-model="productData.category"
                label="Category *"
                variant="outlined"
                prepend-inner-icon="mdi-shape-outline"
                :items="categories"
                :rules="categoryRules"
                :disabled="loading"
                required
              />
            </v-col>

            <!-- Brand -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="productData.brand"
                label="Brand"
                variant="outlined"
                prepend-inner-icon="mdi-certificate-outline"
                :disabled="loading"
              />
            </v-col>

            <!-- Price -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model.number="productData.price"
                label="Price (KSh) *"
                variant="outlined"
                type="number"
                step="0.01"
                min="0"
                prepend-inner-icon="mdi-currency-usd"
                :rules="priceRules"
                :disabled="loading"
                required
              />
            </v-col>

            <!-- Stock -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model.number="productData.stock"
                label="Stock Quantity *"
                variant="outlined"
                type="number"
                min="0"
                prepend-inner-icon="mdi-warehouse"
                :rules="stockRules"
                :disabled="loading"
                required
              />
            </v-col>

            <!-- Discount Percentage -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model.number="productData.discount_percentage"
                label="Discount Percentage"
                variant="outlined"
                type="number"
                min="0"
                max="100"
                step="1"
                prepend-inner-icon="mdi-percent-outline"
                :disabled="loading"
                hint="0-100%"
                persistent-hint
              />
            </v-col>

            <!-- Image Upload -->
            <v-col cols="12">
              <v-card variant="outlined" class="pa-4">
                <div class="text-subtitle-2 mb-3 d-flex align-center">
                  <v-icon class="mr-2">mdi-image-outline</v-icon>
                  Product Image
                </div>

                <!-- Image Preview -->
                <div v-if="imagePreview || productData.image_url" class="mb-4">
                  <v-img
                    :src="imagePreview || productData.image_url"
                    max-height="300"
                    max-width="400"
                    class="rounded"
                    cover
                  />
                  <v-btn
                    v-if="imagePreview"
                    color="red"
                    variant="text"
                    size="small"
                    class="mt-2"
                    @click="removeImage"
                  >
                    Remove Image
                  </v-btn>
                </div>

                <!-- Upload Button -->
                <v-file-input
                  v-model="imageFile"
                  label="Upload Product Image"
                  variant="outlined"
                  prepend-icon=""
                  prepend-inner-icon="mdi-camera"
                  accept="image/jpeg,image/jpg,image/png,image/webp"
                  :disabled="loading || uploading"
                  :loading="uploading"
                  show-size
                  @change="handleImageSelect"
                >
                  <template #append>
                    <v-btn
                      v-if="imageFile"
                      color="orange"
                      variant="flat"
                      :loading="uploading"
                      @click="uploadImage"
                    >
                      Upload
                    </v-btn>
                  </template>
                </v-file-input>

                <!-- Upload Progress -->
                <v-progress-linear
                  v-if="uploading"
                  :model-value="uploadProgress"
                  color="orange"
                  height="6"
                  class="mt-2"
                />

                <div class="text-caption text-grey-darken-1 mt-2">
                  Supported formats: JPEG, PNG, WebP (Max 5MB)
                </div>

                <!-- Or use URL -->
                <v-divider class="my-4" />
                <div class="text-caption text-grey-darken-1 mb-2">
                  Or provide image URL:
                </div>
                <v-text-field
                  v-model="productData.image_url"
                  variant="outlined"
                  density="compact"
                  placeholder="https://example.com/image.jpg"
                  :disabled="loading || uploading || !!imageFile"
                  hide-details
                />
              </v-card>
            </v-col>

            <!-- Description -->
            <v-col cols="12">
              <v-textarea
                v-model="productData.description"
                label="Product Description *"
                variant="outlined"
                prepend-inner-icon="mdi-text"
                :rules="descriptionRules"
                :disabled="loading"
                rows="4"
                required
              />
            </v-col>

            <!-- Submit Buttons -->
            <v-col cols="12">
              <v-divider class="mb-4" />
              <div class="d-flex gap-3">
                <v-btn
                  type="submit"
                  color="orange"
                  variant="flat"
                  size="large"
                  :loading="loading"
                  prepend-icon="mdi-plus-circle"
                >
                  Add Product
                </v-btn>
                <v-btn
                  variant="outlined"
                  size="large"
                  :disabled="loading"
                  @click="resetForm"
                >
                  Clear Form
                </v-btn>
              </div>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
    </v-card>

    <!-- Products List -->
    <v-card v-show="showProductsList" elevation="2">
      <v-card-title
        class="bg-orange-lighten-5 py-4 d-flex justify-space-between align-center"
      >
        <div>
          <v-icon class="mr-2">mdi-format-list-bulleted</v-icon>
          All Products
        </div>
        <v-btn
          color="orange"
          variant="text"
          prepend-icon="mdi-refresh"
          @click="loadProducts"
          :loading="loadingProducts"
        >
          Refresh
        </v-btn>
      </v-card-title>

      <v-card-text class="pa-0">
        <!-- Loading State -->
        <div v-if="loadingProducts" class="text-center py-12">
          <v-progress-circular indeterminate color="orange" size="64" />
          <p class="text-body-1 mt-4">Loading products...</p>
        </div>

        <!-- No Products -->
        <div v-else-if="!products.length" class="text-center py-12">
          <v-icon size="80" color="grey-lighten-1">mdi-package-variant</v-icon>
          <h3 class="text-h6 mt-4 mb-2">No Products Yet</h3>
          <p class="text-body-2 text-grey-darken-1 mb-4">
            Start by adding your first product
          </p>
          <v-btn
            color="orange"
            variant="flat"
            @click="showProductsList = false"
          >
            Add Product
          </v-btn>
        </div>

        <!-- Products Table -->
        <v-data-table
          v-else
          :headers="headers"
          :items="products"
          :items-per-page="10"
          class="elevation-0"
        >
          <!-- Image Column -->
          <template #item.image_url="{ item }">
            <v-avatar size="50" rounded="lg" class="my-2">
              <v-img
                :src="item.image_url || '/images/placeholder.jpg'"
                :alt="item.name"
                cover
              />
            </v-avatar>
          </template>

          <!-- Price Column -->
          <template #item.price="{ item }">
            <div>
              <div class="font-weight-bold">
                KSh {{ formatPrice(item.price) }}
              </div>
              <div
                v-if="item.discount_percentage > 0"
                class="text-caption text-orange"
              >
                -{{ item.discount_percentage }}% off
              </div>
            </div>
          </template>

          <!-- Stock Column -->
          <template #item.stock="{ item }">
            <v-chip
              :color="getStockColor(item.stock)"
              size="small"
              variant="flat"
            >
              {{ item.stock }} units
            </v-chip>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="d-flex gap-2">
              <v-btn
                icon="mdi-pencil"
                size="small"
                variant="text"
                color="primary"
                @click="editProduct(item)"
              />
              <v-btn
                icon="mdi-delete"
                size="small"
                variant="text"
                color="error"
                @click="confirmDelete(item)"
              />
            </div>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="deleteDialog" max-width="500">
      <v-card>
        <v-card-title class="text-h6 bg-error-lighten-5">
          Confirm Delete
        </v-card-title>
        <v-card-text class="pt-4">
          Are you sure you want to delete
          <strong>{{ productToDelete?.name }}</strong
          >? This action cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="deleteDialog = false"> Cancel </v-btn>
          <v-btn
            color="error"
            variant="flat"
            :loading="deleting"
            @click="deleteProduct"
          >
            Delete
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Edit Product Dialog -->
    <v-dialog v-model="editDialog" max-width="800" scrollable>
      <v-card>
        <v-card-title class="bg-orange-lighten-5"> Edit Product </v-card-title>
        <v-card-text class="pt-6">
          <v-form ref="editForm">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="editProductData.name"
                  label="Product Name *"
                  variant="outlined"
                  :rules="nameRules"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="editProductData.category"
                  label="Category *"
                  variant="outlined"
                  :items="categories"
                  :rules="categoryRules"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="editProductData.brand"
                  label="Brand"
                  variant="outlined"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="editProductData.price"
                  label="Price (KSh) *"
                  variant="outlined"
                  type="number"
                  :rules="priceRules"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="editProductData.stock"
                  label="Stock Quantity *"
                  variant="outlined"
                  type="number"
                  :rules="stockRules"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model.number="editProductData.discount_percentage"
                  label="Discount %"
                  variant="outlined"
                  type="number"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="editProductData.image_url"
                  label="Image URL"
                  variant="outlined"
                  :disabled="updating"
                />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  v-model="editProductData.description"
                  label="Description *"
                  variant="outlined"
                  :rules="descriptionRules"
                  :disabled="updating"
                  rows="3"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            variant="text"
            @click="editDialog = false"
            :disabled="updating"
          >
            Cancel
          </v-btn>
          <v-btn
            color="orange"
            variant="flat"
            :loading="updating"
            @click="updateProduct"
          >
            Update Product
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import { uploadProductImage, deleteProductImage } from "@/utils/storage";

const API_BASE = "http://localhost:8080/api/v1";

// State
const form = ref(null);
const editForm = ref(null);
const loading = ref(false);
const loadingProducts = ref(false);
const updating = ref(false);
const deleting = ref(false);
const error = ref("");
const successMessage = ref("");
const showProductsList = ref(false);
const deleteDialog = ref(false);
const editDialog = ref(false);
const productToDelete = ref(null);
const products = ref([]);

// Image Upload State
const imageFile = ref(null);
const imagePreview = ref("");
const uploading = ref(false);
const uploadProgress = ref(0);

// Form Data
const productData = reactive({
  name: "",
  description: "",
  price: 0,
  category: "",
  stock: 0,
  image_url: "",
  brand: "",
  discount_percentage: 0,
});

const editProductData = reactive({
  id: "",
  name: "",
  description: "",
  price: 0,
  category: "",
  stock: 0,
  image_url: "",
  brand: "",
  discount_percentage: 0,
});

// Categories
const categories = [
  "Electronics",
  "Fashion",
  "Home & Garden",
  "Sports",
  "Books",
  "Toys",
  "Beauty",
  "Health",
  "Automotive",
  "Food & Grocery",
];

// Table Headers
const headers = [
  { title: "Image", key: "image_url", sortable: false },
  { title: "Name", key: "name" },
  { title: "Category", key: "category" },
  { title: "Brand", key: "brand" },
  { title: "Price", key: "price" },
  { title: "Stock", key: "stock" },
  { title: "Actions", key: "actions", sortable: false },
];

// Validation Rules
const nameRules = [
  (v) => !!v || "Product name is required",
  (v) => v.length >= 3 || "Name must be at least 3 characters",
];

const descriptionRules = [
  (v) => !!v || "Description is required",
  (v) => v.length >= 10 || "Description must be at least 10 characters",
];

const categoryRules = [(v) => !!v || "Category is required"];

const priceRules = [
  (v) => (v !== null && v !== "") || "Price is required",
  (v) => v > 0 || "Price must be greater than 0",
];

const stockRules = [
  (v) => (v !== null && v !== "") || "Stock is required",
  (v) => v >= 0 || "Stock cannot be negative",
];

// Load products on mount
onMounted(() => {
  loadProducts();
});

// Load Products
async function loadProducts() {
  loadingProducts.value = true;
  error.value = "";

  try {
    const response = await fetch(`${API_BASE}/products?page=1&page_size=100`);

    if (!response.ok) {
      throw new Error("Failed to load products");
    }

    const data = await response.json();
    products.value = data.products || [];
  } catch (err) {
    error.value = err.message;
    console.error("Error loading products:", err);
  } finally {
    loadingProducts.value = false;
  }
}

// Add Product
async function handleAddProduct() {
  const { valid } = await form.value.validate();

  if (!valid) {
    return;
  }

  loading.value = true;
  error.value = "";
  successMessage.value = "";

  try {
    // Use uploaded image URL if available, otherwise use manual URL
    const imageUrl = imagePreview.value || productData.image_url;

    if (!imageUrl) {
      throw new Error("Please upload an image or provide an image URL");
    }

    const response = await fetch(`${API_BASE}/products`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        ...productData,
        image_url: imageUrl,
      }),
    });

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || data.message || "Failed to add product");
    }

    const data = await response.json();
    successMessage.value = `Product "${productData.name}" added successfully!`;

    // Reset form
    resetForm();

    // Reload products if list is visible
    if (showProductsList.value) {
      await loadProducts();
    }

    // Auto-hide success message after 5 seconds
    setTimeout(() => {
      successMessage.value = "";
    }, 5000);
  } catch (err) {
    error.value = err.message;
    console.error("Error adding product:", err);
  } finally {
    loading.value = false;
  }
}

// Reset Form
function resetForm() {
  form.value?.reset();
  productData.name = "";
  productData.description = "";
  productData.price = 0;
  productData.category = "";
  productData.stock = 0;
  productData.image_url = "";
  productData.brand = "";
  productData.discount_percentage = 0;

  // Reset image upload state
  imageFile.value = null;
  imagePreview.value = "";
  uploadProgress.value = 0;
}

// Image Upload Handlers
function handleImageSelect(file) {
  if (!file) {
    removeImage();
    return;
  }

  // Validate file type
  const validTypes = ["image/jpeg", "image/png", "image/webp"];
  if (!validTypes.includes(file.type)) {
    error.value = "Please select a valid image file (JPEG, PNG, or WebP)";
    return;
  }

  // Validate file size (5MB max)
  const maxSize = 5 * 1024 * 1024; // 5MB in bytes
  if (file.size > maxSize) {
    error.value = "Image size must be less than 5MB";
    return;
  }

  imageFile.value = file;
  error.value = "";

  // Create preview
  const reader = new FileReader();
  reader.onload = (e) => {
    imagePreview.value = e.target.result;
  };
  reader.readAsDataURL(file);
}

async function uploadImage() {
  if (!imageFile.value) {
    error.value = "Please select an image to upload";
    return;
  }

  uploading.value = true;
  uploadProgress.value = 0;
  error.value = "";

  try {
    const downloadURL = await uploadProductImage(
      imageFile.value,
      null, // productId (optional, will generate unique name)
      (progress) => {
        uploadProgress.value = progress;
      }
    );

    imagePreview.value = downloadURL;
    productData.image_url = downloadURL;
    successMessage.value = "Image uploaded successfully!";

    // Clear file input after successful upload
    imageFile.value = null;

    setTimeout(() => {
      successMessage.value = "";
    }, 3000);
  } catch (err) {
    error.value = err.message || "Failed to upload image";
    console.error("Error uploading image:", err);
  } finally {
    uploading.value = false;
    uploadProgress.value = 0;
  }
}

function removeImage() {
  imageFile.value = null;
  imagePreview.value = "";
  uploadProgress.value = 0;
  productData.image_url = "";
}

// Edit Product
function editProduct(product) {
  editProductData.id = product.id;
  editProductData.name = product.name;
  editProductData.description = product.description;
  editProductData.price = product.price;
  editProductData.category = product.category;
  editProductData.stock = product.stock;
  editProductData.image_url = product.image_url || "";
  editProductData.brand = product.brand || "";
  editProductData.discount_percentage = product.discount_percentage || 0;
  editDialog.value = true;
}

// Update Product
async function updateProduct() {
  const { valid } = await editForm.value.validate();

  if (!valid) {
    return;
  }

  updating.value = true;

  try {
    const response = await fetch(`${API_BASE}/products/${editProductData.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: editProductData.name,
        description: editProductData.description,
        price: editProductData.price,
        category: editProductData.category,
        stock: editProductData.stock,
        image_url: editProductData.image_url,
        brand: editProductData.brand,
        discount_percentage: editProductData.discount_percentage,
      }),
    });

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || data.message || "Failed to update product");
    }

    successMessage.value = `Product "${editProductData.name}" updated successfully!`;
    editDialog.value = false;
    await loadProducts();

    setTimeout(() => {
      successMessage.value = "";
    }, 5000);
  } catch (err) {
    error.value = err.message;
    console.error("Error updating product:", err);
  } finally {
    updating.value = false;
  }
}

// Confirm Delete
function confirmDelete(product) {
  productToDelete.value = product;
  deleteDialog.value = true;
}

// Delete Product
async function deleteProduct() {
  deleting.value = true;

  try {
    const response = await fetch(
      `${API_BASE}/products/${productToDelete.value.id}`,
      {
        method: "DELETE",
      }
    );

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || data.message || "Failed to delete product");
    }

    successMessage.value = `Product "${productToDelete.value.name}" deleted successfully!`;
    deleteDialog.value = false;
    productToDelete.value = null;
    await loadProducts();

    setTimeout(() => {
      successMessage.value = "";
    }, 5000);
  } catch (err) {
    error.value = err.message;
    console.error("Error deleting product:", err);
  } finally {
    deleting.value = false;
  }
}

// Helper Functions
function formatPrice(price) {
  return new Intl.NumberFormat("en-KE", {
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  }).format(price);
}

function getStockColor(stock) {
  if (stock === 0) return "error";
  if (stock < 10) return "warning";
  return "success";
}
</script>

<style scoped>
.gap-3 {
  gap: 12px;
}
</style>
