<template>
  <v-row class="mt-5 pa-4 ga-2" style="max-width: 1550px; margin: 0 auto">
    <TopBanner />
    <FlashSalesRow />
    <TopDealsRow />
    <BottomSEOSection />
  </v-row>
</template>

<script setup>
import BottomSEOSection from "@/components/bottom/BottomSEOSection.vue";
import FlashSalesRow from "@/components/FlashSalesRow.vue";
import TopBanner from "@/components/TopBanner.vue";
import TopDealsRow from "@/components/TopDealsRow.vue";
</script>
