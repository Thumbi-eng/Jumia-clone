<template>
  <v-card class="flash-wrapper rounded-lg"
   elevation="0"
   style="max-width: 1550px; min-width: 1540px; "
>
    <!-- Header -->
    <div class="flash-header d-flex align-center justify-space-between px-4">
      <div class="d-flex align-center">
        <v-icon color="white" class="mr-2">mdi-flash</v-icon>
        <span class="text-white font-weight-bold">
          Flash Sales | Live Now
        </span>
      </div>

      <div class="text-white font-weight-medium">
        Time Left:
        <span class="ml-1">{{ timeLeft }}</span>
      </div>

      <v-btn variant="text" color="white" size="small">
        See All →
      </v-btn>
    </div>

    <!-- Products -->
    <div class="flash-products d-flex px-3 py-4">
      <div
        v-for="(item, i) in items"
        :key="i"
        class="flash-item mr-3"
      >
        <v-img
          :src="item.image"
          aspect-ratio="1"
          cover
          class="rounded mb-2"
        />

        <div class="text-body-2 truncate">
          {{ item.title }}
        </div>

        <div class="font-weight-bold">
          KSh {{ item.price }}
        </div>

        <div class="text-caption text-grey">
          <s>KSh {{ item.oldPrice }}</s>
        </div>

        <div class="text-caption mt-1">
          {{ item.left }} items left
        </div>

        <v-progress-linear
          :model-value="item.progress"
          height="6"
          rounded
          color="orange"
          class="mt-1"
        />
      </div>
    </div>
  </v-card>
</template>
<script setup>
import { ref } from "vue";

const timeLeft = ref("01h : 50m : 14s");

const items = ref([
  {
    title: "Rashnik RN-999 2 In 1 Blender",
    price: "1,490",
    oldPrice: "1,716",
    left: 50,
    progress: 40,
    image: "/images/5.jpg",
  },
  {
    title: "Nunix 100 Litres Freezer",
    price: "19,665",
    oldPrice: "22,950",
    left: 15,
    progress: 10,
    image: "/images/1.jpg",
  },
  {
    title: "NIVEA Nourishing Cocoa Lotion",
    price: "896",
    oldPrice: "1,280",
    left: 30,
    progress: 55,
    image: "/images/2.jpg",
  },
  {
    title: "Infinix Smart 10, 6.7\"",
    price: "8,100",
    oldPrice: "10,900",
    left: 50,
    progress: 35,
    image: "/images/3.jpg",
  },
  {
    title: "Vision Plus VP8843KV",
    price: "23,999",
    oldPrice: "35,068",
    left: 25,
    progress: 60,
    image: "/images/4.jpg",
  },
  {
    title: "Poco C85, 6.9\" 128GB",
    price: "12,200",
    oldPrice: "15,000",
    left: 89,
    progress: 20,
    image: "/images/2.jpg",
  },
]);
</script>

<style scoped>
.flash-wrapper {
  background: #e50000;
}

.flash-header {
  height: 48px;
  background: #e50000;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.flash-products {
  background: #fff;
  overflow-x: auto;
}

.flash-item {
  width: 170px;
  flex-shrink: 0;
}

.truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
