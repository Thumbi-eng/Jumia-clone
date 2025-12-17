<template>
  <v-card elevation="0" rounded="lg">
    <!-- Header -->
    <div
      class="d-flex align-center justify-space-between px-4"
      style="height: 48px; background: #ffa000"
    >
      <div class="font-weight-bold text-white">
        Krisi Na Jumia | Top Deals
      </div>

      <v-btn
        variant="text"
        color="white"
        size="small"
        class="text-none"
      >
        See All →
      </v-btn>
    </div>

    <!-- Products Grid -->
    <div class="d-flex flex-wrap pa-3">
      <v-card
        v-for="(item, i) in items"
        :key="i"
        elevation="0"
        class="pa-2 mr-2 mb-2"
        style="width: 180px"
      >
        <!-- Image + Discount -->
        <div class="position-relative">
          <v-img
            :src="item.image"
            aspect-ratio="1"
            cover
            class="rounded"
          />

          <v-chip
            v-if="item.discount"
            size="small"
            color="orange-lighten-5"
            class="position-absolute top-0 right-0 ma-1"
            text-color="orange-darken-2"
            label
          >
            -{{ item.discount }}%
          </v-chip>
        </div>

        <!-- Title -->
        <div
          class="text-body-2 mt-2"
          style="line-height: 1.2"
        >
          {{ item.title }}
        </div>

        <!-- Price -->
        <div class="font-weight-bold mt-1">
          KSh {{ item.price }}
        </div>

        <!-- Old Price -->
        <div
          v-if="item.oldPrice"
          class="text-caption text-grey text-decoration-line-through"
        >
          KSh {{ item.oldPrice }}
        </div>
      </v-card>
    </div>
  </v-card>
</template>
<script setup>
import { ref } from "vue";

const items = ref([
  {
    title: "Samsung Galaxy A16, 6.7\"",
    price: "15,940",
    oldPrice: "19,000",
    discount: 16,
    image: "/images/products/galaxy-a16.jpg",
  },
  {
    title: "Hisense 7.5Kg Twin Tub",
    price: "15,599",
    oldPrice: "26,199",
    discount: 40,
    image: "/images/products/hisense-twin.jpg",
  },
  {
    title: "Nunix Hair Red Blow Dryer",
    price: "650",
    oldPrice: "883",
    discount: 26,
    image: "/images/products/blowdryer.jpg",
  },
  {
    title: "Apple MacBook Pro 13\"",
    price: "22,500",
    oldPrice: "30,000",
    discount: 25,
    image: "/images/products/macbook.jpg",
  },
  {
    title: "Hisense 94 Litres Fridge",
    price: "14,999",
    oldPrice: "25,399",
    discount: 41,
    image: "/images/products/fridge.jpg",
  },
  {
    title: "Kuhl S25 Ultra 8GB",
    price: "9,380",
    oldPrice: "13,400",
    discount: 30,
    image: "/images/products/kuhl.jpg",
  },
  {
    title: "Solar Light Three Rotatable",
    price: "1,033",
    oldPrice: "2,000",
    discount: 48,
    image: "/images/products/solar.jpg",
  },
  {
    title: "Euroken 3+1 Electric Cooker",
    price: "14,999",
    oldPrice: "26,999",
    discount: 44,
    image: "/images/products/cooker.jpg",
  },
  {
    title: "Em Electromate Water Dispenser",
    price: "2,399",
    oldPrice: "5,199",
    discount: 54,
    image: "/images/products/dispenser.jpg",
  },
  {
    title: "Kabras Premium Sugar 2kg",
    price: "339",
    oldPrice: "375",
    discount: 10,
    image: "/images/products/sugar.jpg",
  },
  {
    title: "Sunlight Lavender Hand Wash",
    price: "999",
    oldPrice: "1,200",
    discount: 17,
    image: "/images/products/sunlight.jpg",
  },
  {
    title: "HP Chromebook 11 G8",
    price: "6,999",
    oldPrice: "15,799",
    discount: 56,
    image: "/images/products/chromebook.jpg",
  },
]);
</script>
