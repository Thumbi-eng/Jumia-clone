<template>
  <v-app-bar app color="white" elevation="1" style="min-width: 1220px">
    <v-container class="pa-2 mr-8" fluid style="min-width: 1220px">
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
          <!-- <v-btn class="ga-2 hover-primary">
            <v-icon size="32">mdi-account-outline</v-icon>
            Account
            <v-icon size="24">mdi-chevron-down</v-icon>
          </v-btn> -->
          <v-menu
  location="bottom end"
  offset="8"
  transition="scale-transition"
>
  <template #activator="{ props }">
    <v-btn
      v-bind="props"
      class="ga-2 hover-primary"
      variant="text"
    >
      <v-icon size="32">mdi-account-outline</v-icon>
      Account
      <v-icon size="24">mdi-chevron-down</v-icon>
    </v-btn>
  </template>

  <!-- Dropdown Card -->
  <v-card
    min-width="240"
    elevation="8"
    rounded="lg"
  >

      <!-- Header -->
    <v-list-item class="py-3">
      <template #prepend>
        <v-avatar color="grey-lighten-3">
          <v-icon>mdi-account</v-icon>
        </v-avatar>
      </template>

      <v-list-item-title class="font-weight-medium">
        Hi, victor
      </v-list-item-title>
    </v-list-item>

    <v-divider />

    <!-- Menu Items -->
    <v-list density="compact" nav>
      <v-list-item
        v-for="item in menuItems"
        :key="item.title"
        :to="item.to"
        link
      >
        <template #prepend>
          <v-icon>{{ item.icon }}</v-icon>
        </template>

        <v-list-item-title>
          {{ item.title }}
        </v-list-item-title>
      </v-list-item>
    </v-list>

    <v-divider />

    <!-- Logout -->
    <v-list-item
      class="text-center logout-item"
      @click="logout"
    >
      <v-list-item-title class="text-orange font-weight-medium">
        Logout
      </v-list-item-title>
    </v-list-item>
  </v-card>
</v-menu>






          <v-btn text class="pa-1 hover-primary">
            <v-icon size="26">mdi-help-circle-outline</v-icon>
            Help
            <v-icon size="26">mdi-chevron-down</v-icon></v-btn
          >

          <v-btn class="hover-primary">
            <v-icon size="32">mdi-cart-outline</v-icon> Cart
          </v-btn>
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

const menuItems = [
  { title: 'My Account', icon: 'mdi-account-outline', to: '/account' },
  { title: 'Orders', icon: 'mdi-package-variant', to: '/orders' },
  { title: 'Inbox', icon: 'mdi-email-outline', to: '/inbox' },
  { title: 'Wishlist', icon: 'mdi-heart-outline', to: '/wishlist' },
  { title: 'Vouchers', icon: 'mdi-ticket-percent-outline', to: '/vouchers' }
]

const logout = () => {
  // Supabase / Auth logout here
  console.log('Logged out')
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
