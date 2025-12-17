<template>
  <v-sheet style="background: #4b4b4f; color: rgba(255, 255, 255, 0.92)">
    <v-container fluid class="py-12">
      <v-row>
        <v-col cols="12" md="3">
          <div class="subtitle-2 font-weight-bold mb-4">NEED HELP?</div>
          <div class="caption" v-for="(l, i) in helpLinks" :key="i">
            {{ l }}
          </div>
        </v-col>

        <v-col cols="12" md="3">
          <div class="subtitle-2 font-weight-bold mb-4">USEFUL LINKS</div>
          <div class="caption" v-for="(l, i) in usefulLinks" :key="i">
            {{ l }}
          </div>
        </v-col>

        <v-col cols="12" md="2">
          <div class="subtitle-2 font-weight-bold mb-4">ABOUT JUMIA</div>
          <div class="caption" v-for="(l, i) in aboutLinks" :key="i">
            {{ l }}
          </div>
        </v-col>

        <v-col cols="12" md="2">
          <div class="subtitle-2 font-weight-bold mb-4">
            MAKE MONEY WITH JUMIA
          </div>
          <div class="caption" v-for="(l, i) in makeMoneyLinks" :key="i">
            {{ l }}
          </div>
        </v-col>

        <v-col cols="12" md="2">
          <div class="subtitle-2 font-weight-bold mb-4">
            JUMIA INTERNATIONAL
          </div>
          <div class="caption" v-for="(l, i) in intlLinks" :key="i">
            {{ l }}
          </div>
        </v-col>
      </v-row>

      <v-row class="mt-10">
        <v-col cols="12" md="6">
          <div class="subtitle-2 font-weight-bold mb-3">JOIN US ON</div>
          <div class="d-flex align-center">
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-facebook</v-icon></v-btn
            >
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-instagram</v-icon></v-btn
            >
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-twitter</v-icon></v-btn
            >
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-youtube</v-icon></v-btn
            >
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-linkedin</v-icon></v-btn
            >
            <v-btn icon class="mr-2" color="white"
              ><v-icon>mdi-tiktok</v-icon></v-btn
            >
          </div>
        </v-col>

        <v-col cols="12" md="6">
          <div class="subtitle-2 font-weight-bold mb-3">PAYMENT METHODS</div>
          <div class="d-flex align-center">
            <v-icon class="mr-4">mdi-credit-card-outline</v-icon>
            <v-icon class="mr-4">mdi-mastercard</v-icon>
            <v-icon class="mr-4">mdi-bank-transfer</v-icon>
            <v-icon class="mr-4">mdi-wallet</v-icon>
            <div class="ml-4 caption">JUMIA PAY</div>
          </div>
        </v-col>
      </v-row>

      <v-row class="mt-10">
        <v-col cols="12">
          <v-row>
            <v-col
              v-for="(group, idx) in brandGroups"
              :key="idx"
              cols="6"
              md="2"
            >
              <div v-for="(b, i) in group" :key="i" class="caption">
                {{ b }}
              </div>
            </v-col>
          </v-row>
        </v-col>
      </v-row>

      <v-divider class="my-8" color="rgba(255,255,255,0.08)" />

      <v-row align="center" class="py-4">
        <v-col cols="12" class="text-center">
          <v-img
            contain
            max-width="220"
            src=""
            alt="pay-delivery"
            class="mx-auto"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-sheet>
</template>

<script setup>
const helpLinks = ["Chat with us", "Help Center", "Contact Us"];

const usefulLinks = [
  "Track Your Order",
  "Shipping and delivery",
  "Pick-up Stations",
  "Return Policy",
  "How to Order?",
  "Dispute Resolution Policy",
  "Corporate and Bulk Purchase",
  "Advertise with Jumia",
  "Report a Product",
  "Jumia Payment Information Guidelines",
  "Black Friday",
];

const aboutLinks = [
  "About us",
  "Returns and Refunds Policy",
  "Jumia Careers",
  "Jumia Express",
  "Terms and Conditions",
  "Store Credit Terms and Conditions",
  "Privacy Notice",
  "Cookies Notice",
  "Flash Sales",
];

const makeMoneyLinks = [
  "Sell on Jumia",
  "Vendor Hub",
  "Become a Sales Consultant",
  "Jumia City Partner Program",
];

const intlLinks = [
  "Algeria",
  "Ivory Coast",
  "Egypt",
  "Ghana",
  "Morocco",
  "Nigeria",
  "Senegal",
  "Uganda",
];

// brand names grouped into columns to match the image layout
const brandGroups = [
  ["Adidas", "AILYONS", "Apple", "Asus", "Bruhm"],
  ["Canon", "Cantu", "DELL", "Dove", "Dr. Rashel"],
  ["Epson", "Garnier", "Hisense", "HP", "Infinix"],
  ["Itel", "Jameson", "L'Oréal Paris", "Lenovo", "Logitech"],
  ["Maybelline", "Mika", "NIVEA", "Nunix", "Oppo"],
  [
    "Ramtons",
    "Roch",
    "Samsung",
    "Sandisk",
    "Skyworth",
    "Sony",
    "TCL",
    "Tecno",
    "Transcend",
    "Tusker",
    "USN",
    "Vision Plus",
    "Vitron",
    "VON",
    "XIAOMI",
  ],
];
</script>

<!-- All styling uses Vuetify utility classes and inline styles -->
