<template>
  <v-container
    fluid
    class="pa-6 mt-6 bg-grey-lighten-4"
    style="max-width: 1550px"
  >
    <v-row>
      <!-- Left: Checkout Details -->
      <v-col cols="12" md="8">
        <!-- 1. Customer Address -->
        <v-card elevation="0" class="mb-3">
          <v-card-text class="pa-4">
            <v-row align="center" class="mb-2">
              <v-col cols="auto">
                <v-icon color="green" size="small">mdi-check-circle</v-icon>
                <span class="text-body-1 font-weight-medium ml-2">
                  1. CUSTOMER ADDRESS
                </span>
              </v-col>
              <v-spacer></v-spacer>
              <v-col cols="auto">
                <v-btn
                  variant="text"
                  color="primary"
                  size="small"
                  class="text-capitalize"
                >
                  Change
                  <v-icon size="small" class="ml-1">mdi-chevron-right</v-icon>
                </v-btn>
              </v-col>
            </v-row>
            <div class="text-body-1 font-weight-medium">victor Waiguru</div>
            <div class="text-body-2 text-grey">
              King'ong'o | Nyeri - Nyeri Town | +254 748218541
            </div>
          </v-card-text>
        </v-card>

        <!-- 2. Delivery Details -->
        <v-card elevation="0" class="mb-3">
          <v-card-text class="pa-4">
            <v-row align="center" class="mb-3">
              <v-col cols="auto">
                <v-icon color="green" size="small">mdi-check-circle</v-icon>
                <span class="text-body-1 font-weight-medium ml-2">
                  2. DELIVERY DETAILS
                </span>
              </v-col>
              <v-spacer></v-spacer>
              <v-col cols="auto">
                <v-btn
                  variant="text"
                  color="primary"
                  size="small"
                  class="text-capitalize"
                >
                  Change
                  <v-icon size="small" class="ml-1">mdi-chevron-right</v-icon>
                </v-btn>
              </v-col>
            </v-row>

            <!-- Pick-up Station Info -->
            <div class="mb-3">
              <div class="text-body-1 font-weight-medium mb-1">
                Pick-up Station
              </div>
              <div class="text-body-2">
                Delivery between
                <span class="font-weight-medium">05 January</span> and
                <span class="font-weight-medium">08 January</span>
              </div>
            </div>

            <v-divider class="my-4"></v-divider>

            <!-- Place your order header -->
            <v-row align="center" class="mb-4">
              <v-col>
                <h2 class="text-h6 font-weight-regular">Place your order</h2>
              </v-col>
              <v-col cols="auto">
                <v-img src="/Jumia-Logo.jpg" max-width="80" contain></v-img>
              </v-col>
            </v-row>

            <!-- Shipment 1/3 -->
            <v-card variant="outlined" class="mb-3">
              <v-card-text class="pa-3">
                <v-row align="center" class="mb-2">
                  <v-col>
                    <span class="text-body-2 font-weight-medium"
                      >Shipment 1/3</span
                    >
                  </v-col>
                  <v-col cols="auto">
                    <span class="text-body-2 text-grey"
                      >Fulfilled by Jumia</span
                    >
                  </v-col>
                </v-row>

                <div class="bg-grey-lighten-4 pa-3 rounded mb-3">
                  <div class="text-body-2 font-weight-medium mb-1">
                    Pick-up Station
                  </div>
                  <div class="text-body-2 text-grey mb-2">
                    Delivery between 05 January and 06 January
                  </div>
                  <v-img
                    src="/Jumia-Logo.jpg"
                    max-width="100"
                    contain
                    class="mb-2"
                  ></v-img>
                </div>

                <v-row
                  v-for="item in shipment1"
                  :key="item.name"
                  align="center"
                  class="mb-2"
                >
                  <v-col cols="auto">
                    <v-img
                      :src="item.image"
                      width="60"
                      height="60"
                      cover
                      class="rounded"
                    ></v-img>
                  </v-col>
                  <v-col>
                    <div class="text-body-2">x{{ item.quantity }}</div>
                    <div class="text-body-2">{{ item.name }}</div>
                    <div class="text-body-1 font-weight-medium">
                      KSh {{ item.price.toLocaleString() }}
                    </div>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>

            <!-- Shipment 2/3 -->
            <v-card variant="outlined" class="mb-3">
              <v-card-text class="pa-3">
                <v-row align="center" class="mb-2">
                  <v-col>
                    <span class="text-body-2 font-weight-medium"
                      >Shipment 2/3</span
                    >
                  </v-col>
                  <v-col cols="auto">
                    <span class="text-body-2 text-grey">
                      Fulfilled by Elvis & Ken Computer Store
                    </span>
                  </v-col>
                </v-row>

                <div class="bg-grey-lighten-4 pa-3 rounded mb-3">
                  <div class="text-body-2 font-weight-medium mb-1">
                    Pick-up Station
                  </div>
                  <div class="text-body-2 text-grey">
                    Delivery between 07 January and 08 January
                  </div>
                </div>

                <v-row
                  v-for="item in shipment2"
                  :key="item.name"
                  align="center"
                  class="mb-2"
                >
                  <v-col cols="auto">
                    <v-img
                      :src="item.image"
                      width="60"
                      height="60"
                      cover
                      class="rounded"
                    ></v-img>
                  </v-col>
                  <v-col>
                    <div class="text-body-2">x{{ item.quantity }}</div>
                    <div class="text-body-2">{{ item.name }}</div>
                    <div class="text-body-1 font-weight-medium">
                      KSh {{ item.price.toLocaleString() }}
                    </div>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>

            <!-- Shipment 3/3 -->
            <v-card variant="outlined" class="mb-3">
              <v-card-text class="pa-3">
                <v-row align="center" class="mb-2">
                  <v-col>
                    <span class="text-body-2 font-weight-medium"
                      >Shipment 3/3</span
                    >
                  </v-col>
                  <v-col cols="auto">
                    <span class="text-body-2 text-grey">
                      Fulfilled by Mult station gym shop
                    </span>
                  </v-col>
                </v-row>

                <div class="bg-grey-lighten-4 pa-3 rounded mb-3">
                  <div class="text-body-2 font-weight-medium mb-1">
                    Pick-up Station
                  </div>
                  <div class="text-body-2 text-grey">
                    Delivery between 07 January and 08 January
                  </div>
                </div>

                <v-row
                  v-for="item in shipment3"
                  :key="item.name"
                  align="center"
                  class="mb-2"
                >
                  <v-col cols="auto">
                    <v-img
                      :src="item.image"
                      width="60"
                      height="60"
                      cover
                      class="rounded"
                    ></v-img>
                  </v-col>
                  <v-col>
                    <div class="text-body-2">x{{ item.quantity }}</div>
                    <div class="text-body-2">{{ item.name }}</div>
                    <div class="text-body-1 font-weight-medium">
                      KSh {{ item.price.toLocaleString() }}
                    </div>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>

            <div class="text-center my-3">
              <v-btn variant="text" color="orange" @click="modifyCart">
                Modify cart
              </v-btn>
            </div>
          </v-card-text>
        </v-card>

        <!-- 3. Payment Method -->
        <v-card elevation="0" class="mb-3">
          <v-card-text class="pa-4">
            <v-row align="center" class="mb-3">
              <v-col cols="auto">
                <v-icon
                  :color="
                    selectedPaymentMethod && !showPaymentSelection
                      ? 'green'
                      : 'grey'
                  "
                  size="small"
                >
                  mdi-check-circle
                </v-icon>
                <span class="text-body-1 font-weight-medium ml-2">
                  3. PAYMENT METHOD
                </span>
              </v-col>
              <v-spacer></v-spacer>
              <v-col cols="auto" v-if="!showPaymentSelection">
                <v-btn
                  variant="text"
                  color="primary"
                  size="small"
                  class="text-capitalize"
                  @click="openPaymentSelection"
                >
                  Change
                  <v-icon size="small" class="ml-1">mdi-chevron-right</v-icon>
                </v-btn>
              </v-col>
            </v-row>

            <!-- Selected Payment Method (when not in selection mode) -->
            <div v-if="!showPaymentSelection">
              <div v-if="selectedPaymentMethod">
                <div class="text-body-1 font-weight-medium mb-1">
                  {{ selectedPaymentMethod.title }}
                </div>
                <div class="text-body-2 text-grey">
                  {{ selectedPaymentMethod.description }}
                </div>
              </div>
              <div v-else class="text-body-2 text-grey">
                Please select a payment method
              </div>
            </div>

            <!-- Payment Selection (when in selection mode) -->
            <div v-else>
              <!-- Pre-pay Now Section -->
              <div class="text-body-1 font-weight-medium mb-3">Pre-pay Now</div>

              <v-radio-group v-model="tempPaymentMethod" hide-details>
                <!-- Pay Now With Bank Card -->
                <v-card
                  variant="outlined"
                  class="mb-3 pa-3"
                  :class="{
                    'border-orange': tempPaymentMethod === 'bank-card',
                  }"
                >
                  <v-radio value="bank-card" color="orange">
                    <template v-slot:label>
                      <div
                        class="d-flex align-center justify-space-between w-100"
                      >
                        <div>
                          <div class="text-body-1 font-weight-medium">
                            Pay Now With Bank Card
                          </div>
                          <div class="text-body-2 text-grey">
                            Pay now fast and securely with JumiaPay, Mastercard
                          </div>
                        </div>
                        <v-icon color="primary">mdi-credit-card</v-icon>
                      </div>
                    </template>
                  </v-radio>

                  <!-- Jumia Store Credit -->
                  <div
                    v-if="tempPaymentMethod === 'bank-card'"
                    class="mt-3 ml-8"
                  >
                    <div
                      class="d-flex align-center pa-2 bg-blue-lighten-5 rounded"
                    >
                      <v-icon color="blue" class="mr-2"
                        >mdi-credit-card-outline</v-icon
                      >
                      <span class="text-body-2 text-blue">
                        Jumia store credit balance: KSh 0
                      </span>
                    </div>
                  </div>
                </v-card>

                <!-- Lipa Pole Pole -->
                <v-card
                  variant="outlined"
                  class="mb-3 pa-3"
                  :class="{
                    'border-orange': tempPaymentMethod === 'lipa-pole-pole',
                  }"
                >
                  <v-radio value="lipa-pole-pole" color="orange">
                    <template v-slot:label>
                      <div
                        class="d-flex align-center justify-space-between w-100"
                      >
                        <div>
                          <div class="text-body-1 font-weight-medium">
                            Lipa Pole Pole
                          </div>
                          <div class="text-body-2 text-grey">
                            Buy Now Pay Later with a financier of your choice.
                            Enjoy up to 12 months repayment period.
                          </div>
                        </div>
                        <v-icon color="orange">mdi-wallet-outline</v-icon>
                      </div>
                    </template>
                  </v-radio>
                </v-card>

                <!-- Pay Now with Airtel -->
                <v-card
                  variant="outlined"
                  class="mb-3 pa-3"
                  :class="{ 'border-orange': tempPaymentMethod === 'airtel' }"
                >
                  <v-radio value="airtel" color="orange">
                    <template v-slot:label>
                      <div
                        class="d-flex align-center justify-space-between w-100"
                      >
                        <div>
                          <div class="text-body-1 font-weight-medium">
                            Pay Now with Airtel
                          </div>
                          <div class="text-body-2 text-grey">
                            Pay now fast and securely with your Airtel Money
                            account.
                          </div>
                        </div>
                        <v-icon color="red">mdi-cellphone</v-icon>
                      </div>
                    </template>
                  </v-radio>
                </v-card>

                <!-- Pay Now with M-Pesa -->
                <v-card
                  variant="outlined"
                  class="mb-3 pa-3"
                  :class="{ 'border-orange': tempPaymentMethod === 'mpesa' }"
                >
                  <v-radio value="mpesa" color="orange">
                    <template v-slot:label>
                      <div
                        class="d-flex align-center justify-space-between w-100"
                      >
                        <div>
                          <div class="text-body-1 font-weight-medium">
                            Pay Now with M-Pesa
                          </div>
                          <div class="text-body-2 text-grey">
                            Pay now fast and securely with your Mpesa account.
                          </div>
                        </div>
                        <v-icon color="green">mdi-cellphone</v-icon>
                      </div>
                    </template>
                  </v-radio>
                </v-card>
              </v-radio-group>

              <!-- Pay on Delivery Section -->
              <div class="text-body-1 font-weight-medium my-3">
                Pay on Delivery
              </div>

              <v-radio-group v-model="tempPaymentMethod" hide-details>
                <v-card
                  variant="outlined"
                  class="mb-3 pa-3"
                  :class="{
                    'border-orange': tempPaymentMethod === 'cash-on-delivery',
                  }"
                >
                  <v-radio
                    value="cash-on-delivery"
                    color="orange"
                    :disabled="paymentUnavailable"
                  >
                    <template v-slot:label>
                      <div
                        class="d-flex align-center justify-space-between w-100"
                      >
                        <div>
                          <div class="text-body-1 font-weight-medium">
                            Pay on delivery with Mobile Money and Bank Cards
                          </div>
                        </div>
                        <v-icon color="orange">mdi-cash</v-icon>
                      </div>
                    </template>
                  </v-radio>
                </v-card>

                <!-- Error Message -->
                <v-alert
                  v-if="
                    paymentUnavailable &&
                    tempPaymentMethod === 'cash-on-delivery'
                  "
                  type="info"
                  variant="tonal"
                  class="mb-3"
                  density="compact"
                >
                  <div class="text-body-2">
                    <strong>Oops! Unavailable Payment Method.</strong>
                  </div>
                  <div class="text-caption">
                    This may be because: 1) Your order is below the minimum
                    purchase amount of 399 Ksh, or above the maximum purchase
                    amount of 50,000 Ksh; 2) Cash on delivery is not available
                    for your delivery address, delivery method, or selected
                    item. 3) You have had multiple failed delivery attempts or
                    canceled orders 4) The phone number you are using to place
                    the order is a number that has a restriction.
                  </div>
                </v-alert>
              </v-radio-group>

              <!-- Confirm Button -->
              <v-btn
                color="orange"
                variant="flat"
                class="text-white mt-4"
                block
                @click="confirmPaymentMethod"
                :disabled="!tempPaymentMethod"
              >
                Confirm payment method
              </v-btn>
            </div>
          </v-card-text>
        </v-card>

        <!-- Back to shopping -->
        <v-btn variant="text" color="primary" class="mt-3" @click="goBack">
          <v-icon class="mr-2">mdi-chevron-left</v-icon>
          Go back & continue shopping
        </v-btn>
      </v-col>

      <!-- Right: Order Summary -->
      <v-col cols="12" md="4">
        <v-card elevation="0" class="sticky-summary">
          <v-card-title class="text-subtitle-1 font-weight-medium py-3">
            Order summary
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text class="pa-4">
            <!-- Items Total -->
            <v-row align="center" class="mb-2">
              <v-col>
                <span class="text-body-2">Item's total ({{ totalItems }})</span>
              </v-col>
              <v-col cols="auto">
                <span class="text-body-1"
                  >KSh {{ itemsTotal.toLocaleString() }}</span
                >
              </v-col>
            </v-row>

            <!-- Delivery Fees -->
            <v-row align="center" class="mb-3">
              <v-col>
                <span class="text-body-2">Delivery fees</span>
              </v-col>
              <v-col cols="auto">
                <span class="text-body-1"
                  >KSh {{ deliveryFees.toLocaleString() }}</span
                >
              </v-col>
            </v-row>

            <v-divider class="mb-3"></v-divider>

            <!-- Total -->
            <v-row align="center" class="mb-4">
              <v-col>
                <span class="text-body-1 font-weight-medium">Total</span>
              </v-col>
              <v-col cols="auto">
                <span class="text-h6 font-weight-bold">
                  KSh {{ total.toLocaleString() }}
                </span>
              </v-col>
            </v-row>

            <!-- Promo Code -->
            <v-row class="mb-4">
              <v-col cols="8">
                <v-text-field
                  v-model="promoCode"
                  placeholder="Enter code here"
                  density="compact"
                  variant="outlined"
                  hide-details
                  prepend-inner-icon="mdi-ticket-percent-outline"
                ></v-text-field>
              </v-col>
              <v-col cols="4">
                <v-btn
                  variant="outlined"
                  color="grey"
                  block
                  @click="applyPromo"
                  class="text-capitalize"
                >
                  APPLY
                </v-btn>
              </v-col>
            </v-row>

            <!-- Confirm Order Button -->
            <v-btn
              color="orange"
              variant="flat"
              block
              size="large"
              class="text-white font-weight-medium mb-3"
              @click="confirmOrder"
            >
              CONFIRM ORDER
            </v-btn>

            <!-- Terms & Conditions -->
            <div class="text-center mb-4">
              <a href="#" class="text-primary text-body-2 text-decoration-none">
                Terms & Conditions
              </a>
            </div>

            <v-divider class="mb-3"></v-divider>

            <!-- Help Icons -->
            <v-row class="text-center">
              <v-col cols="4">
                <v-icon size="24" color="grey">mdi-phone-outline</v-icon>
                <div class="text-caption">Need Help?</div>
                <div class="text-caption text-primary">Contact us</div>
              </v-col>
              <v-col cols="4">
                <v-icon size="24" color="grey">mdi-refresh</v-icon>
                <div class="text-caption">Easy</div>
                <div class="text-caption">Returns</div>
              </v-col>
              <v-col cols="4">
                <v-icon size="24" color="grey">mdi-shield-check-outline</v-icon>
                <div class="text-caption">Secure</div>
                <div class="text-caption">Payments</div>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const promoCode = ref("");
const showPaymentSelection = ref(false);
const tempPaymentMethod = ref("");
const selectedPaymentMethod = ref(null);
const paymentUnavailable = ref(true);

// Payment methods mapping
const paymentMethods = {
  "bank-card": {
    title: "Pay Now With Bank Card",
    description: "Pay now fast and securely with JumiaPay, Mastercard",
  },
  "lipa-pole-pole": {
    title: "Lipa Pole Pole",
    description:
      "Buy Now Pay Later with a financier of your choice. Enjoy up to 12 months repayment period.",
  },
  airtel: {
    title: "Pay Now with Airtel",
    description: "Pay now fast and securely with your Airtel Money account.",
  },
  mpesa: {
    title: "Pay Now with M-Pesa",
    description: "Pay now fast and securely with your Mpesa account.",
  },
  "cash-on-delivery": {
    title: "Pay on delivery with Mobile Money and Bank Cards",
    description: "",
  },
};

// Shipment data
const shipment1 = ref([
  {
    name: "ELPF901K 5.1CH Home Theater ...",
    image: "/images/1.jpg",
    price: 11649,
    quantity: 1,
  },
  {
    name: "50Q6500H, ...",
    image: "/images/5.jpg",
    price: 41499,
    quantity: 1,
  },
]);

const shipment2 = ref([
  {
    name: "NB-F160 Dual Monitor Desk Mou...",
    image: "/images/2.jpg",
    price: 11500,
    quantity: 1,
  },
]);

const shipment3 = ref([
  {
    name: "12kg Dumbbell Set",
    image: "/images/3.jpg",
    price: 3199,
    quantity: 1,
  },
]);

// Order summary calculations
const itemsTotal = ref(67847);
const deliveryFees = ref(1420);

const total = computed(() => itemsTotal.value + deliveryFees.value);
const totalItems = computed(() => {
  return (
    shipment1.value.length + shipment2.value.length + shipment3.value.length
  );
});

const confirmOrder = () => {
  console.log("Order confirmed!");
  // Navigate to order confirmation page
};

const applyPromo = () => {
  console.log("Applying promo code:", promoCode.value);
};

const modifyCart = () => {
  router.push({ name: "cart" });
};

const goBack = () => {
  router.push({ name: "home" });
};

const openPaymentSelection = () => {
  tempPaymentMethod.value = selectedPaymentMethod.value
    ? Object.keys(paymentMethods).find(
        (key) => paymentMethods[key].title === selectedPaymentMethod.value.title
      )
    : "";
  showPaymentSelection.value = true;
};

const confirmPaymentMethod = () => {
  if (tempPaymentMethod.value) {
    selectedPaymentMethod.value = paymentMethods[tempPaymentMethod.value];
    showPaymentSelection.value = false;
  }
};
</script>

<style scoped>
.sticky-summary {
  position: sticky;
  top: 20px;
}

.border-orange {
  border-color: #ff9800 !important;
  border-width: 2px !important;
}

.w-100 {
  width: 100%;
}
</style>
