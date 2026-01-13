<template>
  <v-container fluid class="pa-0 bg-grey-lighten-4 mt-8">
    <v-row class="ma-0 mt-8 " style="max-width: 1550px; margin: 0 auto !important">
      <!-- Sidebar Navigation -->
      <v-col cols="12" md="3" class="pa-4 mt-5 mb-8">
        <v-card elevation="0" class="pa-0">
          <v-list-item class="px-4 py-3 bg-orange">
            <v-list-item-title class="text-white font-weight-medium">
              Admin Dashboard
            </v-list-item-title>
          </v-list-item>

          <v-divider></v-divider>

          <!-- Main Navigation Items -->
          <v-list density="compact" class="pa-0">
            <v-list-item
              :to="{ name: 'admin-dashboard' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-view-dashboard</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Dashboard
              </v-list-item-title>
            </v-list-item>

            <!-- Products Management -->
            <v-list-group>
              <template v-slot:activator="{ props }">
                <v-list-item v-bind="props" class="px-4 py-3">
                  <template v-slot:prepend>
                    <v-icon>mdi-package-variant</v-icon>
                  </template>
                  <v-list-item-title class="text-body-2">
                    Products
                  </v-list-item-title>
                </v-list-item>
              </template>

              <v-list-item
                :to="{ name: 'admin-products-list' }"
                class="px-8 py-2"
                active-class="bg-grey-lighten-2"
              >
                <v-list-item-title class="text-body-2">
                  All Products
                </v-list-item-title>
              </v-list-item>

              <v-list-item
                :to="{ name: 'admin-products-add' }"
                class="px-8 py-2"
                active-class="bg-grey-lighten-2"
              >
                <v-list-item-title class="text-body-2">
                  Add Product
                </v-list-item-title>
              </v-list-item>

              <v-list-item
                :to="{ name: 'admin-flash-sales' }"
                class="px-8 py-2"
                active-class="bg-grey-lighten-2"
              >
                <v-list-item-title class="text-body-2">
                  Flash Sales
                </v-list-item-title>
              </v-list-item>

              <v-list-item
                :to="{ name: 'admin-top-deals' }"
                class="px-8 py-2"
                active-class="bg-grey-lighten-2"
              >
                <v-list-item-title class="text-body-2">
                  Top Deals
                </v-list-item-title>
              </v-list-item>
            </v-list-group>

            <v-list-item
              :to="{ name: 'admin-orders' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-shopping</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Orders
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'admin-customers' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-account-group</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Customers
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'admin-categories' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-shape</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Categories
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'admin-analytics' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-chart-line</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Analytics
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'admin-settings' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-cog</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Settings
              </v-list-item-title>
            </v-list-item>

            <v-divider class="my-2"></v-divider>

            <!-- Back to Store -->
            <v-list-item class="px-4 py-2" @click="goToStore">
              <template v-slot:prepend>
                <v-icon>mdi-storefront</v-icon>
              </template>
              <v-list-item-title class="text-body-2 text-primary">
                Back to Store
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>

      <!-- Main Content Area -->
      <v-col cols="12" md="9" class="pa-4">
        <router-view />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

const goToStore = () => {
  router.push({ name: "home" });
};
</script>
