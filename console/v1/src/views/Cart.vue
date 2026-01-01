<template>
  <v-container fluid class="pa-6 bg-grey-lighten-4" style="max-width: 1550px">
    <h1 class="text-h5 font-weight-regular mb-6">
      Cart ({{ cartItems.length }})
    </h1>

    <v-row>
      <!-- Left: Cart Items -->
      <v-col cols="12" md="8">
        <v-card
          v-for="(item, index) in cartItems"
          :key="index"
          elevation="0"
          class="mb-3"
        >
          <v-card-text class="pa-4">
            <v-row align="center">
              <!-- Product Image -->
              <v-col cols="2">
                <v-img
                  :src="item.image"
                  :alt="item.name"
                  aspect-ratio="1"
                  cover
                  class="rounded"
                ></v-img>
              </v-col>

              <!-- Product Details -->
              <v-col cols="10">
                <v-row align="center">
                  <!-- Product Info -->
                  <v-col cols="12" md="6">
                    <div>
                      <h3 class="text-body-1 font-weight-regular mb-2">
                        {{ item.name }}
                      </h3>
                      <div v-if="item.express" class="mb-2">
                        <v-img
                          src="/Jumia-Logo.jpg"
                          alt="Jumia Express"
                          max-width="100"
                          contain
                        ></v-img>
                      </div>
                      <div
                        v-if="item.variation"
                        class="text-body-2 text-grey mb-2"
                      >
                        Variation: {{ item.variation }}
                      </div>
                      <div v-if="item.stock" class="text-body-2 text-red mb-2">
                        <v-icon size="small" color="red"
                          >mdi-alert-circle-outline</v-icon
                        >
                        {{ item.stock }}
                      </div>
                      <v-btn
                        variant="text"
                        color="orange"
                        size="small"
                        class="pa-0"
                        @click="removeItem(index)"
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
                          KSh {{ item.price.toLocaleString() }}
                        </div>
                        <div v-if="item.originalPrice" class="text-body-2">
                          <span
                            class="text-grey text-decoration-line-through mr-2"
                          >
                            KSh {{ item.originalPrice.toLocaleString() }}
                          </span>
                          <span class="text-orange">{{ item.discount }}</span>
                        </div>
                      </v-col>

                      <!-- Quantity Controls -->
                      <v-col cols="12" class="d-flex justify-end align-center">
                        <v-btn
                          icon
                          size="small"
                          variant="flat"
                          color="grey-lighten-2"
                          @click="decrementQuantity(index)"
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
                          @click="incrementQuantity(index)"
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
                  KSh {{ subtotal.toLocaleString() }}
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
              Checkout (KSh {{ subtotal.toLocaleString() }})
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed } from "vue";

const cartItems = ref([
  {
    name: "AILYONS ELPF901K 5.1CH Home Theater Multimedia Speaker Sound System With Bluetooth () (1YR WRTY)",
    image: "/images/1.jpg",
    price: 11649,
    originalPrice: 18999,
    discount: "-39%",
    quantity: 1,
    stock: "1 units left",
    express: true,
  },
  {
    name: 'North Bayou NB-F160 Dual Monitor Desk Mount 17"–27" Full Motion Gas Spring Arm Stand Heavy Duty Adjustable Bracket',
    image: "/images/2.jpg",
    price: 11500,
    originalPrice: 14000,
    discount: "-18%",
    quantity: 1,
    stock: "5 units left",
    variation: '17"–27"',
    express: false,
  },
  {
    name: "12kg Dumbbell Set",
    image: "/images/3.jpg",
    price: 3199,
    originalPrice: 4900,
    discount: "-35%",
    quantity: 1,
    stock: "Few units left",
    express: false,
  },
  {
    name: "Wireless Bluetooth Headphones",
    image: "/images/4.jpg",
    price: 2500,
    originalPrice: 3500,
    discount: "-29%",
    quantity: 1,
    stock: null,
    express: false,
  },
]);

const subtotal = computed(() => {
  return cartItems.value.reduce((total, item) => {
    return total + item.price * item.quantity;
  }, 0);
});

const incrementQuantity = (index) => {
  cartItems.value[index].quantity++;
};

const decrementQuantity = (index) => {
  if (cartItems.value[index].quantity > 1) {
    cartItems.value[index].quantity--;
  }
};

const removeItem = (index) => {
  cartItems.value.splice(index, 1);
};

const checkout = () => {
  console.log("Proceeding to checkout with total:", subtotal.value);
};
</script>

<style scoped>
.sticky-summary {
  position: sticky;
  top: 20px;
}
</style>
