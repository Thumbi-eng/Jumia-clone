<template>
  <v-row
    class="mt-5 align-stretch"
    style="max-width: 1350px; margin: 0 auto"
    dense
  >
    <!-- LEFT: Categories + Hover Panel -->
    <v-col cols="12" md="2" class="px-1">
      <div class="menu-wrapper">
        <!-- Categories -->
        <v-card class="pa-2 rounded categories" elevation="1">
          <v-btn
            v-for="cat in categories"
            :key="cat.key"
            flat
            block
            class="hover-primary justify-start"
            @mouseenter="active = cat.key"
          >
            <v-icon class="pr-2">{{ cat.icon }}</v-icon>
            {{ cat.label }}
          </v-btn>
        </v-card>

        <!-- Hover Panel -->
        <v-card
          v-if="active && submenu[active]"
          class="pa-4 rounded hover-panel"
          elevation="1"
          @mouseleave="active = null"
        >
          <v-row>
            <v-col
              v-for="section in submenu[active]"
              :key="section.title"
              cols="4"
            >
              <h4 class="mb-2">{{ section.title }}</h4>
              <div
                v-for="item in section.items"
                :key="item"
                class="submenu-item"
              >
                {{ item }}
              </div>
            </v-col>
          </v-row>
        </v-card>
      </div>
    </v-col>

    <!-- CENTER: Hero carousel -->
    <v-col cols="12" md="7" class="px-1">
      <v-card class="rounded overflow-hidden fill-height" elevation="0">
        <v-carousel
          hide-delimiter-background
          cycle
          height="320"
          v-model="current"
          show-arrows-on-hover
        >
          <v-carousel-item v-for="(slide, i) in slides" :key="i">
            <div
              class="d-flex align-center"
              :style="{ background: slide.bg, height: '320px' }"
            >
              <v-row class="fill-height px-6" align="center" no-gutters>
                <v-col cols="8">
                  <div class="text-white">
                    <div class="text-caption font-weight-bold">
                      {{ slide.kicker }}
                    </div>

                    <div
                      class="text-h5 font-weight-bold mt-2"
                      style="white-space: pre-line"
                    >
                      {{ slide.title }}
                    </div>

                    <v-chip
                      v-if="slide.price"
                      class="mt-3"
                      color="white"
                      text-color="#ff4b00"
                      size="small"
                    >
                      {{ slide.price }}
                    </v-chip>

                    <v-btn
                      class="mt-4"
                      variant="text"
                      color="white"
                      size="small"
                    >
                      SHOP NOW
                    </v-btn>
                  </div>
                </v-col>

                <v-col cols="4" class="text-center">
                  <v-img
                    v-if="slide.image"
                    :src="slide.image"
                    max-width="170"
                    contain
                  />
                </v-col>
              </v-row>
            </div>
          </v-carousel-item>
        </v-carousel>

        <!-- Dots -->
        <v-row class="justify-center py-2">
          <v-btn
            v-for="(_, i) in slides"
            :key="i"
            size="x-small"
            :variant="i === current ? 'flat' : 'text'"
            class="mx-1 dot-btn"
            color="orange"
            @click="current = i"
          />
        </v-row>
      </v-card>
    </v-col>

    <!-- RIGHT: Promo cards -->
    <v-col cols="12" md="3" class="px-1 d-flex flex-column">
      <v-card class="pa-2 mb-2" elevation="1">
        <v-list density="compact">
          <v-list-item v-for="(item, i) in rightItems" :key="i">
            <template #prepend>
              <v-icon size="22">{{ item.icon }}</v-icon>
            </template>
            <v-list-item-title class="text-body-2">
              {{ item.title }}
            </v-list-item-title>
            <v-list-item-subtitle class="text-caption">
              {{ item.subtitle }}
            </v-list-item-subtitle>
          </v-list-item>
        </v-list>
      </v-card>

      <v-card class="pa-4 mt-auto bg-black text-white" elevation="1">
        <div class="d-flex align-center">
          <v-icon color="success" size="42" class="mr-3">
            mdi-phone-in-talk
          </v-icon>
          <div class="font-weight-bold text-body-2">
            CALL OR WHATSAPP<br />
            <strong>0711 011 011</strong><br />
            TO ORDER
          </div>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from "vue";

const current = ref(0);
const active = ref(null);

const categories = [
  { key: "official", label: "Official Stores", icon: "mdi-storefront-outline" },
  { key: "phones", label: "Phones & Tablets", icon: "mdi-cellphone" },
  { key: "tv", label: "TVs & Audio", icon: "mdi-television-play" },
  { key: "appliances", label: "Appliances", icon: "mdi-coffee-maker-outline" },
  { key: "health", label: "Health & Beauty", icon: "mdi-spa" },
  { key: "home", label: "Home & Office", icon: "mdi-home-outline" },
  { key: "fashion", label: "Fashion", icon: "mdi-hanger" },
  { key: "computing", label: "Computing", icon: "mdi-monitor" },
  { key: "gaming", label: "Gaming", icon: "mdi-controller-classic-outline" },
  { key: "supermarket", label: "Supermarket", icon: "mdi-cart-outline" },
  { key: "baby", label: "Baby Products", icon: "mdi-baby-face-outline" },
  { key: "other", label: "Other categories", icon: "mdi-dots-horizontal" }
];

const submenu = {
  phones: [
    { title: "Brands", items: ["Samsung", "Tecno", "Infinix", "Oppo"] },
    { title: "Accessories", items: ["Chargers", "Earphones", "Power Banks"] }
  ],
  appliances: [
    { title: "Kitchen", items: ["Blenders", "Microwaves", "Cookers"] },
    { title: "Home", items: ["Fridges", "Washing Machines"] }
  ],
  fashion: [
    { title: "Men", items: ["T-Shirts", "Shoes", "Watches"] },
    { title: "Women", items: ["Dresses", "Handbags", "Heels"] }
  ],
  supermarket: [
    { title: "Groceries", items: ["Rice", "Cooking Oil", "Sugar"] },
    { title: "Household", items: ["Detergents", "Tissues"] }
  ]
};

const slides = [
  {
    kicker: "KRISI NA JUMIA",
    title: "HOMMY\n3 Column Metallic Wardrobe",
    price: "KSHS 1,749",
    bg: "linear-gradient(90deg,#ff9f00,#ff7a00)",
    image: "/images/slide1"
  },
  {
    kicker: "",
    title: "Discover Amazing Deals",
    bg: "linear-gradient(90deg,#ddd,#bbb)",
    image: "/images/slide2"
  },
  {
    kicker: "",
    title: "Seasonal Offers",
    bg: "linear-gradient(90deg,#e8e8e8,#d0d0d0)",
    image: "/images/slide3"
  }
];

const rightItems = [
  { icon: "mdi-whatsapp", title: "WhatsApp", subtitle: "Text To Order" },
  { icon: "mdi-storefront", title: "CHINA TOWN", subtitle: "NOW ON JUMIA" },
  { icon: "mdi-sale", title: "SELL ON JUMIA", subtitle: "Millions Of Visitors" }
];
</script>


<style scoped>
.menu-wrapper {
  display: flex;
  position: relative;
}

.categories {
  width: 100%;
  height: 450px;
}

.hover-panel {
  position: absolute;
  left: 100%;
  top: 0;
  width: 750px;
  min-height: 450px;
  margin-left: 8px;
  z-index: 10;
}

.hover-primary:hover,
.hover-primary:hover .v-icon {
  color: rgb(var(--v-theme-primary));
}

.submenu-item {
  font-size: 14px;
  margin-bottom: 6px;
  cursor: pointer;
}

.submenu-item:hover {
  color: rgb(var(--v-theme-primary));
}

.dot-btn {
  min-width: 8px;
  height: 8px;
  border-radius: 50%;
}
</style>
