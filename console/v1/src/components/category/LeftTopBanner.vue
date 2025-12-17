<template>
  <div class="menu-wrapper">
    <!-- LEFT: Categories -->
    <v-card
      class="pa-2 rounded categories"
      elevation="2"
    >
      <v-btn
        flat
        block
        class="hover-primary justify-start"
        @mouseenter="active = 'official'"
      >
        <v-icon class="pr-2">mdi-basket-outline</v-icon>
        Official Stores
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'phones'">
        <v-icon class="pr-2">mdi-cellphone</v-icon>
        Phones & Tablets
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'tv'">
        <v-icon class="pr-2">mdi-television-play</v-icon>
        TVs & Audio
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'appliances'">
        <v-icon class="pr-2">mdi-coffee-maker-outline</v-icon>
        Appliances
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'health'">
        <v-icon class="pr-2">mdi-spa</v-icon>
        Health & Beauty
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'home'">
        <v-icon class="pr-2">mdi-home-outline</v-icon>
        Home & Office
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'fashion'">
        <v-icon class="pr-2">mdi-hanger</v-icon>
        Fashion
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'computing'">
        <v-icon class="pr-2">mdi-monitor</v-icon>
        Computing
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'gaming'">
        <v-icon class="pr-2">mdi-controller-classic-outline</v-icon>
        Gaming
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'supermarket'">
        <v-icon class="pr-2">mdi-cart-outline</v-icon>
        Supermarket
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'baby'">
        <v-icon class="pr-2">mdi-baby-face-outline</v-icon>
        Baby Products
      </v-btn>

      <v-btn flat block class="hover-primary justify-start" @mouseenter="active = 'other'">
        <v-icon class="pr-2">mdi-dots-horizontal</v-icon>
        Other categories
      </v-btn>
    </v-card>

    <!-- RIGHT: Hover Panel -->
    <v-card
      v-if="active"
      class="pa-4 rounded hover-panel"
      elevation="2"
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
</template>


<script setup>
import { ref } from 'vue'

const active = ref(null)

const submenu = {
  phones: [
    { title: 'Brands', items: ['Samsung', 'Tecno', 'Infinix', 'Oppo'] },
    { title: 'Accessories', items: ['Chargers', 'Earphones', 'Power Banks'] }
  ],
  tv: [
    { title: 'TV Brands', items: ['Vision Plus', 'Vitron', 'TCL'] },
    { title: 'Audio', items: ['Soundbars', 'Speakers'] }
  ],
  appliances: [
    { title: 'Kitchen', items: ['Blenders', 'Microwaves', 'Cookers'] },
    { title: 'Home', items: ['Fridges', 'Washing Machines'] }
  ],
  fashion: [
    { title: 'Men', items: ['T-Shirts', 'Shoes', 'Watches'] },
    { title: 'Women', items: ['Dresses', 'Handbags', 'Heels'] }
  ],
  supermarket: [
    { title: 'Groceries', items: ['Rice', 'Cooking Oil', 'Sugar'] },
    { title: 'Household', items: ['Detergents', 'Tissues'] }
  ]
}
</script>


<style scoped>
.menu-wrapper {
  display: flex;
  position: relative;
}

.categories {
  width: 250px;
  height: 450px;
}

.hover-panel {
  width: 750px;
  min-height: 450px;
  margin-left: 8px;
}

.hover-primary:hover {
  color: rgb(var(--v-theme-primary));
}

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
</style>
