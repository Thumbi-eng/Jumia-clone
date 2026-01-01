<template>
  <v-container fluid class="pa-6 bg-grey-lighten-4 mt-6">
    <v-row class="mt-6" style="max-width: 1550px; margin: 2px auto !important">
      <!-- Sidebar Navigation -->
      <v-col cols="12" md="3" class="pa-4">
        <v-card elevation="0" class="pa-0">
          <!-- Main Navigation Items -->
          <v-list density="compact" class="pa-0">
            <v-list-item
              :to="{ name: 'customer-account' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-account-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                My Jumia Account
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-orders' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-package-variant</v-icon>
              </template>
              <v-list-item-title class="text-body-2">Orders</v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-inbox' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-email-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">Inbox</v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-reviews' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-comment-edit-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Pending Reviews
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-vouchers' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-ticket-percent-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Vouchers
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-wishlist' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-heart-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Wishlist
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-followed-sellers' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-store-outline</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Followed Sellers
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-recently-viewed' }"
              class="px-4 py-3"
              active-class="bg-grey-lighten-2"
            >
              <template v-slot:prepend>
                <v-icon>mdi-history</v-icon>
              </template>
              <v-list-item-title class="text-body-2">
                Recently Viewed
              </v-list-item-title>
            </v-list-item>

            <v-divider class="my-2"></v-divider>

            <!-- Account Settings -->
            <v-list-item
              :to="{ name: 'customer-account-management' }"
              class="px-4 py-2"
              active-class="bg-grey-lighten-2"
            >
              <v-list-item-title class="text-body-2">
                Account Management
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-payment-settings' }"
              class="px-4 py-2"
              active-class="bg-grey-lighten-2"
            >
              <v-list-item-title class="text-body-2">
                Payment Settings
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-address-book' }"
              class="px-4 py-2"
              active-class="bg-grey-lighten-2"
            >
              <v-list-item-title class="text-body-2">
                Address Book
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-newsletter' }"
              class="px-4 py-2"
              active-class="bg-grey-lighten-2"
            >
              <v-list-item-title class="text-body-2">
                Newsletter Preferences
              </v-list-item-title>
            </v-list-item>

            <v-list-item
              :to="{ name: 'customer-close-account' }"
              class="px-4 py-2"
              active-class="bg-grey-lighten-2"
            >
              <v-list-item-title class="text-body-2">
                Close Account
              </v-list-item-title>
            </v-list-item>

            <v-divider class="my-2"></v-divider>

            <!-- Logout -->
            <v-list-item class="px-4 py-2" @click="handleLogout">
              <v-list-item-title class="text-orange text-body-2">
                Logout
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
const handleLogout = () => {
  // Add logout logic here
  console.log("Logout clicked");
};
</script>
