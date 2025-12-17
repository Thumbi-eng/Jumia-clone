<template>
  <v-app class="bg-orange overflow-y-auto" >
    <v-app-navbar>
      <Navbar />
    </v-app-navbar>

    <!-- Main content would go here -->
    <div class="mt-5 pa-4 ga-2" style="max-width: 1550px; margin: 0 auto">
      <TopBanner />
      <FlashSalesRow/>
      <TopDealsRow/>
      <BottomSEOSection />

    </div>

    <!-- footer section -->
    <v-app-footer class="mt-10">
      <Newsletter/>
      <AppFooter />
    </v-app-footer>
  </v-app>
</template>

<script setup>
import BottomSEOSection from "./components/bottom/BottomSEOSection.vue";
import FlashSalesRow from "./components/FlashSalesRow.vue";
import Navbar from "./components/Navbar.vue";
import TopBanner from "./components/TopBanner.vue";
import TopDealsRow from "./components/TopDealsRow.vue";
import AppFooter from "./components/footer/AppFooter.vue";
import Newsletter from "./components/footer/Newsletter.vue";

//
</script>
