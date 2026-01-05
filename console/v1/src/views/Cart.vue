<template>
  <v-container fluid class="pa-6 bg-grey-lighten-4" style="max-width: 1550px">
    <h1 class="text-h5 font-weight-regular mb-6">
      Cart ({{ cartStore.cartCount }})
    </h1>

    <!-- Empty Cart State -->
    <v-row v-if="cartStore.cartItems.length === 0">
      <v-col cols="12">
        <v-card elevation="0" class="text-center pa-8">
          <v-icon size="80" color="grey-lighten-1">mdi-cart-outline</v-icon>
          <h2 class="text-h5 mt-4 mb-2">Your cart is empty</h2>
          <p class="text-grey mb-4">Add items to get started</p>
          <v-btn color="orange" variant="flat" to="/">
            Continue Shopping
          </v-btn>
        </v-card>
      </v-col>
    </v-row>

    <v-row v-else>
      <!-- Left: Cart Items -->
      <v-col cols="12" md="8">
        <v-card
          v-for="item in cartStore.cartItems"
          :key="item.id"
          elevation="0"
          class="mb-3"
        >
          <v-card-text class="pa-4">
            <v-row align="center">
              <!-- Product Image -->
              <v-col cols="2">
                <v-img
                  :src="item.image_url || '/images/products/placeholder.jpg'"
                  :alt="item.name"
                  aspect-ratio="1"
                  cover
                  class="rounded cursor-pointer"
                  @click="goToProduct(item.id)"
                ></v-img>
              </v-col>

              <!-- Product Details -->
              <v-col cols="10">
                <v-row align="center">
                  <!-- Product Info -->
                  <v-col cols="12" md="6">
                    <div>
                      <h3
                        class="text-body-1 font-weight-regular mb-2 cursor-pointer"
                        @click="goToProduct(item.id)"
                      >
                        {{ item.name }}
                      </h3>
                      <div v-if="item.brand" class="text-body-2 text-grey mb-2">
                        Brand: {{ item.brand }}
                      </div>
                      <div
                        v-if="!item.in_stock"
                        class="text-body-2 text-red mb-2"
                      >
                        <v-icon size="small" color="red"
                          >mdi-alert-circle-outline</v-icon
                        >
                        Out of Stock
                      </div>
                      <div
                        v-else-if="item.stock < 5"
                        class="text-body-2 text-orange mb-2"
                      >
                        <v-icon size="small" color="orange"
                          >mdi-alert-circle-outline</v-icon
                        >
                        Only {{ item.stock }} left
                      </div>
                      <v-btn
                        variant="text"
                        color="orange"
                        size="small"
                        class="pa-0"
                        @click="removeItem(item.id)"
                      >
                        <v-icon size="small" class="mr-1"
                          >mdi-delete-outline</v-icon
                        >
                        Remove
                      </v-btn>
                    </div>
                  </v-col>

                  <!-- Price and Quantity -->
                  <v-col cols="12" md="6">
                    <v-row align="center" justify="end">
                      <!-- Price -->
                      <v-col cols="12" class="text-right">
                        <div class="text-h6 font-weight-medium">
                          KSh
                          {{
                            (item.final_price * item.quantity).toLocaleString()
                          }}
                        </div>
                        <div
                          v-if="item.discount_percentage > 0"
                          class="text-body-2"
                        >
                          <span
                            class="text-grey text-decoration-line-through mr-2"
                          >
                            KSh
                            {{ (item.price * item.quantity).toLocaleString() }}
                          </span>
                          <span class="text-orange"
                            >-{{ Math.round(item.discount_percentage) }}%</span
                          >
                        </div>
                      </v-col>

                      <!-- Quantity Controls -->
                      <v-col cols="12" class="d-flex justify-end align-center">
                        <v-btn
                          icon
                          size="small"
                          variant="flat"
                          color="grey-lighten-2"
                          @click="decrementQuantity(item.id)"
                        >
                          <v-icon>mdi-minus</v-icon>
                        </v-btn>
                        <span class="mx-4 text-body-1">{{
                          item.quantity
                        }}</span>
                        <v-btn
                          icon
                          size="small"
                          variant="flat"
                          color="orange"
                          :disabled="
                            !item.in_stock || item.quantity >= item.stock
                          "
                          @click="incrementQuantity(item.id)"
                        >
                          <v-icon color="white">mdi-plus</v-icon>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-col>
                </v-row>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Right: Cart Summary -->
      <v-col cols="12" md="4">
        <v-card elevation="0" class="sticky-summary">
          <v-card-title class="text-subtitle-1 font-weight-medium py-3">
            CART SUMMARY
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text class="pa-4">
            <v-row align="center" class="mb-4">
              <v-col cols="6">
                <span class="text-body-1">Subtotal</span>
              </v-col>
              <v-col cols="6" class="text-right">
                <span class="text-h6 font-weight-medium">
                  KSh {{ cartStore.cartTotal.toLocaleString() }}
                </span>
              </v-col>
            </v-row>

            <v-btn
              color="orange"
              variant="flat"
              block
              size="large"
              class="text-white font-weight-medium"
              @click="checkout"
            >
              Checkout (KSh {{ cartStore.cartTotal.toLocaleString() }})
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

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
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useCartStore } from "@/stores/cart";

const router = useRouter();
const cartStore = useCartStore();

const snackbar = ref({
  show: false,
  message: "",
  color: "success",
});

const incrementQuantity = (productId) => {
  cartStore.incrementQuantity(productId);
};

const decrementQuantity = (productId) => {
  cartStore.decrementQuantity(productId);
};

const removeItem = (productId) => {
  cartStore.removeFromCart(productId);
  showSnackbar("Item removed from cart", "info");
};

const goToProduct = (productId) => {
  router.push(`/product/${productId}`);
};

const checkout = () => {
  router.push({ name: "checkout" });
};

const showSnackbar = (message, color = "success") => {
  snackbar.value = {
    show: true,
    message,
    color,
  };
};
</script>

<style scoped>
.sticky-summary {
  position: sticky;
  top: 20px;
}

.cursor-pointer {
  cursor: pointer;
}
</style>
