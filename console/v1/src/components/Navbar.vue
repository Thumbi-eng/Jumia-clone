<template>
  <v-app-bar
    app
    color="white"
    elevation="1"
    style="min-width: 1220px"
    class="mb-3"
  >
    <v-container class="pa-2 mr-8" fluid style="min-width: 1220px">
      <v-row align="center" no-gutters class="ga-2">
        <!-- Left: logo -->

        <v-img
          src="/Jumia-Logo.jpg"
          alt="Jumia Logo"
          style="object-fit: contain"
          class=""
          height="100"
          @click="$router.push({ name: 'home' })"
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
              @keyup.enter="onSearch"
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
          <!-- Logged In User Menu -->
          <v-menu
            v-if="authStore.isAuthenticated"
            location="bottom end"
            offset="8"
            transition="scale-transition"
          >
            <template #activator="{ props }">
              <v-btn v-bind="props" class="ga-2 hover-primary" variant="text">
                <v-icon size="32">mdi-account-outline</v-icon>
                Hi, {{ authStore.user?.first_name || "User" }}
                <v-icon size="24">mdi-chevron-down</v-icon>
              </v-btn>
            </template>

            <!-- Dropdown Card -->
            <v-card min-width="240" elevation="8" rounded="lg">
              <!-- Header -->
              <v-list-item class="py-3">
                <template #prepend>
                  <v-avatar color="orange-lighten-4">
                    <v-icon color="orange">mdi-account</v-icon>
                  </v-avatar>
                </template>

                <v-list-item-title class="font-weight-medium">
                  {{ authStore.fullName }}
                </v-list-item-title>
                <v-list-item-subtitle class="text-caption">
                  {{ authStore.user?.email }}
                </v-list-item-subtitle>
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
                @click="handleLogout"
              >
                <v-list-item-title class="text-orange font-weight-medium">
                  Logout
                </v-list-item-title>
              </v-list-item>
            </v-card>
          </v-menu>

          <!-- Not Logged In - Login Button -->
          <v-btn
            v-else
            class="ga-2 hover-primary"
            variant="text"
            @click="$router.push('/login')"
          >
            <v-icon size="32">mdi-account-outline</v-icon>
            Sign In
          </v-btn>

          <v-btn text class="pa-1 hover-primary">
            <v-icon size="26">mdi-help-circle-outline</v-icon>
            Help
            <v-icon size="26">mdi-chevron-down</v-icon></v-btn
          >

          <v-btn class="hover-primary" :to="{ name: 'cart' }">
            <v-badge
              :content="cartStore.cartCount"
              color="orange darken-1 text-white"
              offset-x="5"
              offset-y="5"
              :model-value="cartStore.cartCount > 0"
            >
              <v-icon size="32">mdi-cart-outline</v-icon>
            </v-badge>
            Cart
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>
</template>

<script setup>
import { ref, computed } from "vue";
import { useDisplay } from "vuetify";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useCartStore } from "@/stores/cart";

const router = useRouter();
const authStore = useAuthStore();
const cartStore = useCartStore();
const query = ref("");
const display = useDisplay();

const logoMaxWidth = computed(() => (display.mdAndDown.value ? 110 : 140));
const searchMaxWidth = computed(() => (display.mdAndDown.value ? 420 : 700));

function onSearch() {
  if (!query.value || query.value.trim() === "") {
    return;
  }

  // Navigate to search results page with query parameter
  router.push({
    name: "search",
    query: { q: query.value.trim() },
  });
}

const menuItems = [
  {
    title: "My Account",
    icon: "mdi-account-outline",
    to: "/customer/account/index",
  },
  { title: "Orders", icon: "mdi-package-variant", to: "/customer/order/index" },
  { title: "Inbox", icon: "mdi-email-outline", to: "/customer/inbox/index" },
  {
    title: "Wishlist",
    icon: "mdi-heart-outline",
    to: "/customer/wishlist/index",
  },
  {
    title: "Vouchers",
    icon: "mdi-ticket-percent-outline",
    to: "/customer/vouchers/index",
  },
];

function handleLogout() {
  authStore.logout();
  router.push("/login");
}
</script>

<style scoped>
.hover-primary:hover {
  color: rgb(var(--v-theme-primary));
}

.hover-primary:hover .v-icon {
  color: rgb(var(--v-theme-primary));
}
.logout-item {
  cursor: pointer;
}

.logout-item:hover {
  background-color: rgba(255, 152, 0, 0.08);
}
</style>
