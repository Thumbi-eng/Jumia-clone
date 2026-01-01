<template>
  <v-app class="bg-orange overflow-y-auto">
    <v-app-navbar>
      <Navbar />
    </v-app-navbar>

    <!-- Main content - Router view for page switching -->
    <router-view />

    <!-- footer section -->
    <v-app-footer class="mt-10">
      <Newsletter />
      <AppFooter />
    </v-app-footer>
  </v-app>
</template>

<script setup>
import Navbar from "./components/Navbar.vue";
import AppFooter from "./components/footer/AppFooter.vue";
import Newsletter from "./components/footer/Newsletter.vue";
</script>
