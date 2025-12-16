<template>
  <v-app-bar app color="white" elevation="1" style="min-height: 64px; min-width: 1220px">
    <v-container class="pa-4 mr-8" fluid style="min-width: 1220px">
      <v-row align="center" no-gutters class="ga-2">
        <!-- Left: logo -->

        <v-img
          src="/Jumia-Logo.jpg"
          alt="Jumia Logo"
          style="object-fit: contain"
          class=""
          height="100"
        />

        <!-- Center: search -->
        <v-col>
          <div
            :style="{
              display: 'flex',
              alignItems: 'center',
              width: '100%',
              maxWidth: searchMaxWidth + 'px',
              margin: '0 auto',
            }"
          >
            <v-text-field
              v-model="query"
              placeholder="Search products, brands and categories"
              hide-details
              density="compact"
              variant="outlined"
              class="me-2"
              prepend-inner-icon="mdi-magnify"
              color="grey-darken-3"
              style="flex: 1 1 auto; background: white"
              @click:append-inner="onSearch"
            />
            <v-btn
              color="orange darken-1"
              variant="flat"
              @click="onSearch"
              style="min-width: 92px"
              class="text-white"
            >
              Search
            </v-btn>
          </div>
        </v-col>

        <!-- Right: account / help / cart -->
        <v-col class="ga-3">
          <v-btn class="ga-2 hover-primary" >
            <v-icon size="32">mdi-account-outline</v-icon>
            Account
            <v-icon size="24">mdi-chevron-down</v-icon>
          </v-btn>

          <v-btn text class="pa-1 hover-primary" >
            <v-icon size="26">mdi-help-circle-outline</v-icon>
            Help 
            <v-icon size="26">mdi-chevron-down</v-icon></v-btn>

          <v-btn class="hover-primary" >
             <v-icon size="32">mdi-cart-outline</v-icon> Cart </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>
</template>

<script setup>
import { ref, computed } from "vue";
import { useDisplay } from "vuetify";

const query = ref("");
const display = useDisplay();

const logoMaxWidth = computed(() => (display.mdAndDown.value ? 110 : 140));
const searchMaxWidth = computed(() => (display.mdAndDown.value ? 420 : 700));

function onSearch() {
  // simple behavior: navigate to search or console log for now
  // This can be replaced with proper search routing later
  console.log("Search:", query.value);
}
</script>

<style scoped>
.hover-primary:hover {
  color: rgb(var(--v-theme-primary));
}

.hover-primary:hover .v-icon {
  color: rgb(var(--v-theme-primary));
}
</style>

