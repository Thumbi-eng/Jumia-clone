<template>
  <v-row
    class="mt-5 align-stretch"
    style="max-width: 1550px; "

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
  v-model="current"
  cycle
  height="460"
  hide-delimiter-background
  show-arrows-on-hover
>
  <v-carousel-item
    v-for="(slide, i) in slides"
    :key="i"
  >
    <v-img
      :src="slide.image"
      height="460"
      cover
      class="rounded"
    />
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
    <v-col cols="12" md="3" class="px-2 d-flex flex-column">
      <v-card class="pa-2 mb-2"
      elevation="1"
      height="230"
      >
      <p class='d-flex '>

        <v-btn
        class="ma-2"
        color="green"
        icon="mdi-whatsapp"
      ></v-btn>
      <div>
      <p class='text-h6'>Whatsapp</p>
      <p>Text to order</p>
      </div>

      </p>

       <p class='d-flex pa-1'>

        <img  src='/images/hot-deals.png'
            height='50'
            class="m-0"
            />
      <div>
      <p class='text-h6'>CHINA TOWN</p>
      <p>NOW ON JUMIA</p>
      </div>

      </p>

       <p class='d-flex '>

        <img  src='/images/tl-cash.png'
            height='50'
            class="m-0"
            />
      <div>
      <p class='text-h6'>SELL ON JUMIA</p>
      <p>Millions Of Visitors</p>
      </div>

      </p>


      </v-card>

      <v-card
      class="pa-4
       bg-black
       text-white"
        elevation="1"
        height="220"
        >
        <v-row class='d-md-flex flex-column align-center justify-center' >
          <v-col>
            <v-img src='/images/tl-call.png'
            height='80'/>
             </v-col>

            <p class=' font-weight-bold text-h6   '>CALL OR WHATSAPP</p>

            <p class='font-weight-bold text-h4 text-primary'>0711 011 011</p>


            <p class='font-weight-bold text-h4'>TO ORDER</p>



        </v-row>

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
    
    image: "/images/slide1.png"
  },
  {
 
    image: "/images/slide2.png"
  },
  {
   
    image: "/images/slide3.png"
  },
   {
    
    image: "/images/slide2.png"
  }
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
